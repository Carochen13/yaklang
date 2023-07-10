package analyzer

import (
	"bufio"
	"bytes"
	"github.com/yaklang/yaklang/common/sca/dxtypes"
	"io"
	"net/textproto"
	"strings"

	"github.com/yaklang/yaklang/common/utils"
)

const (
	TypDPKG TypAnalyzer = "dpkg-pkg"

	statusFile = "var/lib/dpkg/status"
	statusDir  = "var/lib/dpkg/status.d/"
	// availableFile = "var/lib/dpkg/available"
	// infoDir       = "var/lib/dpkg/info/"

	statusStatus int = 1
)

func init() {
	RegisterAnalyzer(TypDPKG, NewDpkgAnalyzer())
}

type dpkgAnalyzer struct {
}

// ReadBlock reads Analyzer data block from the underlying reader until Analyzer blank line is encountered.
func ReadBlock(r *bufio.Reader) ([]byte, error) {
	var block bytes.Buffer

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			return block.Bytes(), err
		}

		if line == "\n" {
			break
		}

		block.WriteString(line)
	}

	return block.Bytes(), nil
}

func (a dpkgAnalyzer) parseStatus(s string) bool {
	for _, ss := range strings.Fields(s) {
		if ss == "deinstall" || ss == "purge" {
			return false
		}
	}
	return true
}
func (a dpkgAnalyzer) parseDpkgPkg(header textproto.MIMEHeader) *dxtypes.Package {
	status := header.Get("Status")
	if status == "" {
		return nil
	}
	if isInstalled := a.parseStatus(status); !isInstalled {
		return nil
	}

	pkg := &dxtypes.Package{
		Name:    header.Get("Package"),
		Version: header.Get("Version"),
	}
	if pkg.Name == "" || pkg.Version == "" {
		return nil
	}

	return pkg
}

func (a dpkgAnalyzer) analyzeStatus(r io.Reader) ([]dxtypes.Package, error) {
	pkgs := make([]dxtypes.Package, 0)
	br := bufio.NewReader(r)
	for {
		block, err := ReadBlock(br)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if block == nil {
			break
		}
		reader := textproto.NewReader(bufio.NewReader(bytes.NewReader(block)))
		header, err := reader.ReadMIMEHeader()
		if err != nil && err != io.EOF {
			return nil, utils.Errorf("parse MIME header error: %v ", err)
		}
		pkg := a.parseDpkgPkg(header)
		if pkg != nil {
			pkgs = append(pkgs, *pkg)
		}
	}

	return pkgs, nil
}

func NewDpkgAnalyzer() *dpkgAnalyzer {
	return &dpkgAnalyzer{}
}

func (a dpkgAnalyzer) Match(info MatchInfo) int {
	if strings.HasPrefix(info.path, statusDir) || info.path == statusFile {
		// handler status
		return statusStatus
	}
	return 0
}

func (a dpkgAnalyzer) Analyze(afi AnalyzeFileInfo) ([]dxtypes.Package, error) {
	fi := afi.Self
	switch fi.MatchStatus {
	case statusStatus:
		return a.analyzeStatus(fi.File)
	}

	return nil, nil
}

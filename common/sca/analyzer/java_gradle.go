package analyzer

import (
	"github.com/yaklang/yaklang/common/sca/dxtypes"
	"strings"

	"github.com/aquasecurity/go-dep-parser/pkg/gradle/lockfile"
)

const (
	TypGradle TypAnalyzer = "gradle-lang"

	javaGradleFile = "gradle.lockfile"

	statusGradle int = 1
)

func init() {
	RegisterAnalyzer(TypGradle, NewJavaGradleAnalyzer())
}

type gradleAnalyzer struct{}

func NewJavaGradleAnalyzer() *gradleAnalyzer {
	return &gradleAnalyzer{}
}

func (a gradleAnalyzer) Analyze(afi AnalyzeFileInfo) ([]dxtypes.Package, error) {
	fi := afi.Self
	switch fi.MatchStatus {
	case statusGradle:
		p := lockfile.NewParser()
		pkgs, err := ParseLanguageConfiguration(fi, p)
		if err != nil {
			return nil, err
		}
		return pkgs, nil
	}
	return nil, nil
}

func (a gradleAnalyzer) Match(info MatchInfo) int {
	// Skip `composer.lock` inside `vendor` folder
	if strings.HasSuffix(info.path, javaGradleFile) {
		return statusGradle
	}
	return 0
}
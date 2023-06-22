package vulinbox

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	uuid "github.com/satori/go.uuid"
	"github.com/yaklang/yaklang/common/coreplugin"
	"github.com/yaklang/yaklang/common/crawler"
	"github.com/yaklang/yaklang/common/log"
	utils2 "github.com/yaklang/yaklang/common/utils"
	"github.com/yaklang/yaklang/common/yak"
	"github.com/yaklang/yaklang/common/yak/yaklang"
	"github.com/yaklang/yaklang/common/yak/yaklib"
	"github.com/yaklang/yaklang/common/yakgrpc/yakit"
	"github.com/yaklang/yaklang/common/yakgrpc/ypb"
	"strings"
	"testing"
)

type vulBoxTester struct {
	testCase      map[string][]string
	ignore        map[string][]string
	plugins       map[string]struct{}
	risks         []*yakit.Risk
	serverAdderss string
	t             *testing.T
}

func newVulBoxTester(t *testing.T) *vulBoxTester {
	return &vulBoxTester{
		t:        t,
		testCase: map[string][]string{},
		plugins:  map[string]struct{}{},
		ignore:   map[string][]string{},
	}
}
func (v *vulBoxTester) addIgnoreInfo(pluginName, route string, risks ...string) {
	if v1, ok := v.testCase[route]; ok {
		v.ignore[route] = append(v1, risks...)
	} else {
		v.ignore[route] = risks
	}
	v.plugins[pluginName] = struct{}{}
}
func (v *vulBoxTester) addTestCase(pluginName, route string, risks ...string) {
	if v1, ok := v.testCase[route]; ok {
		v.testCase[route] = append(v1, risks...)
	} else {
		v.testCase[route] = risks
	}
	v.plugins[pluginName] = struct{}{}
}

func (v *vulBoxTester) run() {
	serverAdderss, err := NewVulinServer(context.Background())
	if err != nil {
		panic(err)
	}
	v.serverAdderss = serverAdderss
	yaklib.RiskExports["NewRisk"] = func(target string, opts ...yakit.RiskParamsOpt) {
		r := &yakit.Risk{
			Hash: uuid.NewV4().String(),
		}
		for _, opt := range opts {
			opt(r)
		}
		r.Url = target
		v.risks = append(v.risks, r)
	}
	yaklang.Import("risk", yaklib.RiskExports)
	manager, err := yak.NewMixPluginCaller()
	if err != nil {
		log.Error(err)
	}
	manager.SetFeedback(func(i *ypb.ExecResult) error {
		return nil
	})
	for fileName, _ := range v.plugins {
		var pluginName = strings.TrimSuffix(fileName, ".yak")
		var scriptBytes = coreplugin.GetCorePluginData(pluginName)
		var err = manager.LoadHotPatch(context.Background(), string(scriptBytes))
		if err != nil {
			panic(fmt.Sprintf("load plugin %v failed: %s", pluginName, err))
		}
	}

	manager.SetDividedContext(true)
	manager.SetConcurrent(20)

	ch := make(chan *crawler.Req)

	crawler, err := crawler.NewCrawler(v.serverAdderss, crawler.WithOnRequest(func(req *crawler.Req) {
		ch <- req
	}), crawler.WithMaxDepth(1))
	if err != nil {
		panic(err)
	}
	go func() {
		defer close(ch)
		err := crawler.Run()
		if err != nil {
			log.Error(err)
		}
	}()
	urlPathMap := make(map[string]bool)
	for req := range ch {
		if _, ok := urlPathMap[req.Url()]; ok {
			continue
		} else {
			manager.MirrorHTTPFlowEx(false, req.IsHttps(), req.Url(), req.RequestRaw(), req.ResponseRaw(), req.ResponseBody())
			urlPathMap[req.Url()] = true
		}

	}
	manager.Wait()
	spew.Dump(v.risks)
	for _, risk := range v.risks {
		foundUrl := false
		tmp := make(map[string][]string)
		for url, titles := range v.ignore {
			tmp[url] = titles
		}
		for url, titles := range v.testCase {
			tmp[url] = titles
		}
		for url, titles := range tmp {
			riskUrl, err := utils2.ParseStringUrlToUrlInstance(risk.Url)
			if err != nil {
				v.t.Fatal(fmt.Sprintf("risk url illegal %v", risk.Url))
			}
			if riskUrl.Path == url {
				foundUrl = true
				ok := false
				for i, title := range titles {
					if strings.Contains(risk.Title, title) {
						ok = true
						v.testCase[url] = append(titles[:i], titles[i+1:]...)
						break
					}
				}
				if !ok { // 如果发现的风险信息不在TestCase里则直接Fail
					v.t.Fatal(fmt.Sprintf("there is no %v vulnerability in route %v", risk.Title, risk.Url))
				}
			}
		}
		if !foundUrl { // 如果发现的风险URL不在TestCase里则直接Fail
			v.t.Fatal(fmt.Sprintf("there is no %v vulnerability in route %v", risk.Title, risk.Url))
		}
	}
	// 如果有风险信息没有扫描出来，则Fail
	for url, titles := range v.testCase {
		if len(titles) > 0 {
			v.t.Fatal(fmt.Sprintf("route %v,valnerability: %v was not discovered", url, titles))
		}
	}
}
func TestSSRF(t *testing.T) {
	tester := newVulBoxTester(t)
	tester.addTestCase("启发式SQL注入检测.yak", "/user/name", "Maybe SQL Injection: [param - type:str value:admin single-quote]")
	tester.addTestCase("启发式SQL注入检测.yak", "/user/id", "Maybe SQL Injection: [param - type:str value:1 single-quote]")
	tester.addTestCase("启发式SQL注入检测.yak", "/user/id", "Union-Based SQL Injection: [id:[1]]")
	tester.addTestCase("启发式SQL注入检测.yak", "/user/id-json", "Maybe SQL Injection: [param - type:str value:1 single-quote]")
	tester.addTestCase("启发式SQL注入检测.yak", "/user/id-json", "Union-Based SQL Injection: [id:[1]]")
	tester.addTestCase("启发式SQL注入检测.yak", "/user/id-b64-json", "Maybe SQL Injection: [param - type:str value:1 single-quote]")
	tester.addTestCase("启发式SQL注入检测.yak", "/user/id-b64-json", "Union-Based SQL Injection: [id:[1]]")
	tester.addTestCase("启发式SQL注入检测.yak", "/user/name", "Union-Based SQL Injection: [name:[admin]]")
	tester.addTestCase("启发式SQL注入检测.yak", "/user/by-id-safe")
	tester.addIgnoreInfo("启发式SQL注入检测.yak", "/ping/cmd/shlex", "ORDER BY SQL Injection: [ip:[127.0.0.1]]")
	tester.addIgnoreInfo("启发式SQL注入检测.yak", "/ping/cmd/bash", "ORDER BY SQL Injection: [ip:[127.0.0.1]]")
	tester.run()
}

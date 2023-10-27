package httptpl

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/yaklang/yaklang/common/go-funk"
	"github.com/yaklang/yaklang/common/log"
	"github.com/yaklang/yaklang/common/utils"
	"github.com/yaklang/yaklang/common/utils/lowhttp"
	"github.com/yaklang/yaklang/common/yakgrpc/ypb"
	"net/url"
	"path"
	"strings"
	"sync/atomic"
)

type requestRaw struct {
	Raw     []byte
	IsHttps bool
	Params  map[string]interface{}
	Origin  *YakRequestBulkConfig
}

type RequestBulk struct {
	Requests      []*requestRaw
	RequestConfig *YakRequestBulkConfig
}

func (y *YakTemplate) RenderToFuzzerRequestsByUrl(u string) []*ypb.FuzzerRequests {
	result := []*ypb.FuzzerRequests{}
	for _, sequence := range y.HTTPRequestSequences {
		seq := &ypb.FuzzerRequests{}
		for _, path := range sequence.Paths {
			var firstLine string = fmt.Sprintf("%v %v HTTP/1.1", sequence.Method, path)
			var lines []string
			lines = append(lines, firstLine)
			_, hostOk1 := sequence.Headers["Host"]
			_, hostOk2 := sequence.Headers["host"]
			if !hostOk1 && !hostOk2 {
				lines = append(lines, "Host: "+"{{Hostname}}")
			}
			for k, v := range sequence.Headers {
				lines = append(lines, fmt.Sprintf(`%v: %v`, k, v))
			}
			if len(sequence.Headers) <= 0 {
				lines = append(lines, `User-Agent: Mozilla/5.0 (Windows NT 10.0; rv:78.0) Gecko/20100101 Firefox/78.0`)
			}
			var rawPacket = strings.Join(lines, "\r\n") + "\r\n\r\n"
			rawPacket += sequence.Body
			seq.Requests = append(seq.Requests, &ypb.FuzzerRequest{
				Request:    rawPacket,
				RequestRaw: toBytes(rawPacket),
			})
		}
		result = append(result, seq)
	}
	return result
}
func (y *YakTemplate) GenerateRequestSequences(u string) []*RequestBulk {
	urlIns, err := url.Parse(u)
	if err != nil {
		log.Error(err)
		return nil
	}
	baseUrl := urlIns.String()
	var rootUrl string
	if urlIns.Scheme == "https" {
		if urlIns.Port() == "443" {
			rootUrl = fmt.Sprintf("https://%v", urlIns.Host)
		} else {
			rootUrl = fmt.Sprintf("https://%v", utils.HostPort(urlIns.Host, urlIns.Port()))
		}
	} else {
		if urlIns.Port() == "80" {
			rootUrl = fmt.Sprintf("http://%v", urlIns.Host)
		} else {
			rootUrl = fmt.Sprintf("http://%v", utils.HostPort(urlIns.Host, urlIns.Port()))
		}
	}
	var file string
	pathRaw := urlIns.RequestURI()
	if strings.Contains(pathRaw, "?") {
		pathNoQuery := pathRaw[:strings.Index(pathRaw, "?")]
		_, file = path.Split(pathNoQuery)
	}
	baseUrl = strings.TrimRight(baseUrl, "/")
	rootUrl = strings.TrimRight(rootUrl, "/")
	vars := map[string]string{
		"URL":              urlIns.String(),
		"Host":             urlIns.Host,
		"Port":             urlIns.Port(),
		"Hostname":         urlIns.Hostname(),
		"RootURL":          rootUrl,
		"BaseURL":          baseUrl,
		"Path":             urlIns.RequestURI(),
		"PathTrimEndSlash": strings.TrimRight(urlIns.RequestURI(), "/"),
		"File":             file,
		"Schema":           urlIns.Scheme,
	}
	for k, v := range vars {
		y.Variables.Set(k, v)
	}
	result := []*RequestBulk{}
	for _, sequenceCfg := range y.HTTPRequestSequences {
		seq := &RequestBulk{
			RequestConfig: sequenceCfg,
		}
		for _, path := range sequenceCfg.Paths {
			path, err = ExecNucleiTag(path, y.Variables.ToMap())
			if err != nil {
				log.Error(err)
				continue
			}
			isHttps, packet, err := lowhttp.ParseUrlToHttpRequestRaw(sequenceCfg.Method, path)
			if err != nil {
				log.Error(err)
				continue
			}
			for k, v := range sequenceCfg.Headers {
				packet = lowhttp.ReplaceHTTPPacketHeader(packet, k, v)
			}
			packet = lowhttp.ReplaceHTTPPacketBody(packet, []byte(sequenceCfg.Body), false)
			packetRaw, err := ExecNucleiTag(string(packet), y.Variables.ToMap())
			if err != nil {
				log.Error(err)
				continue
			}
			seq.Requests = append(seq.Requests, &requestRaw{
				Raw:     []byte(packetRaw),
				Origin:  sequenceCfg,
				IsHttps: isHttps,
			})
		}
		for _, request := range sequenceCfg.HTTPRequests {
			req, err := ExecNucleiTag(request.Request, y.Variables.ToMap())
			if err != nil {
				log.Error(err)
				continue
			}
			seq.Requests = append(seq.Requests, &requestRaw{
				Raw:     []byte(req),
				Origin:  sequenceCfg,
				IsHttps: urlIns.Scheme == "https",
			})
		}
		result = append(result, seq)
	}
	return result
}
func (y *YakTemplate) ExecWithUrl(u string, config *Config, opts ...lowhttp.LowhttpOpt) (int, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
			utils.PrintCurrentGoroutineRuntimeStack()
		}
	}()

	if y.SelfContained {
		log.Debugf("skip self-contained template: %v", y.Name)
		return 0, nil
	}

	if config == nil {
		config = NewConfig()
	}

	var count int64 = 0
	if y.ReverseConnectionNeed {
		var err error
		var require func(...float64) (string, string, error)
		if config.OOBRequireCallback == nil {
			require = RequireOOBAddr
		} else {
			require = config.OOBRequireCallback
		}
		reverse_url, reverse_dnslog_token, err := require(config.OOBTimeout)
		y.Variables.Set("interactsh-url", reverse_url)
		y.Variables.Set("interactsh", reverse_url)
		y.Variables.Set("interactsh_url", reverse_url)
		y.Variables.Set("reverse_dnslog_token", reverse_dnslog_token)
		if err != nil {
			log.Errorf("require oob addr failed: %v", err)
			return 0, err
		}
	}

	tplConcurrent := config.ConcurrentInTemplates
	if len(y.HTTPRequestSequences) > 0 {
		swg := utils.NewSizedWaitGroup(tplConcurrent)
		for _, reqSeq := range y.GenerateRequestSequences(u) {
			swg.Add()
			go func(ret *RequestBulk, payload map[string][]string) {
				defer swg.Done()
				rsps, result, extracted, reqCount := y.handleRequestSequences(config, ret.RequestConfig, ret.Requests, payload, func(reqRaw []byte, isHttps bool) (*lowhttp.LowhttpResponse, error) {
					if config.BeforeSendPackage != nil {
						reqRaw = config.BeforeSendPackage(reqRaw, isHttps)
					}
					var packetOpt []lowhttp.LowhttpOpt
					packetOpt = append(
						packetOpt, lowhttp.WithPacketBytes([]byte(reqRaw)), lowhttp.WithHttps(isHttps),
						lowhttp.WithSaveHTTPFlow(true), lowhttp.WithSource(y.Name),
					)
					packetOpt = append(packetOpt, opts...)
					if config.Debug && config.DebugRequest {
						fmt.Printf("--------------REQ---------------\n")
						fmt.Println(reqRaw)
					}

					utils.Debug(func() {
						log.Info("nuclei lowhttp.Exec! ")
						spew.Dump(reqRaw)
					})
					rsp, err := lowhttp.HTTP(packetOpt...)
					if err != nil {
						log.Error(err)
						return nil, err
					}
					if config.Debug && config.DebugResponse {
						fmt.Printf("--------------RSP---------------\n")
						fmt.Println(string(rsp.RawPacket))
					}
					return rsp, nil
				})
				if result {
					log.Infof("[%v]-[%v] matched", y.Name, y.Id)
				}
				atomic.AddInt64(&count, reqCount)
				config.ExecuteResultCallback(y, ret.RequestConfig, rsps, result, extracted)
			}(reqSeq, reqSeq.RequestConfig.Payloads.GetData())
		}
		swg.Wait()
		return int(count), nil
	} else if len(y.TCPRequestSequences) > 0 {
		swg := utils.NewSizedWaitGroup(tplConcurrent)
		for _, tcpReq := range y.TCPRequestSequences {
			swg.Add()
			tcpReq := tcpReq

			go func() {
				defer swg.Done()
				defer func() {
					if err := recover(); err != nil {
						utils.PrintCurrentGoroutineRuntimeStack()
					}
				}()
				p := y.Variables.ToMap()

				lowhttpConfig := lowhttp.NewLowhttpOption()
				for _, opt := range opts {
					opt(lowhttpConfig)
				}

				err := tcpReq.Execute(config, p, y.PlaceHolderMap, lowhttpConfig, func(response []*NucleiTcpResponse, matched bool, extractorResults map[string]any) {
					atomic.AddInt64(&count, 1)
					config.ExecuteTCPResultCallback(y, tcpReq, response, matched, extractorResults)
					if config.Debug || config.DebugResponse {
						fmt.Println("---------------------TCP RESPONSE---------------------")
						spew.Dump(response)
						fmt.Println("------------------------------------------------------")
					}

					if config.Debug {
						fmt.Println("---------------------TCP RESULT---------------------")
						fmt.Printf("%v Matched: %v\n", y.Name, matched)
						fmt.Println("--------------------- EXTRACTOR ----------------------")
						spew.Dump(extractorResults)
					} else {
						log.Infof("%v Matched: %v", y.Name, matched)
					}
				})
				if err != nil {
					log.Errorf("tcpReq.Execute failed: %s", err)
				}
			}()
		}
		swg.Wait()
		return int(count), nil
	} else {
		return 0, utils.Errorf("[%s] tcp/http is all empty!", y.Name)
	}
}
func (y *YakTemplate) Exec(config *Config, isHttps bool, reqOrigin []byte, opts ...lowhttp.LowhttpOpt) (int, error) {
	urlIns, err := lowhttp.ExtractURLFromHTTPRequestRaw(reqOrigin, isHttps)
	if err != nil {
		return 0, err
	}
	return y.ExecWithUrl(urlIns.String(), config, opts...)
}

// handleRequestSequences 渲染、发包、匹配、提取
func (y *YakTemplate) handleRequestSequences(config *Config, reqOrigin *YakRequestBulkConfig, reqSeqs []*requestRaw, payload map[string][]string, sender func(raw []byte, isHttps bool) (*lowhttp.LowhttpResponse, error)) ([]*lowhttp.LowhttpResponse, bool, map[string]interface{}, int64) {

	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
			utils.PrintCurrentGoroutineRuntimeStack()
		}
	}()
	var count int64 = 0
	if reqOrigin == nil {
		log.Error("request origin cannot be nil")
		return nil, false, map[string]interface{}{}, count
	}

	if reqOrigin.Matcher == nil && len(reqOrigin.Extractor) == 0 {
		log.Error("request sequence matcher and extractor all empty!")
		return nil, false, map[string]interface{}{}, count
	}

	extracted := make(map[string]interface{})

	reqConfig := reqOrigin
	matcher := reqConfig.Matcher

	var matchResults []any
	var responses []*lowhttp.LowhttpResponse

	for index, req := range reqSeqs {
		log.Debugf("sequence exec with Req Index:%v", index)
		raw := req.Raw
		reqs, err := FuzzNucleiTag(string(raw), y.Variables.ToMap(), payload, reqConfig.AttackMode)
		if err != nil {
			log.Errorf("cannot mutate request: %v", err)
			continue
		}
		if len(reqs) <= 0 {
			log.Error("mutate request failed, empty requests")
			continue
		}
		for _, reqRaw := range reqs {
			atomic.AddInt64(&count, 1)
			var sufs = []string{fmt.Sprintf("_%v", index+1)}
			_ = reqRaw
			_ = sufs
			rsp, err := sender([]byte(reqRaw), req.IsHttps)
			if err == nil {
				responses = append(responses, rsp)
			} else {
				log.Error(err)
				continue
			}

			for index, extractor := range reqOrigin.Extractor {
				varIns, err := extractor.Execute(rsp.RawPacket, y.Variables.ToMap())
				if err != nil {
					log.Error("extractor execute failed: ", err)
					continue
				}
				_ = index
				if varIns != nil {
					for k, v := range varIns {
						v := ExtractResultToString(v)
						y.Variables.Set(k, v)
						extracted[k] = v
					}
				}
			}
			if reqOrigin.Matcher != nil {
				var varsInResponse = make(map[string]interface{})
				varsInResponse = LoadVarFromRawResponse(rsp.RawPacket, rsp.GetDurationFloat(), sufs...)
				if varsInResponse != nil {
					for k, v := range varsInResponse {
						y.Variables.Set(k, toString(v))
					}
				}

				if !reqOrigin.AfterRequested {
					matchResult, err := matcher.ExecuteWithConfig(config, rsp, y.Variables.ToMap())
					if err != nil {
						log.Error("matcher execute failed: ", err)
					}
					matchResults = append(matchResults, matchResult)
					if matchResult && reqOrigin.StopAtFirstMatch {
						// 第一次匹配就退出
						return responses, true, extracted, count
					}
				}
			}
		}
	}

	if reqOrigin.AfterRequested {
		if matcher != nil {
			for _, rsp := range responses {
				matchResult, err := matcher.ExecuteWithConfig(config, rsp, y.Variables.ToMap())
				if err != nil {
					log.Error("matcher execute failed: ", err)
				}
				matchResults = append(matchResults, matchResult)
			}
		}
	}
	return responses, funk.Any(matchResults...), extracted, count
}

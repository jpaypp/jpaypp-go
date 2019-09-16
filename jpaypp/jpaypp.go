package jpaypp

import (
	"bytes"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	// 当前版本的api地址
	apiLiveBase = "https://api.jpay.live.weidun.biz/v1"
	apiSandBoxBase = "https://api.jpay.sandbox.weidun.biz/v1"

	// 当前版本的api生成生成时间
	apiVersion = "2019-08-30"
	// httpclient等待时间
	defaultHTTPTimeout                  = 30 * time.Second
	TotalBackends                       = 1
	APIBackEnd         SupportedBackEnd = "api"
)

var (
	// 默认错误信息返回语言
	AcceptLanguage = "zh-CN"
	// ping++ api统一需要通过Authentication（http BasicAuth），需要在调用时赋值
	Key string

	LogLevel = 2

	//不用默认的defaultClient，自定义httpClient
	httpClient        = &http.Client{Timeout: defaultHTTPTimeout}
	backends          BackEnds
	AccountPrivateKey string
	OsInfo            string
)

type SupportedBackEnd string

// 定义统一后端处理接口
type BackEnd interface {
	Call(method, path, key string, body *url.Values, params []byte, v interface{}) error
}

// 获取当前sdk的版本
func Version() string {
	return "3.2.1"
}

/*2016-02-16 当前情况下没有代码调用了该函数
func SetHttpClient(client *http.Client) {
	httpClient = client
}*/

type BackEnds struct {
	API BackEnd
}

// 通过不同的参数获取不同的后端对象
func GetBackend(backend SupportedBackEnd) BackEnd {
	var ret BackEnd
	switch backend {
	case APIBackEnd:
		if backends.API == nil {
			backends.API = ApiBackEnd{backend, apiLiveBase, httpClient}
		}

		ret = backends.API
	}
	return ret
}

//设定后端处理对象
func SetBackEnd(backend SupportedBackEnd, b BackEnd) {
	switch backend {
	case APIBackEnd:
		backends.API = b
	}
}

func init() {
	var uname string
	switch runtime.GOOS {
	case "windows":
		uname = "windows"
	default:
		cmd := exec.Command("uname", "-a")
		cmd.Stdin = strings.NewReader("some input")
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		cmd.Run()
		uname = out.String()
	}
	m := map[string]interface{}{
		"lang":             "golang",
		"lang_version":     runtime.Version(),
		"bindings_version": Version(),
		"publisher":        "jpaypp",
		"uname":            uname,
	}
	content, _ := JsonEncode(m)
	OsInfo = string(content)
}

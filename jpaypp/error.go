package jpaypp

const (
	InvalidRequest ErrorStatus = "invalid_request_error"
	APIErr         ErrorStatus = "api_error"
)

type (
	// 错误的类型
	ErrorStatus string
	// 错误的状代码
	ErrorMessage string
	// 错误的状代码
	ErrorData string

	// 错误的数据结构
	Error struct {
		Status         string `json:"status"`
		Message        string `json:"message"`
		Data           string `json:"param,omitempty"`
	}
)

//返回当前Error数据的json字符串
func (e *Error) Error() string {
	er, _ := JsonEncode(e)
	return string(er)
}

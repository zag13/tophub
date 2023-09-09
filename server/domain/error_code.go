package domain

const (
	ErrorCodeSystemBusy       = 10000
	ErrorCodeInternalError    = 10001
	ErrorCodeInvalidParameter = 10002
	ErrorCodeInvalidRequest   = 10003
	ErrorCodeInvalidSignature = 10004
)

var errorCodeMessageMap = map[int]string{
	ErrorCodeSystemBusy:       "系统繁忙，请稍后重试",
	ErrorCodeInternalError:    "内部错误",
	ErrorCodeInvalidParameter: "参数错误",
	ErrorCodeInvalidRequest:   "无效的请求",
	ErrorCodeInvalidSignature: "无效的签名",
}

package tool

// 定义常用的响应状态码常量
const (
	SuccessCode    = 200
	FailureCode    = 400
	TokenErrorCode = 401
)

// Result 通用响应结构体
type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func R(code int, msg string, data interface{}) Result {
	return Result{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
func Success() Result {
	return R(SuccessCode, "操作成功", nil)
}
func SuccessData(data interface{}) Result {
	return R(SuccessCode, "操作成功", data)
}
func SuccessMsgData(msg string, data interface{}) Result {
	return R(SuccessCode, msg, data)
}
func Fail() Result {
	return R(FailureCode, "操作失败", nil)
}
func FailMsg(msg string) Result {
	return R(FailureCode, msg, nil)
}
func FailMsgData(msg string, data interface{}) Result {
	return R(FailureCode, "操作失败", data)
}
func TokenError() Result {
	return R(TokenErrorCode, "Token错误", nil)
}

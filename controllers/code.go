package controllers

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeServerBusy
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:      "success",
	CodeInvalidParam: "请求参数错误",
	CodeUserExist:    "用户已经存在",
	CodeServerBusy:   "服务繁忙",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		return codeMsgMap[CodeServerBusy]
	}
	return msg
}

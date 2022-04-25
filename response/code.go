package response

const (
	CodeOk = 200
)

const (
	CodeParamError = 400 + iota
	CodeAuthError
)

const (
	CodeSystemError = 500 + iota
)

var codeMsgMap = map[int]string{
	CodeOk:         "成功",
	CodeParamError: "参数错误",
	CodeAuthError:  "权限错误",
}

func Text(code int) string {
	msg, ok := codeMsgMap[code]
	if !ok {
		msg = codeMsgMap[CodeSystemError]
	}
	return msg
}

package response

type Response struct {
	StatusCode int32  `json:"status_code"`          // 状态码
	StatusMsg  string `json:"status_msg,omitempty"` // 状态信息
}

func BaseSuccess(msg string) Response {
	return Response{OptSuccess, msg}
}

// BaseInputError 由于客户端输入产生的错误
func BaseInputError(msg string) Response {
	return Response{OptInputError, msg}
}

// BaseServerError 由于服务端执行时的出错
func BaseServerError(msg string) Response {
	return Response{OptServerError, msg}
}

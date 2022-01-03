package response

type Msg struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SuccessMsg(data interface{}) *Msg {
	msg := &Msg{
		Code: 200,
		Msg:  "SUCCESS",
		Data: data,
	}
	return msg
}

func FailMsg(msg string) *Msg {
	ret := &Msg{
		Code: 200,
		Msg:  "ERROR",
		Data: msg,
	}
	return ret
}

func FailCodeMsg(code int, msg string) *Msg {
	msgObj := &Msg{
		Code: code,
		Msg:  msg,
	}
	return msgObj
}

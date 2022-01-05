package response

import "github.com/mini-ecs/back-end/pkg/common/error_msg"

type Msg struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SuccessMsg(data interface{}) *Msg {
	msg := &Msg{
		Code: error_msg.Success,
		Msg:  "SUCCESS",
		Data: data,
	}
	return msg
}

func FailCodeMsg(code int, msg string) *Msg {
	msgObj := &Msg{
		Code: code,
		Msg:  msg,
	}
	return msgObj
}

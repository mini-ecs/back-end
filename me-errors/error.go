package me_errors

type Error struct {
	msg string
}

func (e Error) Error() string {
	return e.msg
}

func New(msg string) error {
	return Error{
		msg: msg,
	}
}

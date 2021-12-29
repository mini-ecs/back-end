package errors

type mError struct {
	msg string
}

func (e mError) Error() string {
	return e.msg
}

func New(msg string) error {
	return mError{
		msg: msg,
	}
}

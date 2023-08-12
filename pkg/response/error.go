package response

type Error struct {
	Code uint
	Err  error
}

func NewError(code uint, err error) *Error {
	return &Error{
		Code: code,
		Err:  err,
	}
}

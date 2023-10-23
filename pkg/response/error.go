package response

type Error struct {
	Code int
	Err  error
}

func NewError(code int, err error) *Error {
	return &Error{
		Code: code,
		Err:  err,
	}
}

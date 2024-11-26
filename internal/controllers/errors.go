package controllers

type ErrGeneric struct {
	msg        string
	statusCode int
}

func (e ErrGeneric) Error() string {
	return e.msg
}

func (e ErrGeneric) StatusCode() int {
	return e.statusCode
}

func NewHTTPError(msg string, statusCode int) ErrGeneric {
	return ErrGeneric{msg, statusCode}
}

package controller

type httpError struct {
	code    int
	message string
}

func NewHttpError(code int, message string) httpError {
	return httpError{code: code, message: message}
}

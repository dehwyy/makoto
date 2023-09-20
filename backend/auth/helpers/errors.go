package helpers

type ErrorsHandler struct {
	err error
}

func NewErrorsHandler() *ErrorsHandler {
	return &ErrorsHandler{
		err: nil,
	}
}

func (e *ErrorsHandler) SetErrorIfNotAlready(err error) {
	if e.err == nil {
		e.err = err
	}
}

func (e *ErrorsHandler) GetError() error {
	return e.err
}

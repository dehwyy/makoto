package helpers

type ErrorsHandler struct {
	err error
}

func New() *ErrorsHandler {
	return &ErrorsHandler{
		err: nil,
	}
}

func (e *ErrorsHandler) SetErrorIfNotAlready(err error)

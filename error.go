package argo

type ErrorString string

func (e ErrorString) Error() string {
	return string(e)
}

const (
	ErrConnIsClosed = ErrorString("connection is closed")
)

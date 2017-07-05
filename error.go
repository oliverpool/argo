package argo

type ErrorString string

func (e ErrorString) Error() string {
	return string(e)
}

const (
	ErrConnIsClosed = ErrorString("connection is closed")
)

// ResponseError indicates the error encountered
type ResponseError struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

func (r ResponseError) Error() string {
	return r.Message
}

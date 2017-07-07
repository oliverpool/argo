package argo

// A Client is an aria2 client (https://aria2.github.io/)
//
// It can be constructed with the http.NewClient method of the subpackage argo/rpc/http
type Client struct {
	Caller Caller
}

// Close gracefully closes the Caller
func (c Client) Close() (err error) {
	return c.Caller.Close()
}

func (c Client) mergeOptions(options ...Option) Option {
	return mergeOptions(options...)
}

// Caller allows to perform Requests
type Caller interface {
	Call(method string, reply interface{}, params ...interface{}) error
	CallWithID(method string, reply interface{}, id *string, params ...interface{}) error
	Close() error
}

// Package argo is a library to communicate with an aria2 (https://aria2.github.io/) daemon in Go.
package argo

// GID is the identifier aria2 uses for each download.
// The GID is hex string of 16 characters, thus [0-9a-zA-Z] are allowed and leading zeros must not be stripped.
// The GID all 0 is reserved and must not be used. The GID must be unique, otherwise error is reported and the download is not added.
type GID string

// String returns the value as string
func (g GID) String() string {
	return string(g)
}

// A Client is an aria2 client (https://aria2.github.io/)
//
// It can be constructed with the http.NewClient method of the subpackage argo/rpc/http
type Client struct {
	Caller Caller
}

// Caller allows to perform requests to an aria2 instance
type Caller interface {
	Call(method string, reply interface{}, params ...interface{}) error
	CallWithID(method string, reply interface{}, id *string, params ...interface{}) error
	Close() error
}

// Close gracefully closes the Caller
func (c Client) Close() (err error) {
	return c.Caller.Close()
}

func (c Client) mergeOptions(options ...Option) Option {
	return mergeOptions(options...)
}

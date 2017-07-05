package rpc

import "net"
import "github.com/oliverpool/argo"

// Request represents a JSON-RPC request sent by a client.
type Request struct {
	// JSON-RPC protocol.
	Version string `json:"jsonrpc"`

	// A String containing the name of the method to be invoked.
	Method string `json:"method"`

	// Object to pass as request parameter to the method.
	Params interface{} `json:"params"`

	// The request id. This can be of any type. It is used to match the
	// response with the request that it is replying to.
	ID string `json:"id"`
}

// Response represents a JSON-RPC response to a request
type Response struct {
	GID   string `json:"result"` // GID of the download
	ID    string `json:"id"`
	Error ResponseError
}

// ResponseError indicates the error encountered
type ResponseError struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

// Caller allows to perform Requests
type Caller interface {
	Call(Request) (Response, error)
	Close() error
}

// Notification from aria2
type Notification struct {
	// JSON-RPC protocol.
	Version string `json:"jsonrpc"`

	// A String containing the name of the method to be invoked.
	Method string `json:"method"`

	// Object to pass as request parameter to the method.
	Params []struct {
		GID string `json:"gid"` // GID of the download
	} `json:"params"`

	// The request id. This can be of any type. It is used to match the
	// response with the request that it is replying to.
	ID string `json:"id"`
}

// Identifier returns the rpc method
func (r Notification) Identifier() string {
	return r.Method
}

// GID gathers the GID of the notification
func (r Notification) GID() []string {
	gid := make([]string, len(r.Params))
	for i, g := range r.Params {
		gid[i] = g.GID
	}
	return gid
}

func IsClosedNetworkConnectionError(err error) bool {
	if e, ok := err.(*net.OpError); ok {
		return "use of closed network connection" == e.Err.Error()
	}
	return false
}

func ConvertClosedNetworkConnectionError(err error) error {
	if IsClosedNetworkConnectionError(err) {
		return argo.ErrConnIsClosed
	}
	return err
}

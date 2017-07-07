// Package rpc handles the rpc communication with the aria2 dameon
package rpc

import (
	"encoding/json"
	"net"

	"github.com/oliverpool/argo"
)

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
	Result json.RawMessage `json:"result"`
	ID     string          `json:"id"`
	Error  argo.ResponseError
}

// Poster allows to perform Requests
type Poster interface {
	Post(Request) (Response, error)
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
		GID argo.GID `json:"gid"` // GID of the download
	} `json:"params"`
}

// Identifier returns the rpc method
func (r Notification) Identifier() string {
	return r.Method
}

// GID gathers the GID of the notification
func (r Notification) GID() []argo.GID {
	gid := make([]argo.GID, len(r.Params))
	for i, g := range r.Params {
		gid[i] = g.GID
	}
	return gid
}

func isClosedNetworkConnectionError(err error) bool {
	if e, ok := err.(*net.OpError); ok {
		return "use of closed network connection" == e.Err.Error()
	}
	return false
}

// ConvertClosedNetworkConnectionError converts a TCP "closed network connection" error to a argo.ErrConnIsClosed
func ConvertClosedNetworkConnectionError(err error) error {
	if isClosedNetworkConnectionError(err) {
		return argo.ErrConnIsClosed
	}
	return err
}

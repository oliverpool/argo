package http

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/oliverpool/argo/rpc"
)

// Poster allows to POST an io.Reader to an URL
type Poster interface {
	Post(url string, contentType string, body io.Reader) (*http.Response, error)
}

// Caller allows to send Calls to an URL
type Caller struct {
	Client Poster
	URL    string
}

// Call performs the RPCRequest
func (j Caller) Call(v rpc.Request) (reply rpc.Response, err error) {
	pay, err := json.Marshal(v)
	if err != nil {
		return
	}
	r, err := j.Client.Post(j.URL, "application/json", bytes.NewReader(pay))
	if err != nil {
		return
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&reply)
	return
}

// Close gracefully closes the connection
func (j Caller) Close() (err error) {
	return nil
}

// NewCaller creates a Caller with the http.DefaultClient
func NewCaller(add string) (j Caller, err error) {
	j = Caller{
		Client: http.DefaultClient,
		URL:    add,
	}
	return
}

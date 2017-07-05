package http

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/oliverpool/argo/rpc"
)

// SubPoster allows to POST an io.Reader to an URL
type SubPoster interface {
	Post(url string, contentType string, body io.Reader) (*http.Response, error)
}

// Poster allows to send rpc.Request via http(s)
type Poster struct {
	Client SubPoster
	URL    string
}

// Post performs the Request
func (j Poster) Post(v rpc.Request) (reply rpc.Response, err error) {
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
func (j Poster) Close() (err error) {
	return nil
}

// NewPoster creates a Poster with the http.DefaultClient
func NewPoster(add string) (j Poster, err error) {
	j = Poster{
		Client: http.DefaultClient,
		URL:    add,
	}
	return
}

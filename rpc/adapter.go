package rpc

import (
	"encoding/json"

	"github.com/oliverpool/argo"
)

// Adapter fulfills the argo.Caller interface, while allowing to choose a custom Poster and a secret tocken
type Adapter struct {
	Poster Poster
	Secret string
}

// Call the given RPC method and unmarshall the response in reply.
func (a Adapter) Call(method string, reply interface{}, options ...interface{}) error {
	r := a.prepareRequest(method, options...)

	resp, err := a.Poster.Post(r)

	if err != nil {
		return err
	}

	if resp.Error.Code != 0 {
		return resp.Error
	}

	err = json.Unmarshal(resp.Result, &reply)

	return err
}

// CallWithID is similar to Call, but write the returned id as well.
func (a Adapter) CallWithID(method string, reply interface{}, id *string, options ...interface{}) error {
	r := a.prepareRequest(method, options...)

	resp, err := a.Poster.Post(r)

	if err != nil {
		return err
	}

	if resp.Error.Code != 0 {
		return resp.Error
	}

	*id = resp.ID
	err = json.Unmarshal(resp.Result, &reply)

	return err
}

func (a Adapter) prepareRequest(method string, options ...interface{}) Request {
	params := make([]interface{}, 0, len(options))

	// secret must be first
	if a.Secret != "" {
		params = append(params, "token:"+a.Secret)
	}

	// id has a separate field
	id := ""
	// position must be last
	var position int
	var hasPosition bool

	for _, p := range options {
		if o, ok := p.(argo.Option); ok {
			id = o.GetID()
			position, hasPosition = o.GetPosition()
		}
		params = append(params, p)
	}
	if hasPosition {
		params = append(params, position)
	}
	return Request{
		Version: "2.0",
		Method:  method,
		Params:  params,
		ID:      id,
	}
}

// Close gracefully closes the Poster
func (a Adapter) Close() (err error) {
	return a.Poster.Close()
}

// Adapt a Poster and a secret to fulfill the argo.Caller interface
func Adapt(c Poster, secret string) Adapter {
	return Adapter{
		Poster: c,
		Secret: secret,
	}
}

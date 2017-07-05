package rpc

import "github.com/oliverpool/argo"

type Adapter struct {
	Poster Poster
	Secret string
}

func (a Adapter) Call(method string, options ...interface{}) (argo.Response, error) {
	params := make([]interface{}, 0, len(options))

	// secret must be first
	if a.Secret != "" {
		params = append(params, "token:"+a.Secret)
	}

	// extract evcentual ID, insert the rest
	id := ""
	for _, p := range options {
		if i, ok := p.(argo.ID); ok {
			id = string(i)
		} else {
			params = append(params, p)
		}
	}
	r := Request{
		Version: "2.0",
		Method:  method,
		Params:  params,
		ID:      id,
	}
	resp, err := a.Poster.Post(r)

	if err == nil && resp.Error.Code != 0 {
		err = resp.Error
	}
	return resp.Response, err
}

func Adapt(c Poster, secret string) Adapter {
	return Adapter{
		Poster: c,
		Secret: secret,
	}
}
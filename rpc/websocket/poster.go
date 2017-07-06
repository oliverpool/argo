package websocket

import "github.com/oliverpool/argo/rpc"

// Poster allows to send rpc.Request via Websocket
type Poster struct {
	*Websocket
}

// Post performs the Request
func (j Poster) Post(v rpc.Request) (reply rpc.Response, err error) {
	reply, err = j.post(v)
	err = rpc.ConvertClosedNetworkConnectionError(err)
	return
}

func (j Poster) post(v rpc.Request) (reply rpc.Response, err error) {
	err = j.Conn.WriteJSON(&v)

	for reply.Response.GID == "" && err == nil {
		// The first response might not be the aknowledgement (but onDownloadStart for instance)
		err = j.Conn.ReadJSON(&reply)
	}
	return
}

// NewPoster creates a Poster with the websocket.DefaultDialer
func NewPoster(add string) (j *Poster, err error) {
	w, err := NewWebsocket(add)
	j = &Poster{
		w,
	}
	return
}

// NewAdapter creates a Caller with the websocket.DefaultDialer
func NewAdapter(address, secret string) (a rpc.Adapter, err error) {
	p, err := NewPoster(address)
	a = rpc.Adapt(p, secret)
	return
}

package websocket

import (
	"github.com/oliverpool/argo"
	"github.com/oliverpool/argo/rpc"
)

// Caller allows to send RPCCalls to an URL
type Caller struct {
	*Websocket
}

// Call performs the RPCRequest
func (j Caller) Call(v rpc.Request) (reply rpc.Response, err error) {
	err = j.send(v)
	if err != nil {
		return
	}
	err = j.receive(&reply)
	return
}

func (j Caller) send(v rpc.Request) error {
	if j.IsClosed() {
		return argo.ErrConnIsClosed
	}
	return j.Conn.WriteJSON(&v)
}

func (j Caller) receive(reply *rpc.Response) error {
	if j.IsClosed() {
		return argo.ErrConnIsClosed
	}
	return j.Conn.ReadJSON(reply)
}

// NewCaller creates a Caller with the websocket.DefaultDialer
func NewCaller(add string) (j *Caller, err error) {
	w, err := NewWebsocket(add)
	j = &Caller{
		w,
	}
	return
}

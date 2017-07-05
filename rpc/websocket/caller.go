package websocket

import "github.com/oliverpool/argo/rpc"

// Caller allows to send RPCCalls to an URL
type Caller struct {
	*Websocket
}

// Call performs the RPCRequest
func (j Caller) Call(v rpc.Request) (reply rpc.Response, err error) {
	reply, err = j.call(v)
	err = rpc.ConvertClosedNetworkConnectionError(err)
	return
}

func (j Caller) call(v rpc.Request) (reply rpc.Response, err error) {
	err = j.Conn.WriteJSON(&v)
	if err != nil {
		return
	}
	err = j.Conn.ReadJSON(&reply)
	return
}

// NewCaller creates a Caller with the websocket.DefaultDialer
func NewCaller(add string) (j *Caller, err error) {
	w, err := NewWebsocket(add)
	j = &Caller{
		w,
	}
	return
}

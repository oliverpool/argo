package websocket

import (
	"github.com/oliverpool/argo"
	"github.com/oliverpool/argo/rpc"
)

// Receiver allows to receive notifications from an URL
type Receiver struct {
	*Websocket
}

// Receive a notification
func (j Receiver) Receive() (notif argo.Notification, err error) {
	return j.receive()
}

func (j Receiver) receive() (notif rpc.Notification, err error) {
	err = j.Conn.ReadJSON(&notif)
	err = rpc.ConvertClosedNetworkConnectionError(err)
	return
}

// NewReceiver creates a Receiver with the websocket.DefaultDialer
func NewReceiver(add string) (j *Receiver, err error) {
	w, err := NewWebsocket(add)
	j = &Receiver{
		w,
	}
	return
}

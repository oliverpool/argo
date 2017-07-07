package websocket

import (
	"github.com/oliverpool/argo"
	"github.com/oliverpool/argo/rpc"
)

// Emitter allows to receive notifications from an URL
type Emitter struct {
	*Websocket
}

// Emit a notification (as soon as one is read via websocket)
func (j Emitter) Emit() (notif argo.Notification, err error) {
	return j.emit()
}

func (j Emitter) emit() (notif rpc.Notification, err error) {
	err = j.Conn.ReadJSON(&notif)
	err = rpc.ConvertClosedNetworkConnectionError(err)
	return
}

// NewEmitter creates a NotificationEmitter with the websocket.DefaultDialer
func NewEmitter(add string) (j *Emitter, err error) {
	w, err := NewWebsocket(add)
	j = &Emitter{
		w,
	}
	return
}

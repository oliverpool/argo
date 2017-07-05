package websocket

import (
	"time"

	"github.com/gorilla/websocket"
)

// Websocket represent a closable websocket connection
type Websocket struct {
	Conn      *websocket.Conn
	WriteWait time.Duration
	isClosed  bool
}

// IsClosed indicated if the connection has been closed
func (w Websocket) IsClosed() bool {
	return w.isClosed
}

// Close gracefully closes the connection (with a CloseMessage)
func (w *Websocket) Close() (err error) {
	w.isClosed = true
	defer func() {
		berr := w.Conn.Close()
		if err == nil {
			err = berr
		}
	}()
	delay := w.WriteWait
	if delay == 0 {
		delay = time.Second
	}
	err = w.Conn.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""), time.Now().Add(w.WriteWait))
	return
}

// NewWebsocket creates a closable websocket with the websocket.DefaultDialer
func NewWebsocket(add string) (w *Websocket, err error) {
	var conn *websocket.Conn
	conn, _, err = websocket.DefaultDialer.Dial(add, nil)
	w = &Websocket{
		Conn:      conn,
		WriteWait: time.Second,
	}
	return
}

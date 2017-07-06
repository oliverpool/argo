package argo

import "fmt"

// NotificationReceiver allows to receive Notifications
type NotificationReceiver interface {
	Receive() (Notification, error)
}

// Notification from aria2
type Notification interface {
	Identifier() string
	GID() []string // GID of the downloads
}

type NotificationHandler interface {
	Started(GID []string)
	Paused(GID []string)
	Stopped(GID []string)
	Completed(GID []string)
	BtCompleted(GID []string)
	Error(GID []string)
	ReceptionError(error) bool // if true, stop receiving notifications
	OtherIdentifier(Identifier string, GID []string)
}

// Response represents a JSON-RPC response to a request
type Response struct {
	GID string `json:"result"` // GID of the download
	ID  string `json:"id"`
}

// Caller allows to perform Requests
type Caller interface {
	Call(method string, params ...interface{}) (Response, error)
	Close() error
}

// Option allows to pass custom options
//
type Option map[string]interface{}

func (o Option) getString(key string) string {
	if v, ok := o[key]; ok {
		return fmt.Sprintf("%v", v)
	}
	return ""
}

func (o Option) getInt(key string) (int, bool) {
	if v, ok := o[key]; ok {
		if s, ok := v.(int); ok {
			return s, true
		}
	}
	return 0, false
}

func (o Option) GetID() string {
	return o.getString("id")
}

func (o Option) GetPosition() (int, bool) {
	return o.getInt("position")
}

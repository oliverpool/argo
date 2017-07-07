package argo

import "fmt"

// A Client is an aria2 client (https://aria2.github.io/)
//
// It can be constructed with the http.NewClient method of the subpackage argo/rpc/http
type Client struct {
	Caller Caller
}

// Close gracefully closes the Caller
func (c Client) Close() (err error) {
	return c.Caller.Close()
}

func (c Client) mergeOptions(options ...Option) Option {
	return mergeOptions(options...)
}

func mergeOptions(options ...Option) Option {
	opt := Option{}
	for _, o := range options {
		for k, v := range o {
			opt[k] = v
		}
	}
	return opt
}

// NotificationReceiver allows to receive Notifications
//
// It can be constructed with the websocket.NewReceiver method of the subpackage argo/rpc/websocket
type NotificationReceiver interface {
	Receive() (Notification, error)
}

// Notification from aria2
type Notification interface {
	Identifier() string
	GID() []string // GID of the downloads
}

// NotificationHandler can handle notifications
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

// Caller allows to perform Requests
type Caller interface {
	Call(method string, reply interface{}, params ...interface{}) error
	CallWithID(method string, reply interface{}, id *string, params ...interface{}) error
	Close() error
}

// Option allows to pass custom options
//
// See the option subpackage
type Option map[string]interface{}

// GetID returns the "id" value if present (empty string otherwise)
func (o Option) GetID() string {
	return o.getString("id")
}

func (o Option) getString(key string) string {
	if v, ok := o[key]; ok {
		return fmt.Sprintf("%v", v)
	}
	return ""
}

// GetPosition returns the "position" value if present
func (o Option) GetPosition() (int, bool) {
	return o.getInt("position")
}

func (o Option) getInt(key string) (int, bool) {
	if v, ok := o[key]; ok {
		if s, ok := v.(int); ok {
			return s, true
		}
	}
	return 0, false
}

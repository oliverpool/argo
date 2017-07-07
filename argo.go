package argo

import "fmt"

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

// Caller allows to perform Requests
type Caller interface {
	Call(method string, reply interface{}, params ...interface{}) error
	CallWithID(method string, reply interface{}, id *string, params ...interface{}) error
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

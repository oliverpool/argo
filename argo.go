package argo

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

// ID allows to identify a request
type ID string

// Caller allows to perform Requests
type Caller interface {
	Call(method string, params ...interface{}) (Response, error)
	Close() error
}

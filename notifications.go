package argo

// Notification from aria2
type Notification interface {
	Identifier() string
	GID() []GID // GID of the downloads
}

// NotificationEmitter emits Notifications
//
// It can be constructed with the websocket.NewEmitter method of the subpackage argo/rpc/websocket
type NotificationEmitter interface {
	Emit() (Notification, error)
}

// NotificationHandler must handle (at least) notifications errors (during communication or decoding).
// It may also handle some of the other notifications of the NotificationFullHandler
//
// It can be used with the notification.Forward method of the argo/notification subpackage
type NotificationHandler interface {
	ReceptionError(error) bool // if true, stop receiving notifications
}

// NotificationFullHandler can handle all notifications
type NotificationFullHandler interface {
	NotificationHandler
	Started(GID []GID)
	Paused(GID []GID)
	Stopped(GID []GID)
	Completed(GID []GID)
	BtCompleted(GID []GID)
	Error(GID []GID)
	OtherIdentifier(Identifier string, GID []GID)
}

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

// NotificationHandler can handle notifications
type NotificationHandler interface {
	Started(GID []GID)
	Paused(GID []GID)
	Stopped(GID []GID)
	Completed(GID []GID)
	BtCompleted(GID []GID)
	Error(GID []GID)
	ReceptionError(error) bool // if true, stop receiving notifications
	OtherIdentifier(Identifier string, GID []GID)
}

// ForwardNotifications forwads the notifications from the emitter to the handler
func ForwardNotifications(conn NotificationEmitter, h NotificationHandler) error {
	for {
		notification, err := conn.Emit()
		if err != nil {
			if h.ReceptionError(err) {
				return err
			}
			continue
		}
		gid := notification.GID()
		switch notification.Identifier() {
		case "aria2.onDownloadStart":
			h.Started(gid)
		case "aria2.onDownloadPause":
			h.Paused(gid)
		case "aria2.onDownloadStop":
			h.Stopped(gid)
		case "aria2.onDownloadComplete":
			h.Completed(gid)
		case "aria2.onDownloadError":
			h.Error(gid)
		case "aria2.onBtDownloadComplete":
			h.BtCompleted(gid)
		default:
			h.OtherIdentifier(notification.Identifier(), gid)
		}
	}
}

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

func GetNotifications(conn NotificationReceiver, h NotificationHandler) error {
	for {
		notification, err := conn.Receive()
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

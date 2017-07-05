package argo

import "log"

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
}

func GetNotifications(conn NotificationReceiver, n NotificationHandler) (err error) {
	for {
		notification, err := conn.Receive()
		if err != nil {
			if err == ErrConnIsClosed {
				return nil
			}
			log.Println("reading ws:", err)
			continue
		}
		gid := notification.GID()
		switch notification.Identifier() {
		case "aria2.onDownloadStart":
			n.Started(gid)
		case "aria2.onDownloadPause":
			n.Paused(gid)
		case "aria2.onDownloadStop":
			n.Stopped(gid)
		case "aria2.onDownloadComplete":
			n.Completed(gid)
		case "aria2.onDownloadError":
			n.Error(gid)
		case "aria2.onBtDownloadComplete":
			n.BtCompleted(gid)
		default:
			log.Printf("unexpected notification: %s: %#v\n", notification.Identifier(), notification)
		}
	}
}

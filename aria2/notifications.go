package aria2

import "github.com/oliverpool/argo"

// ForwardNotifications forwads the notifications from the receiver to the handler
func ForwardNotifications(conn argo.NotificationReceiver, h argo.NotificationHandler) error {
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
		case onDownloadStart:
			h.Started(gid)
		case onDownloadPause:
			h.Paused(gid)
		case onDownloadStop:
			h.Stopped(gid)
		case onDownloadComplete:
			h.Completed(gid)
		case onDownloadError:
			h.Error(gid)
		case onBtDownloadComplete:
			h.BtCompleted(gid)
		default:
			h.OtherIdentifier(notification.Identifier(), gid)
		}
	}
}

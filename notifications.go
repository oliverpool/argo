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
// It may also handle some of the other notifications of the FullNotificationHandler
type NotificationHandler interface {
	ReceptionError(error) bool // if true, stop receiving notifications
}

// FullNotificationHandler can handle all notifications
type FullNotificationHandler interface {
	NotificationHandler
	Started(GID []GID)
	Paused(GID []GID)
	Stopped(GID []GID)
	Completed(GID []GID)
	BtCompleted(GID []GID)
	Error(GID []GID)
	OtherIdentifier(Identifier string, GID []GID)
}

// ForwardNotifications forwads the notifications from the emitter to the handler
func ForwardNotifications(conn NotificationEmitter, p NotificationHandler) error {
	h := completeNotificationHandler(p)
	for {
		notification, err := conn.Emit()
		if err != nil {
			if h.ReceptionError(err) {
				return err
			}
		} else {
			h.handle(notification.Identifier(), notification.GID())
		}
	}
}

// NotificationReceptionErrorHandler handles notification errors (during communication or decoding)
type NotificationReceptionErrorHandler interface {
	ReceptionError(error) bool // if this returns true, stop receiving notifications
}

// NotificationStartedHandler handles onDownloadStart
type NotificationStartedHandler interface {
	Started([]GID)
}

// NotificationPausedHandler handles onDownloadPause
type NotificationPausedHandler interface {
	Paused([]GID)
}

// NotificationStoppedHandler handles onDownloadStop
type NotificationStoppedHandler interface {
	Stopped([]GID)
}

// NotificationCompletedHandler handles onDownloadComplete
type NotificationCompletedHandler interface {
	Completed([]GID)
}

// NotificationBtCompletedHandler handles onBtDownloadComplete
type NotificationBtCompletedHandler interface {
	BtCompleted([]GID)
}

// NotificationErrorHandler handles onDownloadError
type NotificationErrorHandler interface {
	Error([]GID)
}

// NotificationUnknownHandler handles unknown notifications
type NotificationUnknownHandler interface {
	Unknown(identifier string, gid []GID)
}

type fullNotificationHandler struct {
	ReceptionError func(error) bool // if this returns true, stop receiving notifications
	Started        func([]GID)
	Paused         func([]GID)
	Stopped        func([]GID)
	Completed      func([]GID)
	BtCompleted    func([]GID)
	Error          func([]GID)
	Unknown        func(identifier string, gid []GID)
}

func (f fullNotificationHandler) handle(ident string, gid []GID) {
	switch ident {
	case "aria2.onDownloadStart":
		f.Started(gid)
	case "aria2.onDownloadPause":
		f.Paused(gid)
	case "aria2.onDownloadStop":
		f.Stopped(gid)
	case "aria2.onDownloadComplete":
		f.Completed(gid)
	case "aria2.onDownloadError":
		f.Error(gid)
	case "aria2.onBtDownloadComplete":
		f.BtCompleted(gid)
	default:
		f.Unknown(ident, gid)
	}
}

func completeNotificationHandler(h NotificationReceptionErrorHandler) (f fullNotificationHandler) {
	emptyFunc := func([]GID) {}
	f.ReceptionError = h.ReceptionError

	if v, ok := h.(NotificationStartedHandler); ok {
		f.Started = v.Started
	} else {
		f.Started = emptyFunc
	}

	if v, ok := h.(NotificationPausedHandler); ok {
		f.Paused = v.Paused
	} else {
		f.Paused = emptyFunc
	}

	if v, ok := h.(NotificationStoppedHandler); ok {
		f.Stopped = v.Stopped
	} else {
		f.Stopped = emptyFunc
	}

	if v, ok := h.(NotificationCompletedHandler); ok {
		f.Completed = v.Completed
	} else {
		f.Completed = emptyFunc
	}

	if v, ok := h.(NotificationErrorHandler); ok {
		f.Error = v.Error
	} else {
		f.Error = emptyFunc
	}

	if v, ok := h.(NotificationBtCompletedHandler); ok {
		f.BtCompleted = v.BtCompleted
	} else {
		f.BtCompleted = emptyFunc
	}

	if v, ok := h.(NotificationUnknownHandler); ok {
		f.Unknown = v.Unknown
	} else {
		f.Unknown = func(string, []GID) {}
	}
	return
}

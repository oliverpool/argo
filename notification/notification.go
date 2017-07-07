// Package notification allows to forward notifications from an Emitter to a Handler
package notification

import "github.com/oliverpool/argo"

// Forward the notifications from the emitter to the handler
func Forward(conn argo.NotificationEmitter, p argo.NotificationHandler) error {
	h := completeHandler(p)
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

// StartedHandler handles onDownloadStart
type StartedHandler interface {
	Started([]argo.GID)
}

// PausedHandler handles onDownloadPause
type PausedHandler interface {
	Paused([]argo.GID)
}

// StoppedHandler handles onDownloadStop
type StoppedHandler interface {
	Stopped([]argo.GID)
}

// CompletedHandler handles onDownloadComplete
type CompletedHandler interface {
	Completed([]argo.GID)
}

// BtCompletedHandler handles onBtDownloadComplete
type BtCompletedHandler interface {
	BtCompleted([]argo.GID)
}

// ErrorHandler handles onDownloadError
type ErrorHandler interface {
	Error([]argo.GID)
}

// UnknownHandler handles unknown notifications
type UnknownHandler interface {
	Unknown(identifier string, gid []argo.GID)
}

type fullHandler struct {
	ReceptionError func(error) bool // if this returns true, stop receiving notifications
	Started        func([]argo.GID)
	Paused         func([]argo.GID)
	Stopped        func([]argo.GID)
	Completed      func([]argo.GID)
	BtCompleted    func([]argo.GID)
	Error          func([]argo.GID)
	Unknown        func(identifier string, gid []argo.GID)
}

func (f fullHandler) handle(ident string, gid []argo.GID) {
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

func completeHandler(h argo.NotificationHandler) (f fullHandler) {
	emptyFunc := func([]argo.GID) {}
	f.ReceptionError = h.ReceptionError

	if v, ok := h.(StartedHandler); ok {
		f.Started = v.Started
	} else {
		f.Started = emptyFunc
	}

	if v, ok := h.(PausedHandler); ok {
		f.Paused = v.Paused
	} else {
		f.Paused = emptyFunc
	}

	if v, ok := h.(StoppedHandler); ok {
		f.Stopped = v.Stopped
	} else {
		f.Stopped = emptyFunc
	}

	if v, ok := h.(CompletedHandler); ok {
		f.Completed = v.Completed
	} else {
		f.Completed = emptyFunc
	}

	if v, ok := h.(ErrorHandler); ok {
		f.Error = v.Error
	} else {
		f.Error = emptyFunc
	}

	if v, ok := h.(BtCompletedHandler); ok {
		f.BtCompleted = v.BtCompleted
	} else {
		f.BtCompleted = emptyFunc
	}

	if v, ok := h.(UnknownHandler); ok {
		f.Unknown = v.Unknown
	} else {
		f.Unknown = func(string, []argo.GID) {}
	}
	return
}

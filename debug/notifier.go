package debug

import (
	"log"

	"net"

	"github.com/oliverpool/argo"
)

// NotificationLogger prints the notification to the logger
type NotificationLogger struct {
	Logger *log.Logger
}

var _ argo.NotificationFullHandler = NotificationLogger{}

// Started handles the corresponding event
func (d NotificationLogger) Started(GID []argo.GID) {
	d.Logger.Printf("%s started.\n", GID)
}

// Paused handles the corresponding event
func (d NotificationLogger) Paused(GID []argo.GID) {
	d.Logger.Printf("%s paused.\n", GID)
}

// Stopped handles the corresponding event
func (d NotificationLogger) Stopped(GID []argo.GID) {
	d.Logger.Printf("%s stopped.\n", GID)
}

// Completed handles the corresponding event
func (d NotificationLogger) Completed(GID []argo.GID) {
	d.Logger.Printf("%s completed.\n", GID)
}

// Error handles the corresponding event
func (d NotificationLogger) Error(GID []argo.GID) {
	d.Logger.Printf("%s error.\n", GID)
}

// BtCompleted handles the corresponding event
func (d NotificationLogger) BtCompleted(GID []argo.GID) {
	d.Logger.Printf("bt %s completed.\n", GID)
}

// Unknown handles the corresponding event
func (d NotificationLogger) Unknown(ident string, GID []argo.GID) {
	d.Logger.Printf("unknown %s for %s.\n", ident, GID)
}

// ReceptionError handles the corresponding event
func (d NotificationLogger) ReceptionError(err error) bool {
	if err == argo.ErrConnIsClosed {
		d.Logger.Printf("connection closed.\n")
		return true
	}
	d.Logger.Printf("notification error %#v.\n", err)
	if e, ok := err.(*net.OpError); ok {
		d.Logger.Printf("NetOp %s.\n", e.Err.Error())

	}
	return false
}

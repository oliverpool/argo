package debug

import (
	"log"

	"net"

	"github.com/oliverpool/argo"
)

type NotificationReceiver struct {
	Logger *log.Logger
}

func (d NotificationReceiver) Started(GID []argo.GID) {
	d.Logger.Printf("%s started.\n", GID)
}
func (d NotificationReceiver) Paused(GID []argo.GID) {
	d.Logger.Printf("%s paused.\n", GID)
}
func (d NotificationReceiver) Stopped(GID []argo.GID) {
	d.Logger.Printf("%s stopped.\n", GID)
}
func (d NotificationReceiver) Completed(GID []argo.GID) {
	d.Logger.Printf("%s completed.\n", GID)
}
func (d NotificationReceiver) Error(GID []argo.GID) {
	d.Logger.Printf("%s error.\n", GID)
}
func (d NotificationReceiver) BtCompleted(GID []argo.GID) {
	d.Logger.Printf("bt %s completed.\n", GID)
}
func (d NotificationReceiver) OtherIdentifier(ident string, GID []argo.GID) {
	d.Logger.Printf("unknown %s for %s.\n", ident, GID)
}
func (d NotificationReceiver) ReceptionError(err error) bool {
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

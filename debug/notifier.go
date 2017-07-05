package debug

import "log"

type NotificationReceiver struct {
	Logger *log.Logger
}

func (d NotificationReceiver) Started(GID []string) {
	d.Logger.Printf("%s started.\n", GID)
}
func (d NotificationReceiver) Paused(GID []string) {
	d.Logger.Printf("%s paused.\n", GID)
}
func (d NotificationReceiver) Stopped(GID []string) {
	d.Logger.Printf("%s stopped.\n", GID)
}
func (d NotificationReceiver) Completed(GID []string) {
	d.Logger.Printf("%s completed.\n", GID)
}
func (d NotificationReceiver) Error(GID []string) {
	d.Logger.Printf("%s error.\n", GID)
}
func (d NotificationReceiver) BtCompleted(GID []string) {
	d.Logger.Printf("bt %s completed.\n", GID)
}

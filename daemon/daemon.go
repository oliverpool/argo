// Package daemon simplifies the configuration of the aria2c daemon
package daemon

import (
	"net"
	"os/exec"
)

// Aria2 stores the config to launch the daemon
type Aria2 struct {
	Name string
	args []string
}

// New creates a default daemon configuration
func New() Aria2 {
	a := Aria2{
		Name: "aria2c",
	}
	a.Option(EnableRPC, LogLevel("warn"))
	return a
}

// Cmd return the exec.Cmd to launch the aria2c daemon
func (a Aria2) Cmd() *exec.Cmd {
	return exec.Command(a.Name, a.args...)
}

// IsRunningOn tests if the address is listening for TCP connections
func IsRunningOn(address string) bool {
	conn, err := net.Dial("tcp", address)
	if conn != nil {
		conn.Close()
	}
	return err == nil
}

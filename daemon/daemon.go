package daemon

import (
	"net"
	"os/exec"
)

type Aria2 struct {
	Name string
	args []string
}

func New() Aria2 {
	a := Aria2{
		Name: "aria2c",
	}
	a.Option(EnableRPC, Log("warn"))
	return a
}

func (a Aria2) Cmd() *exec.Cmd {
	return exec.Command(a.Name, a.args...)
}

type RunError struct {
	error
	CombinedOutput []byte
}

func IsRunningOn(address string) bool {
	conn, err := net.Dial("tcp", address)
	if conn != nil {
		conn.Close()
	}
	return err == nil
}

/*

// LaunchAria2cDaemon launchs aria2 daemon to listen for RPC calls, locally.
func (id *client) LaunchAria2cDaemon() (info VersionInfo, err error) {
	if info, err = id.GetVersion(); err == nil {
		return
	}
	args := []string{"--enable-rpc", "--rpc-listen-all"}
	if id.token != "" {
		args = append(args, "--rpc-secret="+id.token)
	}
	cmd := exec.Command("aria2c", args...)
	if err = cmd.Start(); err != nil {
		return
	}
	cmd.Process.Release()
	timeout := false
	timer := time.AfterFunc(time.Second, func() {
		timeout = true
	})
	for !timeout {
		if info, err = id.GetVersion(); err == nil {
			break
		}
	}
	timer.Stop()
	return
}
*/

package daemon_test

import (
	"fmt"
	"time"

	"github.com/oliverpool/argo"
	"github.com/oliverpool/argo/daemon"
)

func Example() {
	fmt.Println("Creating daemon configuration")
	aria2 := daemon.New()
	aria2.Option(daemon.Port("6800"), daemon.Secret("secretToken"), argo.Option{"max-concurrent-downloads": 1})

	fmt.Println("Launching daemon (in a goroutine)")
	cmd := aria2.Cmd()

	done := make(chan struct{})
	go func() {
		out, err := cmd.CombinedOutput()
		fmt.Println("Daemon stopped")
		fmt.Printf("[%s: %s]\n", err, string(out))
		close(done)
	}()

	fmt.Println("Wait until the adresse is ready to listen")
	for !daemon.IsRunningOn(":6800") {
		time.Sleep(time.Second)
	}

	fmt.Println("Killing daemon (not very nice... prefer to send a aria2.shutdown command)")
	cmd.Process.Kill()
	<-done

	fmt.Println("Bye")

	// Output:
	// Creating daemon configuration
	// Launching daemon (in a goroutine)
	// Wait until the adresse is ready to listen
	// Killing daemon (not very nice... prefer to send a aria2.shutdown command)
	// Daemon stopped
	// [signal: killed: ]
	// Bye
}

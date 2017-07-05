package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	"strings"

	"github.com/oliverpool/argo"
	"github.com/oliverpool/argo/aria2"
	"github.com/oliverpool/argo/daemon"
	"github.com/oliverpool/argo/debug"
	"github.com/oliverpool/argo/rpc"
	"github.com/oliverpool/argo/rpc/http"
	"github.com/oliverpool/argo/rpc/websocket"
)

var (
	rpcSecret = flag.String("secret", "", "set --rpc-secret for aria2c")
	ariaURL   = flag.String("uri", "ws://localhost:6800/jsonrpc", "URL of the aria websocket endpoint")
)

func main() {
	flag.Parse()

	var err error
	var j2, jhttp rpc.Poster

	aria := daemon.New()
	secret := "secret"
	aria.Option(daemon.Port("6800"), daemon.Secret(secret))
	go func() {
		out, err := aria.Cmd().CombinedOutput()
		if err != nil {
			log.Printf("%s: %s", err, string(out))
		}
	}()

	for !daemon.IsRunningOn(":6800") {
		time.Sleep(time.Second)
	}
	httpURL := "http" + strings.TrimLeft(*ariaURL, "ws")

	notifier := debug.NotificationReceiver{
		log.New(os.Stderr, "", log.LstdFlags),
	}

	j, err := websocket.NewReceiver(*ariaURL)
	if err != nil {
		log.Fatal(err)
	}

	j2, err = websocket.NewPoster(*ariaURL)
	if err != nil {
		log.Fatal(err)
	}
	j2r := rpc.Adapt(j2, secret)

	jhttp, err = http.NewPoster(httpURL)
	if err != nil {
		log.Fatal(err)
	}
	jhttpr := rpc.Adapt(jhttp, secret)

	ctx := context.Background()
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(3*time.Second))
	defer cancel()

	go func() {

		uri := []string{"https://static.ranchcomputing.com/css/topbar.1.css"}

		reply, err := jhttpr.Call("aria2.addUri", uri)
		if err != nil {
			panic(err)
		}
		log.Printf("%#v", reply)
		reply, err = j2r.Call("aria2.addUri", uri, argo.ID("-1"))
		if err != nil {
			panic(err)
		}
		log.Printf("%#v", reply)
		<-ctx.Done()
		j.Close()
		j2.Close()
	}()

	aria2.ForwardNotifications(j, notifier)

	reply, err := jhttpr.Call("aria2.shutdown")
	if err != nil {
		panic(err)
	}
	log.Printf("%#v", reply)

	for daemon.IsRunningOn(":6800") {
		time.Sleep(time.Second)
	}
	log.Printf("Bye")

	/*


		var err error
		rpcc, err = rpc.New(*rpcURI, *rpcSecret)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(2)
		}
		if flag.NArg() == 0 {
			fmt.Fprintf(os.Stderr, "usage: argo {CMD} {PARAMETERS}...\n")
			flag.PrintDefaults()
			os.Exit(1)
		}
		args := flag.Args()
		if cmd, ok := cmds[args[0]]; ok {
			err = cmd(args[1:]...)
		} else {
			err = errInvalidCmd
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	*/
}

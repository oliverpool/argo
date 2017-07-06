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
	"github.com/oliverpool/argo/option"
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
	var j2, jhttpr argo.Caller

	aria := daemon.New()
	secret := "secret"
	aria.Option(daemon.Port("6800"), daemon.Secret(secret), daemon.AppendArg("--max-concurrent-downloads=1"))
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

	j2, err = websocket.NewAdapter(*ariaURL, secret)
	if err != nil {
		log.Fatal(err)
	}

	jhttpr = http.NewAdapter(httpURL, secret)

	ctx := context.Background()
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(3*time.Second))
	defer cancel()

	dir := option.Dir("/tmp/aria")

	dir0 := argo.Option{"id": "123", "dir": "/tmp/aria", "position": 0}

	// iid := argo.Option{"id": "123", "dir": "/tmp/aria", "position": 0}

	/*
		if o, ok := p.(argo.Option); ok {
			if i, ok := o["ID"]; ok {
				id = string(i)
			}
	*/

	uri1 := []string{"https://static.ranchcomputing.com/css/topbar.1.css"}
	uri2 := []string{"https://static.ranchcomputing.com/css/topbar.1.css"}
	uri3 := []string{"https://static.ranchcomputing.com/css/topbar.1.css"}

	go func() {

		//uri := []string{"https://static.ranchcomputing.com/css/topbar.1.css"}

		reply, err := jhttpr.Call("aria2.addUri", uri1, dir)
		if err != nil {
			log.Printf("%#v", err)
			panic(err)
		}
		log.Printf("%#v", reply)

		reply, err = jhttpr.Call("aria2.addUri", uri2, dir)
		if err != nil {
			log.Printf("%#v", err)
			panic(err)
		}
		log.Printf("%#v", reply)

		reply, err = jhttpr.Call("aria2.addUri", uri3, dir0)
		if err != nil {
			log.Printf("%#v", err)
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

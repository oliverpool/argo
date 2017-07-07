package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	"strings"

	"io/ioutil"

	"github.com/oliverpool/argo"
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
	var j2, jhttpr argo.Client

	aria := daemon.New()
	secret := "secret"
	aria.Option(daemon.Port("6800"), daemon.Secret(secret), argo.Option{"max-concurrent-downloads": 1})
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

	j, err := websocket.NewEmitter(*ariaURL)
	if err != nil {
		log.Fatal(err)
	}

	j2, err = websocket.NewClient(*ariaURL, secret)
	if err != nil {
		log.Fatal(err)
	}

	jhttpr = http.NewClient(httpURL, secret)

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

	uri1 := []string{"http://bitlove.org/nitramred/staatsbuergerkunde-opus/SBK064_Mauerweg.opus.torrent"}
	uri2 := []string{"https://static.ranchcomputing.com/css/topbar.1.css"}
	uri3 := []string{"https://static.ranchcomputing.com/css/topbar.1.css"}

	_ = uri1

	torrentContent, err := ioutil.ReadFile("/home/olivier/Downloads/temp/SBK064_Mauerweg.opus.torrent")
	if err != nil {
		panic(err)
	}
	_ = torrentContent

	go func() {

		//uri := []string{"https://static.ranchcomputing.com/css/topbar.1.css"}

		//*
		reply, err := jhttpr.AddURI(uri1, dir, argo.Option{"id": "torrent"})
		/*/
		reply, err := jhttpr.AddTorrent(torrentContent, dir)

		// time.Sleep(15 * time.Second)
		//*/
		if err != nil {
			log.Printf("%#v", err)
			panic(err)
		}
		log.Printf("%#v", reply)

		reply, err = jhttpr.AddURI(uri2, dir, argo.Option{"id": 42})
		if err != nil {
			log.Printf("%#v", err)
			panic(err)
		}
		log.Printf("%#v", reply)

		reply, err = jhttpr.AddURI(uri3, dir0)
		if err != nil {
			log.Printf("%#v", err)
			panic(err)
		}
		log.Printf("%#v", reply)
		g := reply.GID

		time.Sleep(1000 * time.Millisecond)

		//*

		_ = g
		jhttpr.ChangeOption(g, argo.Option{"max-download-limit": "20K"})
		reply2, err := jhttpr.ListNotifications()
		if err != nil {
			log.Printf("%#v", err)
		}
		log.Printf("%#v", reply2)

		//*/

		<-ctx.Done()
		j.Close()
		j2.Close()
	}()

	argo.ForwardNotifications(j, notifier)

	reply, err := jhttpr.Shutdown()
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

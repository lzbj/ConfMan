package main

import (
	"fmt"
	serv "github.com/lzbj/ConfMan/server"
	cf "github.com/lzbj/ConfMan/server/conf"
	"google.golang.org/grpc"

	"flag"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	grpcServer := grpc.NewServer()
	// Get backend type
	betype := cf.Cfg.GetString("backend.type")

	// Get backend address
	beaddress := cf.Cfg.GetString("backend.address")
	ConfServer := serv.NewConfManServer(grpcServer)
	err := serv.PersistInit(betype, beaddress)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	port, _ := strconv.Atoi(cf.Cfg.GetString("server.port"))
	go func() {
		err = ConfServer.Serve(port)
		if err != nil {
			fmt.Sprintf("error happened : %s", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

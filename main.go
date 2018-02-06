package main

import (
	"fmt"
	serv "github.com/lzbj/ConfMan/server"
	cf "github.com/lzbj/ConfMan/server/conf"
	"google.golang.org/grpc"

	"os"
	"strconv"
)

func main() {
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
	err = ConfServer.Serve(port)
	if err != nil {
		fmt.Sprintf("error happened : %s", err)
		os.Exit(1)
	}

}

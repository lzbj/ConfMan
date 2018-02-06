package main

import (
	"log"

	cf "github.com/lzbj/ConfMan/proto/ConfMan"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:8081"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c2 := cf.NewConfManClient(conn)

	r2, err := c2.GetConf(context.Background(), &cf.GetConfRequest{ServiceName: "invoice", HashKey: "key2"})
	if err != nil {
		log.Fatalf("could not get value: %v", err)
	}
	log.Printf("service : %s", r2.HashValue)
}

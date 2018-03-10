package main

import (
	"log"

	"fmt"
	cf "github.com/lzbj/ConfMan/proto/ConfMan"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"os"
	"os/signal"
	"time"
)

const (
	address = "localhost:8081"
)

var (
	c2                                        cf.ConfManClient
	serviceName, hashkey, hashvalue, newvalue string
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c2 = cf.NewConfManClient(conn)
	serviceName = "invoice"
	hashkey = "dbconnection"
	hashvalue = "localhost:3306"
	newvalue = "localhost:3307"

	fmt.Println("##########Create kv########")
	create(serviceName, hashkey, hashvalue)

	fmt.Println("##########Get kv########")
	get(serviceName, hashkey)

	fmt.Println("##########Delete kv#########")
	delete(serviceName, hashkey)

	fmt.Println("##########Get kv########")
	get(serviceName, hashkey)

	fmt.Println("Simulate etcd watch for key update")
	c1 := make(chan struct{})
	c2 := make(chan os.Signal, 1)
	signal.Notify(c2, os.Interrupt)
	go func() {
		fmt.Println("##########First we create kv########")
		create(serviceName, hashkey, hashvalue)

		fmt.Println("##########Then we update kv########")
		create(serviceName, hashkey, newvalue)
		time.Sleep(5 * time.Second)
		c1 <- struct{}{}
	}()

	go func() {
		for {
			time.Sleep(1 * time.Second)
			select {
			case _ = <-c1:

				fmt.Println("##########We get updated kv########")
				get(serviceName, hashkey)
				c2 <- os.Interrupt
			default:
				fmt.Println("##########wating for kv update#####")
			}
		}
	}()
	<-c2

}

func create(serviceName, key, value string) error {
	ur := &cf.UpdateConfRequest{ServiceName: serviceName, HashKey: key, HashValue: value}
	cm, err := c2.CreateConf(context.Background(), ur)
	if err != nil {
		log.Printf("could not create kv: %s", err)
	}
	log.Printf("created service %s, key %s, value %s", cm.ServiceName, cm.HashKey, cm.HashValue)
	return err
}

func get(service, key string) error {
	gr := &cf.GetConfRequest{ServiceName: service, HashKey: key}
	cm, err := c2.GetConf(context.Background(), gr)
	if err != nil {
		log.Printf("could not get service %s with key %s", gr.ServiceName, gr.HashKey)
	} else {
		log.Printf("get service %s, key %s, value %s", cm.ServiceName, cm.HashKey, cm.HashValue)

	}

	return err
}

func delete(service, key string) error {
	dr := &cf.DeleteConfRequest{ServiceName: service, HashKey: key}
	rep, err := c2.DeleteConf(context.Background(), dr)
	if err != nil {
		log.Printf("could not delete service %s with key %s", dr.ServiceName, dr.HashKey)
	}
	log.Printf("delete hask key code %d, and status %s", rep.Code, rep.Status)
	return err
}

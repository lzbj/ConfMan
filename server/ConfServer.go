package server

import (
	"fmt"
	cm "github.com/lzbj/ConfMan/proto/ConfMan"
	p "github.com/lzbj/ConfMan/server/persist"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type ConfManServer struct {
	server *grpc.Server
}

var persist *p.Persistence

func PersistInit(betype, beaddress string) error {
	ps, err := p.NewPersistence(betype, beaddress)
	if err != nil {
		return err
	}
	persist = ps
	return nil
}
func NewConfManServer(server *grpc.Server) *ConfManServer {

	return &ConfManServer{
		server: server,
		//cPersist: ps,
	}
}

func (s *ConfManServer) Serve(port int) error {
	address := fmt.Sprintf(":%d", port)

	log.Printf("server is running on port %s", address)

	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	s.server = grpc.NewServer()
	cm.RegisterConfManServer(s.server, s)

	err = s.server.Serve(l)

	if err != nil {
		return err
	}

	return nil
}

func (s *ConfManServer) GetConf(c context.Context, rq *cm.GetConfRequest) (*cm.ConfigurationModel, error) {
	fmt.Println(rq.HashKey)
	//combinedKey := rq.ServiceName + "/" + rq.HashKey
	res, err := persist.GetConfig(c, rq.HashKey)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ConfManServer) CreateConf(c context.Context, ur *cm.UpdateConfRequest) (*cm.ConfigurationModel, error) {
	c1 := &cm.ConfigurationModel{}
	c1.ServiceName = ur.ServiceName
	c1.HashKey = ur.HashKey
	c1.HashValue = ur.HashValue
	cm, err := persist.Update(c, c1)
	if err != nil {
		return nil, err
	}
	return cm, nil
}

func (s *ConfManServer) UpdateConf(c context.Context, rq *cm.UpdateConfRequest) (*cm.ConfigurationModel, error) {
	c1 := &cm.ConfigurationModel{}
	c1.ServiceName = rq.ServiceName
	c1.HashKey = rq.HashKey
	c1.HashValue = rq.HashValue
	cm, err := persist.Update(c, c1)
	if err != nil {
		return nil, err
	}
	return cm, nil
}

func (s *ConfManServer) DeleteConf(c context.Context, rq *cm.DeleteConfRequest) (*cm.DeleteConfResponse, error) {
	c1 := &cm.ConfigurationModel{}
	c1.ServiceName = rq.ServiceName
	c1.HashKey = rq.HashKey
	_, err := persist.Delete(c, rq.HashKey)
	resp := &cm.DeleteConfResponse{}
	if err != nil {
		resp.Code = 400
		resp.Status = "delete failed"
		return resp, err
	}
	resp.Code = 200
	resp.Status = "delete succeed"
	return resp, nil
}

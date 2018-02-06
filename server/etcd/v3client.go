package etcd

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"log"
	"time"
)

var (
	dialTimeout    = 2 * time.Second
	requestTimeout = 5 * time.Second
)

func NewClient(cfs string) (kv clientv3.KV, err error) {
	cli, err := clientv3.New(clientv3.Config{DialTimeout: dialTimeout,
		//Endpoints: []string{"127.0.0.1:2379"},
		Endpoints: []string{cfs},
	})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	kv = clientv3.NewKV(cli)
	err = nil
	return
}

func GetSingleValue(ctx context.Context, kv clientv3.KV) {

}

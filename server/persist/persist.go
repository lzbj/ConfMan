package persist

import (
	"context"
	"errors"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/lzbj/ConfMan/proto/ConfMan"
	"github.com/lzbj/ConfMan/server/etcd"
)

type Persistence struct {
	kv clientv3.KV
}

func NewPersistence(betype, address string) (*Persistence, error) {
	switch betype {
	case "etcd":
		kv, err := etcd.NewClient(address)
		if err != nil {
			return nil, errors.New("connect backend fails")
		}
		return &Persistence{kv: kv}, nil
	default:
		return nil, errors.New("Unsupported backend type")
	}
	return nil, nil
}

func (p *Persistence) GetConfig(c context.Context, key string) (*ConfMan.ConfigurationModel, error) {
	fmt.Println("key is :" + key)
	resp, err := p.kv.Get(c, key)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%v", resp)
	fmt.Printf("result: %d ", resp.Count)
	if resp.Count == 0 {
		return nil, errors.New("no value find for the key")
	}
	value := string(resp.Kvs[0].Value)
	cm := &ConfMan.ConfigurationModel{HashKey: key, HashValue: value}
	return cm, nil
}

func (p *Persistence) Update(c context.Context, cf *ConfMan.ConfigurationModel) (*ConfMan.ConfigurationModel, error) {
	fmt.Println("key is :" + cf.HashKey)
	resp, err := p.kv.Put(c, cf.HashKey, cf.HashValue)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%v", resp)
	fmt.Printf("result: %d ", resp.Header.Revision)

	cm := &ConfMan.ConfigurationModel{ServiceName: cf.ServiceName, HashKey: cf.HashKey, HashValue: cf.HashValue}
	return cm, nil
}

func (p *Persistence) Save(c *ConfMan.ConfigurationModel) (*ConfMan.ConfigurationModel, error) {
	return nil, nil
}

func (p *Persistence) Delete(c context.Context, key string) (bool, error) {
	fmt.Println("key is :" + key)
	resp, err := p.kv.Delete(c, key)
	if err != nil {
		return false, err
	}

	fmt.Printf("%v", resp)
	fmt.Printf("result: %d ", resp.Header.Revision)

	return true, nil
}

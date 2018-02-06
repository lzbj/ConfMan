package persist

import (
	"context"
	"errors"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/lzbj/ConfMan/proto/ConfMan"
	"github.com/lzbj/ConfMan/server/etcd"
	"github.com/lzbj/ConfMan/server/model"
	"github.com/lzbj/ConfMan/server/util"
	//"log"
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
	mc := &model.ConfigurationModel{}
	fmt.Println("key is :" + key)
	resp, err := p.kv.Get(c, key)
	if err != nil {
		res, _ := util.Convert(mc)
		return res, err
	}

	fmt.Printf("%v", resp)
	fmt.Printf("result: %d ", resp.Count)
	if resp.Count == 0 {
		res, _ := util.Convert(mc)
		return res, err
	}
	mc.HashValue = string(resp.Kvs[0].Value)
	res, _ := util.Convert(mc)
	return res, nil
}

func (p *Persistence) Update(cf *ConfMan.ConfigurationModel) (*ConfMan.ConfigurationModel, error) {
	return nil, nil
}

func (p *Persistence) Save(c *ConfMan.ConfigurationModel) (*ConfMan.ConfigurationModel, error) {
	return nil, nil
}

func (p *Persistence) Delete(key string) (bool, error) {
	return false, nil
}

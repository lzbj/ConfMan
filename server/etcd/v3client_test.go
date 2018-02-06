package etcd

import (
	"github.com/coreos/etcd/pkg/testutil"
	"github.com/lzbj/ConfMan/server/conf"
	"testing"
)

func TestNewClient(t *testing.T) {
	backend := conf.Cfg.GetString("backend.address")
	_, err := NewClient(backend)
	testutil.AssertNil(t, err)
	// Only apply when etcd is ready
	//testutil.AssertNotNil(t,kv)
}

func TestGetSingleValue(t *testing.T) {

}

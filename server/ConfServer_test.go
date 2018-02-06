package server

import (
	"github.com/coreos/etcd/pkg/testutil"
	"github.com/lzbj/ConfMan/server/conf"
	"testing"
)

func TestNewConfManServer(t *testing.T) {
	cfs := conf.NewConfig()
	bty := cfs.GetString("backend.type")
	dpass := cfs.GetString("database.pass")
	testutil.AssertEqual(t, "etcd", bty, "backend type equals")
	testutil.AssertEqual(t, "password", dpass, "db password equals")
}

package conf

import (
	"github.com/coreos/etcd/pkg/testutil"
	"testing"
)

func TestNewConfig(t *testing.T) {
	cfs := NewConfig()
	bty := cfs.GetString("backend.type")
	dpass := cfs.GetString("database.pass")
	testutil.AssertEqual(t, "etcd", bty, "backend type equals")
	testutil.AssertEqual(t, "password", dpass, "db password equals")
}

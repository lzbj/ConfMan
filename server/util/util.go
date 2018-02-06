package util

import (
	//"github.com/golang/protobuf/ptypes"
	cm "github.com/lzbj/ConfMan/proto/ConfMan"
	"github.com/lzbj/ConfMan/server/model"
)

func Convert(m *model.ConfigurationModel) (*cm.ConfigurationModel, error) {
	pcf := &cm.ConfigurationModel{}
	pcf.HashKey = m.HashKey
	pcf.ServiceName = m.ServiceName
	pcf.HashValue = m.HashValue
	return pcf, nil

}

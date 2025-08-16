package server

import (
	"alat/pkg/core/device"

	"google.golang.org/protobuf/proto"
)

type ServerConfig struct {
	DeviceInfo device.DeviceInfo
}

var (
	config       ServerConfig
	infoResponse []byte
)

func Configure(conf ServerConfig) (err error) {
	inf, err := conf.DeviceInfo.ToPBuf()
	if err != nil {
		return err
	}
	infoResponse, err = proto.Marshal(&inf)
	return err
}

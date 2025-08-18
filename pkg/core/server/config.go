package server

import (
	"alat/pkg/core/device"
	"alat/pkg/core/pbuf"

	"google.golang.org/protobuf/proto"
)

type ServerConfig struct {
	DeviceInfo     device.DeviceInfo
	OnPairRequest  func(*pbuf.PairRequest) int
	OnPairResponse func(*pbuf.PairResponse) int
}

var (
	config       *ServerConfig
	infoResponse []byte
)

func Configure(conf *ServerConfig) (err error) {
	config = conf
	inf, err := conf.DeviceInfo.ToPBuf()
	if err != nil {
		return err
	}
	infoResponse, err = proto.Marshal(&inf)
	return err
}

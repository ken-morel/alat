package server

import (
	"alat/pkg/core/device"
	"alat/pkg/core/pbuf"

	"github.com/wailsapp/wails/v2/pkg/options"
	"google.golang.org/protobuf/proto"
)

type ServerConfig struct {
	DeviceName  string
	DeviceCode  string
	DeviceColor options.RGBA
	DeviceType  device.DeviceType
}

var (
	config       ServerConfig
	infoResponse []byte
)

func Configure(conf ServerConfig) (err error) {
	config = conf
	infoResponse, err = proto.Marshal(&pbuf.DeviceInfo{
		Code: conf.DeviceCode,
		Name: conf.DeviceName,
		Type: pbuf.DeviceInfo_DeviceType(conf.DeviceType),
		Color: &pbuf.DeviceColor{
			R: uint32(conf.DeviceColor.R),
			G: uint32(conf.DeviceColor.G),
			B: uint32(conf.DeviceColor.B),
		},
	})
	return
}

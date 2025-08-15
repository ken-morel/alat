package server

import (
	"alat/pkg/core/device"

	"github.com/wailsapp/wails/v2/pkg/options"
)

type ServerConfig struct {
	DeviceName  string
	DeviceCode  string
	DeviceColor options.RGBA
	DeviceType  device.DeviceType
}

var config ServerConfig

func Configure(conf ServerConfig) {
	config = conf
}

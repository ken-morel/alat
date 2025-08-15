// Package device stores high level device queries
package device

import (
	"alat/pkg/core/address"
	"alat/pkg/core/pbuf"

	"github.com/wailsapp/wails/v2/pkg/options"
)

type DeviceInfo struct {
	Address address.Address
	Name    string
	Color   options.RGBA
	Code    string
	Type    DeviceType
}

type DeviceType int

const (
	DeviceTypeDesktop DeviceType = 0
	DeviceTypeMobile  DeviceType = 1
	DeviceTypeTV      DeviceType = 2
	DeviceTypeWeb     DeviceType = 3
)

func NewDeviceInfo(addr address.Address, info *pbuf.DeviceInfo) DeviceInfo {
	return DeviceInfo{
		Address: addr,
		Name:    info.Name,
		Color: options.RGBA{
			R: uint8(info.Color.R),
			G: uint8(info.Color.R),
			B: uint8(info.Color.B),
		},
		Code: info.Code,
		Type: DeviceType(info.Type.Number()),
	}
}

// Package device stores high level device queries
package device

import (
	"alat/pkg/core/address"
	"alat/pkg/core/pbuf"
	"alat/pkg/core/service"

	"github.com/wailsapp/wails/v2/pkg/options"
)

type DeviceInfo struct {
	Address  address.Address   `yaml:"address"`
	Name     string            `yaml:"name"`
	Color    options.RGBA      `yaml:"color"`
	Code     string            `yaml:"code"`
	Type     DeviceType        `yaml:"type"`
	Services []service.Service `yaml:"services"`
}

func (d *DeviceInfo) ToPBuf() (pbuf.DeviceInfo, error) {
	var services []*pbuf.Service
	for _, srv := range d.Services {
		pb := srv.ToPBuf()
		services = append(services, &pb)
	}
	return pbuf.DeviceInfo{
		Code: d.Code,
		Name: d.Name,
		Type: pbuf.DeviceType(d.Type),
		Color: &pbuf.DeviceColor{
			R: uint32(d.Color.R),
			G: uint32(d.Color.G),
			B: uint32(d.Color.B),
		},
		Services: services,
	}, nil
}

type DeviceType int

const (
	DeviceTypeDesktop DeviceType = 0
	DeviceTypeMobile  DeviceType = 1
	DeviceTypeTV      DeviceType = 2
	DeviceTypeWeb     DeviceType = 3
)

func NewDeviceInfo(addr address.Address, info *pbuf.DeviceInfo) DeviceInfo {
	var services []service.Service
	for _, src := range info.Services {
		services = append(services, service.FromPBuf(src))
	}
	return DeviceInfo{
		Address: addr,
		Name:    info.Name,
		Color: options.RGBA{
			R: uint8(info.Color.R),
			G: uint8(info.Color.R),
			B: uint8(info.Color.B),
		},
		Code:     info.Code,
		Type:     DeviceType(info.Type.Number()),
		Services: services,
	}
}

var ThisDeviceInfo DeviceInfo

func SetThisDeviceInfo(d DeviceInfo) {
	ThisDeviceInfo = d
}

// Package device stores high level device queries
package device

import (
	"alat/pkg/core/address"
	"alat/pkg/core/pbuf"
	"alat/pkg/core/service"

	"github.com/wailsapp/wails/v2/pkg/options"
)

type DeviceInfo struct {
	Address  address.Address   `yaml:"address" json:"address"`
	Name     string            `yaml:"name" json:"name"`
	Color    options.RGBA      `yaml:"color" json:"color"`
	Code     string            `yaml:"code" json:"code"`
	Type     DeviceType        `yaml:"type" json:"type"`
	Services []service.Service `yaml:"services" json:"services"`
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
		Ip:       d.Address.IP.String(),
		Port:     uint32(d.Address.Port),
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
	for _, src := range info.GetServices() {
		services = append(services, service.FromPBuf(src))
	}

	return DeviceInfo{
		Address: addr,
		Name:    info.GetName(),
		Color: options.RGBA{
			R: uint8(info.Color.GetR()),
			G: uint8(info.Color.GetG()),
			B: uint8(info.Color.GetB()),
			A: 255,
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

package device

import (
	"alat/pkg/core/device/color"
	"alat/pkg/core/security"
	"alat/pkg/pbuf"
)

type Details struct {
	Color       color.Color
	Name        string
	Type        DeviceType
	Certificate security.Certificate
}

func (d *Details) ToPBUF() *pbuf.DeviceDetails {
	return &pbuf.DeviceDetails{
		Certificate: d.Certificate[:],
		Name:        d.Name,
		Type:        d.Type.ToPBUF(),
		Color:       d.Color.ToPBUF(),
	}
}

func (d *Details) GetInfo() *Info {
	return &Info{
		ID:    d.Certificate.ID(),
		Name:  d.Name,
		Type:  d.Type,
		Color: d.Color,
	}
}

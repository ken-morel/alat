package device

import "alat/pkg/core/security"

type Details struct {
	Color       Color
	Name        string
	Type        DeviceType
	Certificate security.Certificate
}

func (d *Details) GetInfo() *Info {
	return &Info{
		ID:    d.Certificate.ID(),
		Name:  d.Name,
		Type:  d.Type,
		Color: d.Color,
	}
}

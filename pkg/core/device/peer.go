// Package device: holds peer related things
package device

import (
	"alat/pkg/core/security"
	"alat/pkg/pbuf"
)

type Color struct {
	R uint8
	G uint8
	B uint8
}

func (c Color) ToPBUF() *pbuf.Color {
	return &pbuf.Color{
		R: uint32(c.R),
		G: uint32(c.G),
		B: uint32(c.B),
	}
}

type PairedDevice struct {
	Certificate security.Certificate
	Token       security.PairToken
}

func (d *PairedDevice) ToPBUF() *pbuf.PairedDevice {
	return &pbuf.PairedDevice{
		Certificate: d.Certificate[:],
		Token:       d.Token[:],
	}
}

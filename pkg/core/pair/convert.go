// Package pair: converts between internal and pbuf types
package pair

import (
	"alat/pkg/core/device"
	"alat/pkg/core/security"
	"alat/pkg/pbuf"
)

func PBUFToPairedDevice(pbufDevice *pbuf.PairedDevice) *device.PairedDevice {
	var cert security.Certificate
	copy(cert[:], pbufDevice.Certificate)
	var token security.PairToken
	copy(token[:], pbufDevice.Token)
	return &device.PairedDevice{
		Certificate: cert,
		Token:       token,
	}
}

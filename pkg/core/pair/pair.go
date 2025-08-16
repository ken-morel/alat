// Package pair: hold pair device info structures
package pair

import (
	"alat/pkg/core/device"
	"alat/pkg/core/service"
	"crypto/rand"
)

type Pair struct {
	DeviceInfo       device.DeviceInfo `yaml:"deviceinfo"`
	Token            string            `yaml:"token"`
	OldToken         string            `yaml:"oldtoken"`
	Services         []service.Service `yaml:"services"`
	ExposingServices []service.Service `yaml:"exposingservices"`
}

func GeneratePairToken() string {
	data := make([]byte, 255)
	rand.Read(data)
	return string(data)
}

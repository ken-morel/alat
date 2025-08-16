// Package pair: hold pair device info structures
package pair

import (
	"alat/pkg/core/device"
	"alat/pkg/core/service"

	"github.com/google/uuid"
)

type Pair struct {
	DeviceInfo       device.DeviceInfo `yaml:"deviceinfo"`
	Token            string            `yaml:"token"`
	OldToken         string            `yaml:"oldtoken"`
	Services         []service.Service `yaml:"services"`
	ExposingServices []service.Service `yaml:"exposingservices"`
}

func GeneratePairToken() string {
	token, _ := uuid.NewUUID()
	return token.String()
}

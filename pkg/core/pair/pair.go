// Package pair: hold pair device info structures
package pair

import (
	"alat/pkg/core/device"
	"alat/pkg/core/service"
)

type PairedDevice struct {
	DeviceInfo device.DeviceInfo
	Token      string
	OldToken   string
	Services   []service.Service
}

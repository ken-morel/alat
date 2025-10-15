// Package config, manages an interface to application settings
package config

import (
	"alat/pkg/core/device"
	"alat/pkg/core/device/color"
	"alat/pkg/core/security"
)

type AppConfig struct {
	SetupComplete bool                 `yaml:"setupComplete"         json:"setupComplete"`
	DeviceName    string               `yaml:"deviceName"            json:"deviceName"`
	DeviceColor   color.Color          `yaml:"deviceColor"           json:"deviceColor"`
	Certificate   security.Certificate `yaml:"certificate"           json:"certificate"`
	DeviceType    device.DeviceType    `yaml:"deviceType"            json:"deviceType"`
}

type SysInfoConfig struct {
	Enabled      bool  `yaml:"enabled"      json:"enabled"`
	CacheSeconds uint8 `yaml:"cacheSeconds" json:"cacheSeconds"`
}
type FileSendConfig struct {
	Enabled    bool   `yaml:"enabled"    json:"enabled"`
	MaxSize    uint64 `yaml:"maxSize"    json:"maxSize"`
	SaveFolder string `yaml:"saveFolder" json:"saveFolder"`
}
type ServiceConfig struct {
	SysInfo  SysInfoConfig  `yaml:"sysinfo"  json:"sysinfo"`
	FileSend FileSendConfig `yaml:"filesend" json:"filesend"`
}

// Package config, manages an interface to application settings
package config

import (
	"alat/pkg/core/device"
	"alat/pkg/core/device/color"
	"alat/pkg/core/security"
	"math/rand"
	"os"
	"path"
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

func DefaultAppConfig() AppConfig {
	name, err := os.Hostname()
	if err != nil {
		name = "alat"
	}
	cert, _ := security.GenerateCertificate()
	return AppConfig{
		SetupComplete: false,
		DeviceName:    name,
		DeviceColor:   color.Colors[rand.Int()%len(color.Colors)],
		Certificate:   cert,
		DeviceType:    device.UnspecifiedDevice,
	}
}

func DefaultServiceConfig() ServiceConfig {
	home, err := os.UserHomeDir()
	if err != nil {
		home = "."
	} else {
		dloads := path.Join(home, "Downloads")
		info, err := os.Stat(dloads)
		if err == nil && info.IsDir() {
			home = dloads
		}
	}
	return ServiceConfig{
		SysInfo: SysInfoConfig{
			Enabled:      true,
			CacheSeconds: 10,
		},
		FileSend: FileSendConfig{
			Enabled:    true,
			MaxSize:    0,
			SaveFolder: home,
		},
	}
}

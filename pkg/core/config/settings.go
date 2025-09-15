// Package config, manages an interface to application settings
package config

import (
	"math/rand"
	"os"
	"path"

	"alat/pkg/core/device/color"
	"alat/pkg/core/security"

	"gopkg.in/yaml.v3"
)

type AppSettings struct {
	SetupComplete bool                 `yaml:"setupComplete"         json:"setupComplete"`
	DeviceName    string               `yaml:"deviceName"            json:"deviceName"`
	DeviceColor   color.Color          `yaml:"deviceColor"           json:"deviceColor"`
	Certificate   security.Certificate `yaml:"certificate,omitempty" json:"certificate,omitempty"`
}

type SysInfoSettings struct {
	Enabled      bool  `yaml:"enabled"      json:"enabled"`
	CacheSeconds uint8 `yaml:"cacheSeconds" json:"cacheSeconds"`
}
type FileSendSettings struct {
	Enabled    bool   `yaml:"enabled"    json:"enabled"`
	MaxSize    uint64 `yaml:"maxSize"    json:"maxSize"`
	SaveFolder string `yaml:"saveFolder" json:"saveFolder"`
}
type ServiceSettings struct {
	SysInfo  SysInfoSettings  `yaml:"sysinfo"  json:"sysinfo"`
	FileSend FileSendSettings `yaml:"filesend" json:"filesend"`
}

func DefaultAppSettings() *AppSettings {
	defaultName, _ := os.Hostname()
	cert, _ := security.GenerateCertificate()
	return &AppSettings{
		SetupComplete: false,
		DeviceName:    defaultName,
		DeviceColor:   color.Colors[rand.Int()%len(color.Colors)],
		Certificate:   cert,
	}
}

func LoadAppSettings(filePath string) (*AppSettings, error) {
	defaults := DefaultAppSettings()

	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			if err := SaveAppSettings(defaults, filePath); err != nil {
				return nil, err
			}
			return defaults, nil
		}
		return nil, err
	}

	var settings AppSettings
	if err := yaml.Unmarshal(data, &settings); err != nil {
		return nil, err
	}
	return &settings, nil
}

func SaveAppSettings(settings *AppSettings, filePath string) error {
	data, err := yaml.Marshal(settings)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0o644)
}

func LoadServiceSettings(filePath string) (*ServiceSettings, error) {
	home, _ := os.UserHomeDir()

	defaults := &ServiceSettings{
		SysInfo: SysInfoSettings{
			Enabled:      true,
			CacheSeconds: 10,
		},
		FileSend: FileSendSettings{
			Enabled:    true,
			MaxSize:    0,
			SaveFolder: path.Join(home, "Downloads"),
		},
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			if err := SaveServiceSettings(defaults, filePath); err != nil {
				return nil, err
			}
			return defaults, nil
		}
		return nil, err
	}

	var settings ServiceSettings
	if err := yaml.Unmarshal(data, &settings); err != nil {
		return nil, err
	}
	return &settings, nil
}

func SaveServiceSettings(settings *ServiceSettings, filePath string) error {
	data, err := yaml.Marshal(settings)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0o644)
}

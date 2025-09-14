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
	SetupComplete bool                 `yaml:"setupComplete"`
	DeviceName    string               `yaml:"deviceName"`
	DeviceColor   color.Color          `yaml:"deviceColor"`
	Certificate   security.Certificate `yaml:"certificate,omitempty"`
}

type SysInfoSettings struct {
	Enabled      bool  `yaml:"enabled"`
	CacheSeconds uint8 `yaml:"cacheseconds"`
}
type FileSendSettings struct {
	Enabled    bool   `yaml:"enabled"`
	MaxSize    uint64 `yaml:"maxsize"`
	SaveFolder string `yaml:"savefolder"`
}
type ServiceSettings struct {
	SysInfo  SysInfoSettings `yaml:"sysinfo"`
	FileSend FileSendSettings
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

func LoadAppSettings(configDir string) (*AppSettings, error) {
	p := path.Join(configDir, "settings.yml")

	defaults := DefaultAppSettings()

	data, err := os.ReadFile(p)
	if err != nil {
		if os.IsNotExist(err) {
			if err := SaveAppSettings(defaults, configDir); err != nil {
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

func SaveAppSettings(settings *AppSettings, configDir string) error {
	p := path.Join(configDir, "settings.yml")
	data, err := yaml.Marshal(settings)
	if err != nil {
		return err
	}
	return os.WriteFile(p, data, 0o644)
}

func LoadServiceSettings(configDir string) (*ServiceSettings, error) {
	p := path.Join(configDir, "services.yml")
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

	data, err := os.ReadFile(p)
	if err != nil {
		if os.IsNotExist(err) {
			if err := SaveServiceSettings(defaults, configDir); err != nil {
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

func SaveServiceSettings(settings *ServiceSettings, configDir string) error {
	p := path.Join(configDir, "services.yml")
	data, err := yaml.Marshal(settings)
	if err != nil {
		return err
	}
	return os.WriteFile(p, data, 0o644)
}

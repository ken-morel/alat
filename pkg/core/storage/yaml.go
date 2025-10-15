package storage

import (
	"alat/pkg/core/config"
	"alat/pkg/core/device"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

type YAMLNodeStorage struct {
	path string
}

func (ns *YAMLNodeStorage) GetPairedDevices() ([]device.PairedDevice, error) {
	data, err := os.ReadFile(ns.path)
	if err != nil {
		if os.IsNotExist(err) {
			return []device.PairedDevice{}, nil
		}
		return nil, err
	}

	var devices []device.PairedDevice
	if err := yaml.Unmarshal(data, &devices); err != nil {
		return nil, err
	}

	return devices, nil
}

func (ns *YAMLNodeStorage) AddPairedDevice(newDevice device.PairedDevice) error {
	devices, err := ns.GetPairedDevices()
	if err != nil {
		return err
	}

	for _, d := range devices {
		if d.Certificate == newDevice.Certificate {
			return nil
		}
	}

	devices = append(devices, newDevice)

	data, err := yaml.Marshal(devices)
	if err != nil {
		return err
	}

	return os.WriteFile(ns.path, data, 0o644)
}

func (ns *YAMLNodeStorage) GetAppConfig(defaults config.AppConfig) (*config.AppConfig, error) {
	return LoadAppConfig(path.Join(ns.path, "app.yml"), defaults)
}
func (ns *YAMLNodeStorage) SetAppConfig(conf config.AppConfig) error {
	return SaveAppConfig(conf, path.Join(ns.path, "app.yml"))
}

func (ns *YAMLNodeStorage) GetServiceConfig(defaults config.ServiceConfig) (*config.ServiceConfig, error) {
	return LoadServiceConfig(path.Join(ns.path, "services.yml"), defaults)
}
func (ns *YAMLNodeStorage) SetServiceConfig(conf config.ServiceConfig) error {
	return SaveServiceConfig(conf, path.Join(ns.path, "services.yml"))
}

func CreateYAMLNodeStorage(path string) *YAMLNodeStorage {
	return &YAMLNodeStorage{
		path: path,
	}
}

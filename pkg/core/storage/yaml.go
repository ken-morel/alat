package storage

import (
	"alat/pkg/core/config"
	"alat/pkg/core/device"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

type YAMLNodeStorage struct {
	path                 string
	defaultAppConfig     config.AppConfig
	defaultServiceConfig config.ServiceConfig
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
func (ns *YAMLNodeStorage) AppConfigPath() string {
	return path.Join(ns.path, "app.yml")
}
func (ns *YAMLNodeStorage) ServiceConfigPath() string {
	return path.Join(ns.path, "services.yml")

}

func (ns *YAMLNodeStorage) GetAppConfig() (*config.AppConfig, error) {
	data, err := os.ReadFile(ns.AppConfigPath())
	if err != nil {
		return &ns.defaultAppConfig, nil
	}
	var settings config.AppConfig
	if err := yaml.Unmarshal(data, &settings); err != nil {
		return nil, err
	}
	return &settings, nil
}
func (ns *YAMLNodeStorage) SetAppConfig(conf config.AppConfig) error {
	data, err := yaml.Marshal(conf)
	if err != nil {
		return err
	}
	return os.WriteFile(ns.AppConfigPath(), data, 0o644)

}
func (ns *YAMLNodeStorage) DefaultAppConfig(conf config.AppConfig) {
	ns.defaultAppConfig = conf
}

func (ns *YAMLNodeStorage) GetServiceConfig() (*config.ServiceConfig, error) {
	data, err := os.ReadFile(ns.ServiceConfigPath())
	if err != nil {

		return &ns.defaultServiceConfig, nil
	}

	var settings config.ServiceConfig
	if err := yaml.Unmarshal(data, &settings); err != nil {
		return nil, err
	}
	return &settings, nil

}
func (ns *YAMLNodeStorage) SetServiceConfig(conf config.ServiceConfig) error {
	data, err := yaml.Marshal(conf)
	if err != nil {
		return err
	}
	return os.WriteFile(ns.ServiceConfigPath(), data, 0o644)
}
func (ns *YAMLNodeStorage) DefaultServiceConfig(conf config.ServiceConfig) {
	ns.defaultServiceConfig = conf
}

func CreateYAMLNodeStorage(path string) *YAMLNodeStorage {
	return &YAMLNodeStorage{
		path: path,
	}
}

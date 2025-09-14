package storage

import (
	"alat/pkg/core/device"
	"os"

	"gopkg.in/yaml.v3"
)

type YAMLNodeStorage struct {
	path string
}

func (ns *YAMLNodeStorage) GetPaired() ([]device.PairedDevice, error) {
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

func (ns *YAMLNodeStorage) AddPaired(newDevice device.PairedDevice) error {
	devices, err := ns.GetPaired()
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

func CreateYAMLNodeStorage(path string) *YAMLNodeStorage {
	return &YAMLNodeStorage{
		path: path,
	}
}

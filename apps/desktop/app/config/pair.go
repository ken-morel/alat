package config

import (
	"alat/pkg/core/pair"
	"fmt"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

func GetPairsConfigFile() string {
	return path.Join(AlatConfigDir, "paireddevices.yaml")
}

func GetPairedDevices() ([]pair.Pair, error) {
	var devices []pair.Pair
	path := GetPairsConfigFile()
	file, err := os.Open(path)
	if err != nil {
		return devices, err
	}
	defer file.Close()

	err = yaml.NewDecoder(file).Decode(&devices)
	if err != nil {
		// If the file is empty, Decode returns an error. We can treat this as an empty list.
		return []pair.Pair{}, nil
	}
	return devices, nil
}

func AddPairedDevice(newDevice pair.Pair) error {

devices, err := GetPairedDevices()
	if err != nil {
		return err
	}

	// Check if device is already paired
	for _, d := range devices {
		if d.DeviceInfo.Code == newDevice.DeviceInfo.Code {
			// TODO: Handle updating existing device instead of just returning
			fmt.Printf("Device %s is already paired.\n", newDevice.DeviceInfo.Name)
			return nil
		}
	}


devices = append(devices, newDevice)

	path := GetPairsConfigFile()
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	enc := yaml.NewEncoder(file)
	defer enc.Close()

	return enc.Encode(devices)
}


func InitPair() (err error) {
	path := GetPairsConfigFile()
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()
		
		enc := yaml.NewEncoder(file)
		defer enc.Close()
		return enc.Encode([]pair.Pair{})
	}
	return
}
package config

import (
	"alat/pkg/core/pair"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

func GetPairsConfigFile() string {
	return path.Join(AlatConfigDir, "paireddevices.yaml")
}

func GetPairedDevices() ([]pair.Pair, error) {
	var devices []pair.Pair
	var err error
	path := GetPairsConfigFile()
	file, err := os.Open(path)
	if err != nil {
		return devices, err
	}
	err = yaml.NewDecoder(file).Decode(devices)
	return devices, err
}

func InitPair() (err error) {
	path := GetPairsConfigFile()
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		enc := yaml.NewEncoder(file)
		defer enc.Close()
		enc.Encode([]pair.Pair{})
	}
	return
}

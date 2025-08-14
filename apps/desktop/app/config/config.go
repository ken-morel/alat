// Package config manages app settings and preferences
package config

import (
	"alat/pkg/core"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/wailsapp/wails/v2/pkg/options"
	"gopkg.in/yaml.v3"
)

type Config struct {
	DeviceName  string
	DeviceColor options.RGBA
	DeviceCode  string
}

var (
	config        Config
	AlatConfigDir string
)
var Ready bool = false

func GetConfigDir() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatalln("Could not get app config dir")
		os.Exit(1)
	}
	return path.Join(configDir, core.AppID)
}

func GetMainConfigFile() string {
	return path.Join(AlatConfigDir, "config.yaml")
}

func LoadConfig() error {
	filePath := GetMainConfigFile()
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	err = yaml.NewDecoder(file).Decode(&config)
	if err != nil {
		return err
	} else {
		Ready = true
		return nil
	}
}

func Init() error {
	if err := LoadConfig(); err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

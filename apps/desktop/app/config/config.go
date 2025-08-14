// Package config manages app settings and preferences
package config

import (
	"alat/pkg/core"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/google/uuid"
	"github.com/wailsapp/wails/v2/pkg/options"
	"gopkg.in/yaml.v3"
)

type Config struct {
	DeviceName  string       `yaml:"deviceName"`
	DeviceColor options.RGBA `yaml:"deviceColor"`
	DeviceCode  string       `yaml:"deviceCode"`
	Language    string       `yaml:"language"`
	AutoStart   bool         `yaml:"autoStart"`
	Theme       string       `yaml:"theme"`
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
	defer file.Close()
	err = yaml.NewDecoder(file).Decode(&config)
	if err != nil {
		return err
	}

	Ready = true
	return nil
}

func SaveConfig(cfg *Config) error {
	filePath := GetMainConfigFile()
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	defer encoder.Close()

	return encoder.Encode(cfg)
}

func GetConfig() Config {
	return config
}

func GenerateDeviceCode() string {
	return uuid.NewString()
}

func Init() error {
	AlatConfigDir = GetConfigDir()
	if err := os.MkdirAll(AlatConfigDir, 0750); err != nil {
		return fmt.Errorf("could not create config directory: %w", err)
	}

	if err := LoadConfig(); err != nil {
		hostname, _ := os.Hostname()
		config = Config{
			DeviceName:  hostname,
			DeviceColor: options.RGBA{R: 0, G: 0, B: 0, A: 255},
			DeviceCode:  GenerateDeviceCode(),
			Language:    "en-cm",
			AutoStart:   false,
			Theme:       "light",
		}
		if err := SaveConfig(&config); err != nil {
			return fmt.Errorf("could not save initial config: %w", err)
		}
	}
	return nil
}

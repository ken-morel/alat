package main

import (
	"os"
	"path"

	"alat/pkg/core"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	Autostart bool `yaml:"autostart"`
}

var appConfig = AppConfig{
	Autostart: true,
}

func (a *AppConfig) Load() error {
	f, err := os.Open(path.Join(getConfigDir(), "config.yml"))
	if err != nil {
		return err
	}
	defer f.Close()

	return yaml.NewDecoder(f).Decode(a)
}

func (a *AppConfig) Save() error {
	data, err := yaml.Marshal(a)
	if err != nil {
		return err
	}

	return os.WriteFile(path.Join(getConfigDir(), "config.yml"), data, 0o644)
}

func getConfigDir() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		configDir = path.Join(os.TempDir(), core.DesktopAppID)
	} else {
		configDir = path.Join(configDir, core.DesktopAppID)
	}
	_ = os.MkdirAll(configDir, 0o755)
	return configDir
}

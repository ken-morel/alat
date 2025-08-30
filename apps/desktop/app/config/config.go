// Package config: Manages app configuration
package config

import (
	"alat/pkg/core"
	"os"
	"path"
)

var configDir string

func Init() error {
	dir, err := os.UserConfigDir()
	if err != nil {
		dir = path.Join(os.TempDir(), core.AppID)
	} else {
		dir = path.Join(dir, core.AppID)
	}

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	configDir = dir
	return nil
}

func GetConfigDir() string {
	return configDir
}

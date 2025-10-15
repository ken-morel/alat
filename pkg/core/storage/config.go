package storage

import (
	"alat/pkg/core/config"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadAppConfig(filePath string, defaults *config.AppConfig) (*config.AppConfig, error) {

	data, err := os.ReadFile(filePath)
	if err != nil {
		if err := SaveAppConfig(defaults, filePath); err != nil {
			return nil, err
		}
		return defaults, nil
	}

	var settings config.AppConfig
	if err := yaml.Unmarshal(data, &settings); err != nil {
		return nil, err
	}
	return &settings, nil
}

func SaveAppConfig(settings *config.AppConfig, filePath string) error {
	data, err := yaml.Marshal(settings)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0o644)
}

func LoadServiceSettings(filePath string, defaults *config.ServiceConfig) (*config.ServiceConfig, error) {

	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			if err := SaveServiceSettings(defaults, filePath); err != nil {
				return nil, err
			}
			return defaults, nil
		}
		return nil, err
	}

	var settings config.ServiceConfig
	if err := yaml.Unmarshal(data, &settings); err != nil {
		return nil, err
	}
	return &settings, nil
}

func SaveServiceSettings(settings *config.ServiceConfig, filePath string) error {
	data, err := yaml.Marshal(settings)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0o644)
}

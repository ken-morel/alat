package main

import (
	"alat/pkg/core/config"
	"alat/pkg/core/device"
	"alat/pkg/core/node"
	"alat/pkg/core/storage"
)

func createNode() (*node.Node, error) {
	configDir := getConfigDir()
	appConfig := config.DefaultAppConfig()
	appConfig.DeviceType = device.DesktopDevice
	serviceConfig := config.DefaultServiceConfig()

	store := storage.CreateYAMLNodeStorage(configDir, appConfig, serviceConfig)
	return node.CreateNode(store)
}

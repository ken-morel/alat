package main

import (
	"log"

	"alat/pkg/core/config"
	"alat/pkg/core/device"
	"alat/pkg/core/node"
	"alat/pkg/core/storage"
)

func createNode() *node.Node {
	configDir := getConfigDir()
	appConfig := config.DefaultAppConfig()
	appConfig.DeviceType = device.DesktopDevice
	serviceConfig := config.DefaultServiceConfig()

	store := storage.CreateYAMLNodeStorage(configDir, appConfig, serviceConfig)
	n, err := node.CreateNode(store)
	if err != nil {
		log.Fatal("Error creating node", err)
	}
	return n
}

package main

import (
	"alat/apps/desktop/app"
	"embed"
)

import (
	"log"
	"os"
	"path"

	"alat/pkg/core"
	"alat/pkg/core/config"
	"alat/pkg/core/device"
	"alat/pkg/core/node"
	"alat/pkg/core/storage"
)

//go:embed all:frontend/build
var assets embed.FS

func main() {

	configDir, err := os.UserConfigDir()
	if err != nil {
		configDir = path.Join(os.TempDir(), core.DesktopAppID)
	} else {
		configDir = path.Join(configDir, core.DesktopAppID)
	}
	_ = os.MkdirAll(configDir, 0755)

	appConfig := config.DefaultAppConfig()
	appConfig.DeviceType = device.DesktopDevice
	serviceConfig := config.DefaultServiceConfig()

	store := storage.CreateYAMLNodeStorage(configDir, appConfig, serviceConfig)
	n, err := node.CreateNode(store)
	if err != nil {
		log.Fatal("Error creating node", err)
	}
	err = app.NewApp(assets, n).Run()
	if err != nil {
		println("Error:", err.Error())
	}
	err = app.NewApp(assets, n).Run()
	if err != nil {
		println("Error:", err.Error())
	}
}

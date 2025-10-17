package main

import (
	"embed"
	"log"
	"os"
	"path"

	"alat/apps/desktop/app"
	"alat/pkg/core"
	"alat/pkg/core/config"
	"alat/pkg/core/device"
	"alat/pkg/core/node"
	"alat/pkg/core/storage"

	"github.com/getlantern/systray"
)

//go:embed all:frontend/build
var assets embed.FS

//go:embed logos/tray.png
var iconPngData []byte

func createNode() *node.Node {
	configDir, err := os.UserConfigDir()
	if err != nil {
		configDir = path.Join(os.TempDir(), core.DesktopAppID)
	} else {
		configDir = path.Join(configDir, core.DesktopAppID)
	}
	_ = os.MkdirAll(configDir, 0o755)

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

func main() {
	n := createNode()

	a := app.NewApp(assets, n)

	systray.Register(func() {
		if len(iconPngData) > 0 {
			systray.SetIcon(iconPngData)
		}
		systray.SetTitle("Alat")
		systray.SetTooltip("Alat desktop application")

		mShow := systray.AddMenuItem("Show", "Show the app")
		mHide := systray.AddMenuItem("Hide", "Hide the app")
		mQuit := systray.AddMenuItem("Quit", "Close and stop alat")

		go func() {
			for {
				select {
				case <-mShow.ClickedCh:
					a.Show()
				case <-mHide.ClickedCh:
					a.Hide()
				case <-mQuit.ClickedCh:
					systray.Quit()
				}
			}
		}()
	}, func() {
		a.Quit()
	})
	a.Run()
}

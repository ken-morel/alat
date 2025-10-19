package main

import (
	"embed"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"time"

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

//go:embed logos/tray-x.png
var iconXPngData []byte

//go:embed logos/tray.ico
var iconIcoData []byte

//go:embed logos/tray-x.ico
var iconXIcoData []byte

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
	conf, _ := n.GetAppConfig()
	if conf != nil && conf.SetupComplete {
		n.Start()
	}

	a := app.NewApp(assets, n)

	systray.Register(func() {
		systray.SetTitle("Alat")
		systray.SetTooltip("Alat desktop application")

		mShow := systray.AddMenuItem("Show", "Show the app")
		mHide := systray.AddMenuItem("Hide", "Hide the app")
		systray.AddSeparator()
		mNode := systray.AddMenuItem("Node", "")
		mNodeStatus := mNode.AddSubMenuItem("Node stalked", "Node status")
		mNodeStatus.Disable()
		mNodeStop := mNode.AddSubMenuItem("Stop", "Stop the node")
		mNodeStart := mNode.AddSubMenuItem("Start", "Start the node")
		systray.AddSeparator()

		go func() {
			lastRunning := false
			lastPort := 0
			firstTime := true
			for {
				status := n.GetStatus()
				running := status.ServerRunning && status.DiscoveryRunning && status.WorkerRunning
				port := status.Port
				if lastPort != port || lastRunning != running || firstTime {
					firstTime = false
					if running {
						mNodeStatus.SetTitle(fmt.Sprintf("Node running at port %d", port))
						if runtime.GOOS == "windows" {
							systray.SetIcon(iconIcoData)
						} else {
							systray.SetIcon(iconPngData)
						}
					} else {
						mNodeStatus.SetTitle("Node stalked")
						if runtime.GOOS == "windows" {
							systray.SetIcon(iconXIcoData)
						} else {
							systray.SetIcon(iconXPngData)
						}
					}
				}
				time.Sleep(time.Second)
			}
		}()
		systray.AddSeparator()
		mSendFiles := systray.AddMenuItem("Send files", "Select files and a device to send them")
		systray.AddSeparator()
		mQuit := systray.AddMenuItem("Quit", "Close and stop alat")

		go func() {
			for {
				select {
				case <-mShow.ClickedCh:
					a.Show()
				case <-mHide.ClickedCh:
					a.Hide()
				case <-mNodeStop.ClickedCh:
					n.Stop()
				case <-mNodeStart.ClickedCh:
					n.Start()
				case <-mSendFiles.ClickedCh:
					a.OpenSendFilesPage()
				case <-mQuit.ClickedCh:
					systray.Quit()
				}
			}
		}()
	}, func() {
		n.Stop()
		a.Quit()
	})
	a.Run()
}

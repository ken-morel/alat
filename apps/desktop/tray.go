package main

import (
	"fmt"
	"runtime"
	"time"

	"alat/apps/desktop/app"
	"alat/pkg/core/node"

	"github.com/getlantern/systray"
)

func showTray(n *node.Node, a *app.App) {
	systray.Register(func() {
		systray.SetTitle("Alat")
		systray.SetTooltip("Alat desktop application")

		mShow := systray.AddMenuItem("Show", "Show the app")
		mHide := systray.AddMenuItem("Hide", "Hide the app")
		systray.AddSeparator()
		mAutoStart := systray.AddMenuItemCheckbox("Autostart", "Launch alat on startup", appConfig.Autostart)
		systray.AddSeparator()
		mNode := systray.AddMenuItem("Node", "")
		mNodeStatus := mNode.AddSubMenuItem("Node stalked", "Node status")
		mNodeStatus.Disable()
		mNodeStop := mNode.AddSubMenuItem("Stop", "Stop the node")
		mNodeStart := mNode.AddSubMenuItem("Start", "Start the node")

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
				case <-mAutoStart.ClickedCh:
					appConfig.Autostart = !mAutoStart.Checked()
					if mAutoStart.Checked() {
						mAutoStart.Uncheck()
					} else {
						mAutoStart.Check()
					}
					appConfig.Save()
				}
			}
		}()
	}, func() {
		n.Stop()
		a.Quit()
	})
}

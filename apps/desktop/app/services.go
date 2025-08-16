package app

import (
	"alat/apps/desktop/app/config"
	"alat/pkg/core/pair"
	"alat/pkg/core/service/rcfile"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (app *App) GetAndSendFiles(device pair.Pair) {
	name, err := runtime.OpenFileDialog(app.ctx, runtime.OpenDialogOptions{
		Title:           "Choose a file to send to " + device.DeviceInfo.Name,
		ShowHiddenFiles: true,
		ResolvesAliases: true,
	})
	if err != nil {
		return
	}
	go app.doSendFile(name, device)
}

func (app *App) doSendFile(path string, device pair.Pair) {
	channel := make(chan rcfile.SendFileStatus)
	fmt.Println("Sendinf file...")
	go rcfile.SendFile(channel, path, device.DeviceInfo.Address, device.Token)

	for msg := range channel {
		fmt.Println("Message ", msg)
		if msg.Error != nil {
			fmt.Println("Error sending file", msg.Error)
			go runtime.MessageDialog(app.ctx, runtime.MessageDialogOptions{
				Type:    runtime.ErrorDialog,
				Message: "Error sending file: " + msg.Error.Error(),
				Title:   "Error",
			})
		} else if msg.Received {
			go runtime.MessageDialog(app.ctx, runtime.MessageDialogOptions{
				Type:    runtime.InfoDialog,
				Message: "Sent file succesfully ",
				Title:   "Success",
			})
		}

	}

	fmt.Println("Done sending file")
}

func (app *App) SetupServices() {
	conf := config.GetConfig().Services
	rcfile.Init(conf.RCFile)
}

package app

import (
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

func (app *App) onFileReceive(channel chan *rcfile.ReceiveFileStatus, name string, size int) {
	if size == 0 {
		return
	}
	percent := float32(0.0)
	for status := range channel {
		if status.Percent-percent > float32(1000000.0)/float32(size) {
			percent = status.Percent
			runtime.WindowExecJS(app.ctx, fmt.Sprintf("window.downloadCallback('%s', %f);", name, percent))
		}
	}
	runtime.WindowExecJS(app.ctx, fmt.Sprintf("window.downloadCallback('%s', 100)", name))
}

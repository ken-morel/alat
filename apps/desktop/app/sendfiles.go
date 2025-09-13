package app

import (
	"fmt"
	"os"

	"alat/pkg/core/connected"
	"alat/pkg/core/service/filesend"

	rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

type SendFile struct {
	Path string
	Size uint32
}

func (app *App) AskFilesSend(to string) (sendFiles []SendFile, err error) {
	names, err := rt.OpenMultipleFilesDialog(app.ctx, rt.OpenDialogOptions{
		Title:           "Select files to send to " + to,
		ShowHiddenFiles: true,
		ResolvesAliases: true,
	})
	if err != nil {
		return nil, err
	}
	for _, name := range names {
		stat, err := os.Stat(name)
		if err != nil {
			continue
		}
		if stat.IsDir() {
			continue
		}
		sendFiles = append(sendFiles, SendFile{
			Path: name,
			Size: uint32(stat.Size()),
		})
	}
	return sendFiles, err
}

func (app *App) ServiceStartSendFilesToDevice(peer connected.Connected, files []string) error {
	for _, file := range files {
		fmt.Println("File: ", file)
		go func() {
			err := app.serviceRegistery.FileSend.SendFile(app.ctx, &peer, app.nodeDetails, file)
			if err != nil {
				rt.MessageDialog(app.ctx, rt.MessageDialogOptions{
					Type:    rt.ErrorDialog,
					Title:   "File send",
					Message: "Error sending file: " + err.Error(),
				})
			} else {
				rt.MessageDialog(app.ctx, rt.MessageDialogOptions{
					Type:    rt.InfoDialog,
					Title:   "File send",
					Message: "No error ",
				})
			}
		}()
	}
	return nil
}

func (app *App) ServiceGetFileSendStatus() filesend.FileTransfersStatus {
	return app.serviceRegistery.FileSend.GetStatus()
}

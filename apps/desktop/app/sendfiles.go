package app

import (
	"alat/pkg/core/connected"
	"alat/pkg/core/service/filesend"
	"os"

	rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

type SendFile struct {
	Path string
	Size uint32
}

func (app *App) AskFilesSend() (sendFiles []SendFile, err error) {
	names, err := rt.OpenMultipleFilesDialog(app.ctx, rt.OpenDialogOptions{
		Title:           "Select files to send",
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
	var err error
	for sendError := range app.node.Services.FileSend.SendFiles(&peer, files) {
		if sendError != nil {
			err = sendError
		}
	}
	return err
}

func (app *App) ServiceGetFileSendStatus() filesend.FileTransfersStatus {
	return *app.node.Services.FileSend.GetStatus()
}

package app

import (
	"os"

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

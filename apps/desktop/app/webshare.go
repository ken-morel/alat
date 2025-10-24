package app

import "alat/pkg/core/service/webshare"

func (app *App) WebShareStart() (int, error) {
	return app.node.Services.WebShare.Start()
}

func (app *App) WebShareStop() error {
	return app.node.Services.WebShare.Stop()
}

func (app *App) WebShareGetStatus() *webshare.Status {
	return app.node.Services.WebShare.GetStatus()
}

func (app *App) WebShareAddSharedFiles(paths []string) error {
	return app.node.Services.WebShare.AddSharedFiles(paths)
}

func (app *App) WebShareRemoveSharedFile(uuid string) {
	app.node.Services.WebShare.RemoveSharedFilesByUUIDS([]string{uuid})
}

func (app *App) WebShareClearSharedFiles() {
	app.node.Services.WebShare.ClearSharedFiles()
}

func (app *App) WebShareSetPasscode(passcode string) {
	app.node.Services.WebShare.SetPasscode(passcode)
}

func (app *App) WebShareGetPasscode() string {
	return app.node.Services.WebShare.GetPasscode()
}

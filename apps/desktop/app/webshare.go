package app

import "alat/pkg/core/service/webshare"

func (app *App) WebShareStart() (int, error) {
	return app.node.GetServices().WebShare.Start()
}

func (app *App) WebShareStop() error {
	return app.node.GetServices().WebShare.Stop()
}

func (app *App) WebShareGetStatus() *webshare.Status {
	return app.node.GetServices().WebShare.GetStatus()
}

func (app *App) WebShareAddSharedFiles(paths []string) error {
	return app.node.GetServices().WebShare.AddSharedFiles(paths)
}

func (app *App) WebShareRemoveSharedFile(uuid string) {
	app.node.GetServices().WebShare.RemoveSharedFile(uuid)
}

func (app *App) WebShareClearSharedFiles() {
	app.node.GetServices().WebShare.ClearSharedFiles()
}

func (app *App) WebShareSetPasscode(passcode string) {
	app.node.GetServices().WebShare.SetPasscode(passcode)
}

func (app *App) WebShareGetPasscode() string {
	return app.node.GetServices().WebShare.GetPasscode()
}
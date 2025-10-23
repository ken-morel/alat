package app

func (app *App) IsSearchingDevices() bool {
	status := app.node.GetStatus()
	return status.WorkerRunning
}

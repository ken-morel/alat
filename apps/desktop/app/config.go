package app

func (app *App) ConfigReady() bool {
	return app.settings.SetupComplete
}

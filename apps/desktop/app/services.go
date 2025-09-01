package app

import (
	"alat/pkg/core/service"
)

func (app *App) initServices() error {
	app.serviceRegistery = service.NewRegistery()
	return nil
}

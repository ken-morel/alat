// Package app contains the app
package app

import (
	"context"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

// App struct
type App struct {
	ctx    context.Context
	assets embed.FS
}

// NewApp creates a new App application struct
func NewApp(assets embed.FS) *App {
	return &App{assets: assets}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (app *App) startup(ctx context.Context) {
	app.ctx = ctx
}

func (app *App) shutdown(ctx context.Context) {
}

func (app *App) onSecondInstanceLaunch(secondInstanceData options.SecondInstanceData) {
}

func (app *App) Run() error {
	return wails.Run(&options.App{
		Title:  "desktop",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: app.assets,
		},
		BackgroundColour: &options.RGBA{R: 230, G: 240, B: 240, A: 1},
		OnStartup:        app.startup,
		OnShutdown:       app.shutdown,
		SingleInstanceLock: &options.SingleInstanceLock{
			UniqueId:               "cm.rbs.engon.alat",
			OnSecondInstanceLaunch: app.onSecondInstanceLaunch,
		},

		Bind: []any{
			app,
		},
	})
}

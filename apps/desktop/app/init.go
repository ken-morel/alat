package app

import (
	"context"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

func NewApp(fs embed.FS) *App {
	return &App{assets: fs}
}

func (app *App) startup(ctx context.Context) {
	app.ctx = ctx
	if app.initConfig() {
		app.initNode()
	}
}

func (app *App) Run() error {
	return wails.Run(&options.App{
		Title:  "Alat desktop",
		Width:  800,
		Height: 600,
		AssetServer: &assetserver.Options{
			Assets: app.assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 200},
		OnStartup:        app.startup,
		OnShutdown:       app.shutdown,
		Bind: []any{
			app,
		},
		MinWidth:      800,
		MinHeight:     600,
		OnBeforeClose: app.beforeClose,
		SingleInstanceLock: &options.SingleInstanceLock{
			UniqueId:               "cm.engon.alat",
			OnSecondInstanceLaunch: app.onSecondInstance,
		},
		Windows: &windows.Options{
			WindowIsTranslucent: true,
			Theme:               windows.Dark,
			BackdropType:        windows.Acrylic,
			WindowClassName:     "cm.engon.alat",
		},
		Linux: &linux.Options{
			WindowIsTranslucent: true,
			ProgramName:         "cm.engon.alat",
		},
	})
}
func (app *App) onSecondInstance(options.SecondInstanceData) {}

func (app *App) beforeClose(context.Context) bool {
	return false
}
func (app *App) shutdown(context.Context) {}

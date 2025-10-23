package app

import (
	"alat/pkg/core"
	"context"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

func (app *App) startup(ctx context.Context) {
	app.ctx = ctx
	app.started = true
	rt.WindowSetDarkTheme(ctx)
}

func (app *App) Run() error {
	conf, _ := app.node.GetAppConfig()
	setupComplete := conf != nil && conf.SetupComplete
	return wails.Run(&options.App{
		Title:  "Alat",
		Width:  800,
		Height: 1000,
		AssetServer: &assetserver.Options{
			Assets: app.assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 100},
		OnStartup:        app.startup,
		OnShutdown:       app.shutdown,
		StartHidden:      setupComplete,
		WindowStartState: options.Maximised,
		Bind: []any{
			app,
		},
		MinWidth:      800,
		MinHeight:     600,
		OnBeforeClose: app.beforeClose,
		SingleInstanceLock: &options.SingleInstanceLock{
			UniqueId:               core.DesktopAppID,
			OnSecondInstanceLaunch: app.onSecondInstance,
		},
		Windows: &windows.Options{
			WindowIsTranslucent:  true,
			WebviewIsTransparent: true,
			BackdropType:         windows.Acrylic,
			WindowClassName:      core.DesktopAppID,
			Theme:                windows.Dark,
		},
		Linux: &linux.Options{
			WindowIsTranslucent: true,
			ProgramName:         core.DesktopAppID,
		},
	})
}

func (app *App) onSecondInstance(options.SecondInstanceData) {
	rt.WindowUnminimise(app.ctx)
	rt.Show(app.ctx)
}

func (app *App) beforeClose(context.Context) bool {
	app.Hide()
	return true
}
func (app *App) shutdown(context.Context) {}

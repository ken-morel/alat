package app

import (
	"alat/pkg/core"
	"alat/pkg/core/node"
	"context"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

func NewApp(fs embed.FS, n *node.Node) *App {
	return &App{assets: fs, node: n}
}

func (app *App) startup(ctx context.Context) {
	app.ctx = ctx
	rt.WindowSetDarkTheme(ctx)

}

func (app *App) Run() error {
	return wails.Run(&options.App{
		Title:  "Alat",
		Width:  800,
		Height: 600,
		AssetServer: &assetserver.Options{
			Assets: app.assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 100},
		OnStartup:        app.startup,
		OnShutdown:       app.shutdown,
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
		},
		Linux: &linux.Options{
			WindowIsTranslucent: true,
			ProgramName:         core.DesktopAppID,
		},
	})
}
func (app *App) onSecondInstance(options.SecondInstanceData) {}

func (app *App) beforeClose(context.Context) bool {
	return false
}
func (app *App) shutdown(context.Context) {}

// Package app contains the app
package app

import (
	"alat/apps/desktop/app/config"
	"alat/pkg/core"
	"alat/pkg/core/server"
	"context"
	"embed"
	"fmt"

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

func (app *App) startup(ctx context.Context) {
	app.ctx = ctx
	err := config.Init()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if config.Ready {
			config.SetupServer()
			server.Start()
		}
	}
}

func (app *App) shutdown(ctx context.Context) {
}

func (app *App) onSecondInstanceLaunch(secondInstanceData options.SecondInstanceData) {
}

func (app *App) IsSetupComplete() bool {
	return config.Ready
}

func (app *App) GetConfig() config.Config {
	return config.GetConfig()
}

func (app *App) SaveConfig(cfg config.Config) error {
	err := config.SaveConfig(&cfg)
	if err == nil {
		config.SetupServer()
		if !server.Running {
			server.Start()
		}
	}
	return nil
}

// GenerateDeviceCode generates a new unique device code.
func (app *App) GenerateDeviceCode() string {
	return config.GenerateDeviceCode()
}

func (app *App) Run() error {
	return wails.Run(&options.App{
		Title:  "desktop",
		Width:  800,
		Height: 600,
		AssetServer: &assetserver.Options{
			Assets: app.assets,
		},
		BackgroundColour: &options.RGBA{R: 4, G: 2, B: 4, A: 255},
		//&options.RGBA{R: 17, G: 23, B: 17, A: 255},
		OnStartup:  app.startup,
		OnShutdown: app.shutdown,
		SingleInstanceLock: &options.SingleInstanceLock{
			UniqueId:               core.AppID,
			OnSecondInstanceLaunch: app.onSecondInstanceLaunch,
		},

		Bind: []any{
			app,
		},
	})
}

// Package app contains the app
package app

import (
	"alat/apps/desktop/app/config"
	"alat/pkg/core"
	"alat/pkg/core/client"
	"alat/pkg/core/device"
	"alat/pkg/core/pair"
	"alat/pkg/core/server"
	"alat/pkg/core/service"
	"context"
	"embed"
	"fmt"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx          context.Context
	assets       embed.FS
	pandingPairs []PendingPair
}

type PendingPair struct {
	Device   device.DeviceInfo
	Token    string
	Services []service.Service
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
			config.SetupDevice()
			app.setupServer()
			server.Start()
		}
	}
}

func (app *App) setupServer() {
	fmt.Println("Seting up server with info", device.ThisDeviceInfo)
	server.Configure(&server.ServerConfig{
		DeviceInfo:     device.ThisDeviceInfo,
		OnPairRequest:  app.HandlePairRequest,
		OnPairResponse: app.HandlePairResponse,
	})
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

func (app *App) AskDirectory() (path string, err error) {
	path, err = runtime.OpenDirectoryDialog(app.ctx, runtime.OpenDialogOptions{
		Title: "Choose directory",
	})
	return
}

func (app *App) GetServices() []service.Service {
	return config.GetServices()
}

func (app *App) SaveConfig(cfg config.Config) error {
	fmt.Println("Saving config", cfg)
	err := config.SaveConfig(&cfg)
	if err == nil {
		app.setupServer()
		if !server.Running {
			server.Start()
		}
	} else {
		fmt.Println("Errror saving config", err)
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

func (app *App) RequestPair(deviceInfo device.DeviceInfo, services []service.Service) error {
	token := pair.GeneratePairToken()
	res, err := client.SendPairRequest(
		deviceInfo.Address,
		token,
		services,
	)
	if err != nil {
		return err
	}

	if res {
		app.pandingPairs = append(app.pandingPairs, PendingPair{
			Device:   deviceInfo,
			Token:    token,
			Services: services,
		})
		return fmt.Errorf("pairing request pending by device")
	} else {
		return fmt.Errorf("pairing request rejected by device")
	}
}

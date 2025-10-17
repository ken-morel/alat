// Package app stores the app initialization and bindings.
package app

import (
	"context"
	"embed"
	"os"

	"alat/pkg/core/device/color"
	"alat/pkg/core/node"

	"github.com/emersion/go-autostart"
	rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx       context.Context
	assets    embed.FS
	node      *node.Node
	started   bool
	autostart *autostart.App
}

func NewApp(fs embed.FS, n *node.Node) *App {
	execPath, _ := os.Executable()

	return &App{assets: fs, node: n, started: false, autostart: &autostart.App{
		Name:        "alat",
		DisplayName: "Alat",
		Exec:        []string{execPath},
	}}
}

func (app *App) IsAutostartEnabled() bool {
	return app.autostart.IsEnabled()
}

func (app *App) EnableAutostart() {
	_ = app.autostart.Enable()
}

func (app *App) DisableAutostart() {
	_ = app.autostart.Disable()
}

func (app *App) GetAlatColors() []color.Color {
	return color.Colors
}

func (app *App) Show() {
	if app.started {
		rt.Show(app.ctx)
	}
}

func (app *App) Hide() {
	if app.started {
		rt.Hide(app.ctx)
	}
}

func (app *App) Quit() {
	if app.started {
		rt.Quit(app.ctx)
	}
}

func (app *App) OpenSendFilesPage() {
	rt.EventsEmit(app.ctx, "send-files")
}

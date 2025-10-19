// Package app stores the app initialization and bindings.
package app

import (
	"context"
	"embed"

	"alat/pkg/core/device/color"
	"alat/pkg/core/node"

	rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx     context.Context
	assets  embed.FS
	node    *node.Node
	started bool
}

func NewApp(fs embed.FS, n *node.Node) *App {
	return &App{assets: fs, node: n, started: false}
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

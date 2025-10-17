// Package app stores the app initialization and bindings.
package app

import (
	"context"
	"embed"

	"alat/pkg/core/device/color"
	"alat/pkg/core/node"
)

type App struct {
	ctx    context.Context
	assets embed.FS
	node   *node.Node
}

func (app *App) GetAlatColors() []color.Color {
	return color.Colors
}

// Package app stores the app initialization and bindings.
package app

import (
	"context"
	"embed"

	core_config "alat/pkg/core/config"
	"alat/pkg/core/device"
	"alat/pkg/core/device/color"
	"alat/pkg/core/node"
	"alat/pkg/core/service"
	"alat/pkg/core/storage"
)

// The app sourve code

type App struct {
	ctx              context.Context
	assets           embed.FS
	settings         *core_config.AppSettings
	serviceSettings  *core_config.ServiceSettings
	nodeStore        storage.NodeStorage
	node             *node.Node
	serviceRegistery *service.Registry
	nodeDetails      *device.Details
}

func (app *App) GetAlatColors() []color.Color {
	return color.Colors
}

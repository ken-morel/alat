// Package app stores the app initialization and bindings.
package app

import (
	"alat/apps/desktop/app/config"
	"alat/pkg/core/device"
	"alat/pkg/core/device/color"
	"alat/pkg/core/node"
	"alat/pkg/core/service"
	"alat/pkg/core/storage"
	"context"
	"embed"
)

type App struct {
	ctx              context.Context
	assets           embed.FS
	settings         *config.AppSettings
	serviceSettings  *config.ServiceSettings
	nodeStore        storage.NodeStorage
	node             *node.Node
	serviceRegistery *service.Registry
	nodeDetails      *device.Details
}

func (app *App) GetAlatColors() []color.Color {
	return color.Colors
}

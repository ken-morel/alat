package app

import "alat/pkg/core/connected"

func (app *App) GetConnectedDevices() []connected.Connected {
	return app.node.Connected.GetConnectedDevices()
}

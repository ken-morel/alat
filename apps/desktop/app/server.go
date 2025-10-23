package app

import "alat/pkg/core/node"

func (app *App) GetNodeStatus() *node.Status {
	if app.node == nil {
		return nil
	} else {
		return app.node.GetStatus()
	}
}

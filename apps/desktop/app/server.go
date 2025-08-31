package app

import "alat/pkg/core/node"

type ServerStatus struct {
	MDNSRunning    bool
	ServiceRunning bool
}

func (app *App) GetNodeStatus() *node.Status {
	if app.node == nil {
		return nil
	} else {
		return app.node.GetStatus()
	}
}

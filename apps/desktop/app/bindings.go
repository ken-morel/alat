package app

import (
	"alat/apps/desktop/app/config"
	"fmt"
)

func (app *App) WasSetup() bool {
	fmt.Println("Getting setup state", config.Ready)
	return config.Ready
}

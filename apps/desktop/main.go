package main

import (
	"alat/apps/desktop/app"
)

func main() {
	appConfig.Load()
	autostartCheck() // exits
	n := createNode()
	conf, _ := n.GetAppConfig()
	if conf != nil && conf.SetupComplete {
		n.Start()
	}

	a := app.NewApp(assets, n)

	n.OnPairRequest(a.HandlePairRequest)

	showTray(n, a)
	a.Run()
}

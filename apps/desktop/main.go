package main

import (
	"alat/apps/desktop/app"
	"log"
)

func main() {
	println("Alat starting, loading configuration")
	if err := appConfig.Load(); err != nil {
		println("note loading configuration: ", err)
	}
	if err := appConfig.Save(); err != nil {
		println("Error saving configuration after loading: ", err)
	}
	// autostartCheck() // exits
	println("Creating node")
	n, err := createNode()

	if err != nil {
		log.Fatal("Error creating node", err)
	}
	conf, _ := n.GetAppConfig()
	if conf != nil && conf.SetupComplete {
		n.Start()
	}
	a := app.NewApp(assets, n)
	showTray(n, a)
	a.Run()
}

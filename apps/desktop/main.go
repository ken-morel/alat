package main

import (
	"alat/apps/desktop/app"
	"embed"
)

//go:embed all:frontend/build
var assets embed.FS

func main() {
	app := app.NewApp(assets)

	if err := app.Run(); err != nil {
		println("Error:", err.Error())
	}
}

package main

import (
	"alat/apps/desktop/app"
	"embed"
)

//go:embed all:frontend/build
var assets embed.FS

func main() {
	err := app.NewApp(assets).Run()
	if err != nil {
		println("Error:", err.Error())
	}
}

package main

import "embed"

//go:embed all:frontend/build
var assets embed.FS

//go:embed logos/tray.png
var iconPngData []byte

//go:embed logos/tray-x.png
var iconXPngData []byte

//go:embed logos/tray.ico
var iconIcoData []byte

//go:embed logos/tray-x.ico
var iconXIcoData []byte

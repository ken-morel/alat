package app

import "github.com/wailsapp/wails/v2/pkg/menu"

func (app *App) menu() *menu.Menu {
	m := menu.NewMenu()
	m.Append(menu.AppMenu())
	return m
}

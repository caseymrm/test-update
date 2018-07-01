package main

import (
	"github.com/caseymrm/menuet"
)

var version = "v0.1"

func main() {
	app := menuet.App()
	app.Name = "Updater"
	app.AutoUpdate.Version = version
	app.AutoUpdate.Repo = "caseymrm/test-update"
	menuet.App().SetMenuState(&menuet.MenuState{
		Title: "Version " + version,
	})
	app.RunApplication()
}

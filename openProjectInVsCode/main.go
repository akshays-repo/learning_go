package main

import (
	"embed"
	"io/ioutil"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	bytes, _ := ioutil.ReadFile("./appi.png")

	// Create application with options
	err := wails.Run(&options.App{
		Title:            "openProjectInVsCode",
		Width:            1024,
		Height:           768,
		Assets:           assets,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,

		Bind: []interface{}{
			app,
		},
		Linux: &linux.Options{
			Icon: bytes,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

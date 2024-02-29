package main

import (
	"embed"
	"os"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
)

//go:embed all:frontend/build
var assets embed.FS

var homeDir, _ = os.UserHomeDir()

func main() {
	// Create an instance of the app structure
	initLogger()
	print("Init")
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "nvim-wails",
		Width:     1008,
		Height:    800,
		MinWidth:  1008,
		MinHeight: 80,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 65, G: 69, B: 89, A: 128},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		AlwaysOnTop: true,
		Linux: &linux.Options{
			WebviewGpuPolicy:    linux.WebviewGpuPolicyNever,
			WindowIsTranslucent: true,
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}

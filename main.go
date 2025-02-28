package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/wailsapp/wails/v2/pkg/application"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"log"
	"os"
	"runtime/pprof"
	"terminal/logic"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	logicApp := logic.NewApp()
	app := application.NewWithOptions(&options.App{
		Title:     "WuTerm",
		Width:     924,
		Height:    520,
		MinWidth:  800,
		MinHeight: 600,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Frameless:        true,
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 0},
		OnStartup: func(ctx context.Context) {
			runtime.WindowSetDarkTheme(ctx)
			logicApp.Ctx = ctx
		},
		Bind: []interface{}{
			logicApp,
		},
		Windows: &windows.Options{
			WebviewIsTransparent:              false,
			WindowIsTranslucent:               false,
			BackdropType:                      windows.Mica,
			DisableWindowIcon:                 false,
			DisableFramelessWindowDecorations: false,
			WebviewUserDataPath:               "",
			WebviewBrowserPath:                "",
			Theme:                             windows.Dark,
			ResizeDebounceMS:                  10,
			WebviewGpuIsDisabled:              false,
		},
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
				HideTitle:                  false,
				HideTitleBar:               false,
				FullSizeContent:            false,
				UseToolbar:                 false,
				HideToolbarSeparator:       true,
			},
			Appearance:           mac.NSAppearanceNameDarkAqua,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  false,
			About: &mac.AboutInfo{
				Title:   "WuTerm",
				Message: "© 2025 Wuly",
				//Icon:    icon,
			},
		},
		Linux: &linux.Options{
			//Icon: icon,
			WindowIsTranslucent: true,
			WebviewGpuPolicy:    linux.WebviewGpuPolicyAlways,
			ProgramName:         "WuTerm",
		},
	})
	//startPprof()
	err := app.Run()
	if err != nil {
		println("Error:", err.Error())
	}
}

// startPprof
func startPprof() (func(), bool) {
	enablePprof := flag.Bool("pprof", false, "enable pprof")
	flag.Parse()
	if *enablePprof {
		fmt.Println("starting pprof")
		f, _ := os.OpenFile("cpu.pprof", os.O_CREATE|os.O_RDWR, 0644)
		err := pprof.StartCPUProfile(f)
		if err != nil {
			log.Fatal(err)
		}
		cancel := func() {
			_ = f.Close()
			pprof.StopCPUProfile()
		}
		return cancel, true
	}
	return nil, false
}

package main

import (
	"flag"

	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/asticode/go-astilog"
	"github.com/pkg/errors"
)

// Vars
var (
	AppName string
	BuiltAt string
	debug   = flag.Bool("d", true, "if yes, the app is in debug mode")
	window  *astilectron.Window
)

func main() {
	// Init
	flag.Parse()
	astilog.FlagInit()

	// Run bootstrap
	if err := bootstrap.Run(bootstrap.Options{
		Asset: Asset,
		AstilectronOptions: astilectron.Options{
			AppName:            AppName,
			AppIconDarwinPath:  "resources/oneclient.icns",
			AppIconDefaultPath: "resources/oneclient.png",
		},
		Debug:    *debug,
		Homepage: "index.html",
		MenuOptions: []*astilectron.MenuItemOptions{},
		MessageHandler: handleMessages,
		OnWait: func(_ *astilectron.Astilectron, w *astilectron.Window, _ *astilectron.Menu, t *astilectron.Tray, _ *astilectron.Menu) error {
			// Store global variables
			window = w
			// Add listeners on tray
			t.On(astilectron.EventNameTrayEventClicked, func(e astilectron.Event) (deleteListener bool) { astilog.Info("Tray has been clicked!"); return })
			return nil
		},
		RestoreAssets: RestoreAssets,
		TrayOptions: &astilectron.TrayOptions{
			Image:   astilectron.PtrStr("resources/tray.png"),
			Tooltip: astilectron.PtrStr("OneClient Tray"),
		},
		WindowOptions: &astilectron.WindowOptions{
			BackgroundColor: astilectron.PtrStr("#333"),
			Center:          astilectron.PtrBool(true),
			Height:          astilectron.PtrInt(600),
			Width:           astilectron.PtrInt(600),
		},
	}); err != nil {
		astilog.Fatal(errors.Wrap(err, "Running bootstrap failed"))
	}
}

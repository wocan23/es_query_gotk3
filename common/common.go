package common

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/andlabs/ui"
)

var (
	GlobalWin *gtk.ApplicationWindow
	MainTab ui.Tab

	CurrentConn string
	CurrentIndex string
	CurrentDoc string
)

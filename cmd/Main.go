package main

import (
	"../component"
	"github.com/gotk3/gotk3/gtk"
	"log"
	"os"
	"github.com/gotk3/gotk3/glib"
	"../common"
)

func main(){
	mainFunc()
}

func mainFunc()  {
	// Create Gtk Application, change appID to your application domain name reversed.
	const appID = "org.gtk.es"
	application, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	// Check to make sure no errors when creating Gtk Application
	if err != nil {
		log.Fatal("Could not create application.", err)
	}
	application.Connect("activate", func() { onActivate(application) })
	os.Exit(application.Run(os.Args))
}

func onActivate(application *gtk.Application) {
	// Create ApplicationWindow
	appWindow, err := gtk.ApplicationWindowNew(application)
	if err != nil {
		log.Fatal("Could not create application window.", err)
	}
	// Set ApplicationWindow Properties
	appWindow.SetTitle("Basic Application.")
	appWindow.SetDefaultSize(common.WindowWidth, common.WindowHeight)
	appWindow.Add(Layout())
	appWindow.Show()
}


func Layout() *gtk.Box{

	box,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)
	// heaer
	box.Add(component.CreateHeader())
	// left
	hbox,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)
	subBox,_ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,0)

	left := component.CreateLeft()
	subBox.Add(left)
	// main
	subBox.Add(component.CreateMain())

	subBox.SetMarginTop(10)
	subBox.SetVExpand(true)
	subBox.SetMarginBottom(10)
	hbox.Add(subBox)

	box.Add(hbox)
	box.SetVExpand(true)
	box.SetMarginBottom(10)
	box.SetMarginTop(10)
	box.SetMarginStart(10)
	box.ShowAll()


	return box
}





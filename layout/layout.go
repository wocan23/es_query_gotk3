package layout

import (


	//"../common"
	//"../helper"
	//"../component"
	"github.com/gotk3/gotk3/gtk"
	"log"
	"github.com/gotk3/gotk3/glib"
	"os"
)


func Start(){

	gtk.Init(nil)
	const appID = "org.gtk.example"
	application, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	// Check to make sure no errors when creating Gtk Application
	if err != nil {
		log.Fatal("Could not create application.", err)
	}

	application.Connect("activate", func() { onActivate(application) })
	// Run Gtk application
	os.Exit(application.Run(os.Args))
}

// Callback signal from Gtk Application
func onActivate(application *gtk.Application) {
	// Create ApplicationWindow
	appWindow, err := gtk.ApplicationWindowNew(application)
	if err != nil {
		log.Fatal("Could not create application window.", err)
	}
	// Set ApplicationWindow Properties
	appWindow.SetTitle("Basic Application.")
	appWindow.SetDefaultSize(400, 400)
	appWindow.Show()
}



//func CreateHeader() *gtk.Fixed{
//
//	header := gtk.NewFixed()
//	header.SetSizeRequest(common.WindowWidth,common.NavItemHeight)
//
//	hBox := gtk.NewHBox(true,0)
//
//	connItem := component.GenerateNavItem(common.NavItemWidth,common.NavItemHeight,"/Users/zhaoshuai/Documents/go_workspace_wocan/es_query/images/conn.png","conn")
//	docItem := component.GenerateNavItem(common.NavItemWidth,common.NavItemHeight,"/Users/zhaoshuai/Documents/go_workspace_wocan/es_query/images/doc.png","doc")
//	indexItem := component.GenerateNavItem(common.NavItemWidth,common.NavItemHeight,"/Users/zhaoshuai/Documents/go_workspace_wocan/es_query/images/index.png","index")
//	editItem := component.GenerateNavItem(common.NavItemWidth,common.NavItemHeight,"/Users/zhaoshuai/Documents/go_workspace_wocan/es_query/images/editDoc.png","edit")
//	addItem := component.GenerateNavItem(common.NavItemWidth,common.NavItemHeight,"/Users/zhaoshuai/Documents/go_workspace_wocan/es_query/images/addDoc.png","add")
//	searchItem := component.GenerateNavItem(common.NavItemWidth,common.NavItemHeight,"/Users/zhaoshuai/Documents/go_workspace_wocan/es_query/images/search.png","search")
//
//	hBox.PackStart(connItem,false,false,5)
//	hBox.PackStart(docItem,false,false,5)
//	hBox.PackStart(indexItem,false,false,5)
//	hBox.PackStart(editItem,false,false,5)
//	hBox.PackStart(addItem,false,false,5)
//	hBox.PackStart(searchItem,false,false,5)
//
//	header.Put(hBox,20,10)
//	header.ShowAll()
//	return header
//}
//
//
//func CreateLeft() *gtk.VBox{
//	left := gtk.NewVBox(false,0)
//	left.SetSizeRequest(common.WindowLeftWidth,common.WindowLeftHeight)
//	return left
//}
//
//func CreateMain() *gtk.VBox{
//
//	left := gtk.NewVBox(false,0)
//	left.SetSizeRequest(common.WindowWidth-common.WindowLeftWidth,common.WindowLeftHeight)
//	return left
//}

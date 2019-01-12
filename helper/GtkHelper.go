package helper

import (
	"log"
	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/gdk"
)

func CreateWindow(title string,width int,height int) *gtk.Window{
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Simple Example")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	win.SetDefaultSize(width, height)
	win.SetTitle(title)

	win.ShowAll()

	return win

}

func CreateImage(width int,height int,imagePath string) *gtk.Image{
	spixbuf,_ := gdk.PixbufNewFromFile(imagePath)
	pixbuf,_ := spixbuf.ScaleSimple(width,height,gdk.INTERP_HYPER)
	image,_ := gtk.ImageNewFromPixbuf(pixbuf)
	return image
}

func SetSize(widget gtk.IWidget,width int, height int){
	widget.Set("width",width)
	widget.Set("height",height)
}

func ChangeBgColor(id string,widget *gtk.Widget,rgb string){
		// 染色
	cssProvider,_ := gtk.CssProviderNew()
	cssProvider.LoadFromData("."+id+"{ background-color:"+rgb+";}")
	screen,_ := widget.GetScreen()
	style,_ := widget.GetStyleContext()
	gtk.AddProviderForScreen(screen,cssProvider,1)

	style.AddClass(id)
}

func ChangeMeunBgColor(id string,widget *gtk.Menu,rgb string){
	// 染色
	cssProvider,_ := gtk.CssProviderNew()
	cssProvider.LoadFromData("."+id+"{ background-color:"+rgb+";}")
	screen,_ := widget.GetScreen()
	style,_ := widget.GetStyleContext()
	gtk.AddProviderForScreen(screen,cssProvider,3)

	style.AddClass(id)
}


func RemoveAndAddNew(curBox *gtk.Box,newBox *gtk.Box){
	curBox.GetChildren().Foreach(func(item interface{}) {
		ie := item.(*gtk.Widget)
		curBox.Remove(ie)
	})
	curBox.Add(newBox)
	curBox.ShowAll()
}


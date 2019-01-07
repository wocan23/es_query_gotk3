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


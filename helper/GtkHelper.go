package helper

import (
	"log"
	"github.com/gotk3/gotk3/gtk"
)

func CreateWindow(title string,width int,height int) *gtk.Window{
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Simple Example")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	win.SetSizeRequest(width, height)
	win.ShowAll()

	return win
}
//
//func CreateImage(width int,height int,imagePath string) *gtk.Image{
//	//srcpixBuf,_ := gdkpixbuf.NewPixbufFromFile(imagePath)
//	//pixBuf := srcpixBuf.ScaleSimple(width,width,gdkpixbuf.INTERP_TILES)
//	image := gtk.ImageNewFromPixbuf()
//	return image
//}
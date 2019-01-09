package component

import (
	"../common"
	"github.com/gotk3/gotk3/gtk"
)


func CreateLeft()*gtk.Box{
	box,_ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,0)
	box.SetSizeRequest(common.WindowLeftWidth,common.WindowLeftHeight)

	box.Add(CreateLeftDetail())

	return box
}

func CreateLeftDetail()*gtk.Box{
	box,_ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,0)



	//cssProvider,_ := gtk.CssProviderNew()
	//cssProvider.LoadFromData(`.left{
	//		width:200;
	//	}`)
	//screen,_ := box.GetScreen()
	//style,_ := box.GetStyleContext()
	//gtk.AddProviderForScreen(screen,cssProvider,1)
	//
	//style.AddClass("left")
	//tree := component.CreateTreeC()
	tree := TreeTest()
	box.Add(tree)

	return box
}

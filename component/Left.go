package component

import (
	"../common"
	"github.com/gotk3/gotk3/gtk"
)


func CreateLeft()*gtk.Box{
	box,_ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,0)
	//box.SetSizeRequest(common.WindowLeftWidth,common.WindowLeftHeight)

	adjust,_ := gtk.AdjustmentNew(1,200,400,1,1,100)
	//scrollbar,_ := gtk.ScrollbarNew(gtk.ORIENTATION_VERTICAL,adjust)

	win,_ := gtk.ScrolledWindowNew(adjust,adjust)
	win.SetSizeRequest(common.WindowLeftWidth,common.WindowLeftHeight)

	win.Add(CreateLeftDetail())
	box.Add(win)

	return box
}

func CreateLeftDetail()*gtk.Box{
	box,_ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,5)


	//helper.ChangeBgColor(box,"#34f901")

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

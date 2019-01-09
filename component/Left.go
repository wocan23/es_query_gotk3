package component

import (
	"../common"
	"github.com/gotk3/gotk3/gtk"
)


func CreateLeft()*gtk.Box{
	box,_ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,0)

	adjust,_ := gtk.AdjustmentNew(common.LeftScrollInital,common.LeftScrollLower,common.LeftScrollUpper,common.LeftScrollStepIncrement,common.LeftScrollPageIncrement,common.LeftScrollPageSize)
	//scrollbar,_ := gtk.ScrollbarNew(gtk.ORIENTATION_VERTICAL,adjust)

	win,_ := gtk.ScrolledWindowNew(nil,adjust)
	//win,_ := gtk.ScrolledWindowNew(adjust,adjust)
	win.SetSizeRequest(common.WindowLeftWidth,common.WindowLeftHeight)

	win.Add(CreateLeftDetail())
	//box.Add(scrollbar)
	box.Add(win)

	return box
}

func CreateLeftDetail()*gtk.Box{
	box,_ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,0)

	tree := TreeTest()
	box.Add(tree)

	return box
}

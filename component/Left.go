package component

import (
	"../common"
	"github.com/gotk3/gotk3/gtk"
)


func CreateLeft()*gtk.Box{
	box,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)

	box.SetSizeRequest(common.WindowLeftWidth,common.WindowLeftHeight)

	box.SetBorderWidth(2)

	//tree := component.CreateTreeC()
	box.Add(TreeTest())

	//img := helper.CreateImage(common.WindowLeftWidth,common.WindowLeftHeight,"/Users/zhaoshuai/Documents/go_workspace_wocan/es_query_gotk3/images/index.png")
	//box.Add(img)

	box.SetMarginStart(20)

	box.ShowAll()
	return box
}

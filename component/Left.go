package component

import (
	//"../common"
	"github.com/gotk3/gotk3/gtk"
)


func CreateLeft()*gtk.Box{
	box,_ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,0)


	//box.SetSizeRequest(common.WindowLeftWidth,common.WindowLeftHeight)
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
	tree.SetVExpand(true)
	box.Add(tree)

	//img := helper.CreateImage(common.WindowLeftWidth,common.WindowLeftHeight,"/Users/zhaoshuai/Documents/go_workspace_wocan/es_query_gotk3/images/index.png")
	//box.Add(img)

	box.SetMarginStart(20)
	box.SetMarginBottom(20)
	box.SetVExpand(true)
	return box
}

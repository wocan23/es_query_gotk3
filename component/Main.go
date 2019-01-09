package component

import (
	"github.com/gotk3/gotk3/gtk"
	"../common"
)



func CreateMain()*gtk.Box{
	box,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)

	//box.SetSizeRequest(common.WindowWidth-common.WindowLeftWidth,common.WindowLeftHeight)

	box.SetBorderWidth(2)

	box.Add(CreateMainDetail())

	box.SetVExpand(true)
	box.ShowAll()
	return box
}

func CreateMainDetail() *gtk.Box{
	mainBox,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)
	//width,_ := mainBox.GetSizeRequest()


	// 输入区域
	text,_ := gtk.TextViewNew()
	text.SetVExpand(true)

	//text.SetSizeRequest(-1,common.ShowHeight)

	// 展示区域
	showBox,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)
	// 列表

	scrollbar,_ := gtk.ScrollbarNew(gtk.ORIENTATION_VERTICAL,nil)
	showBox.Add(scrollbar)

	// 工具条
	bar,_ :=  gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,common.ShowBarSpace)

	btn,_ := gtk.ButtonNewWithLabel("查询")
	bar.Add(btn)


	showBox.Add(bar)

	mainBox.Add(text)
	mainBox.Add(showBox)
	mainBox.SetVExpand(true)

	return mainBox
}
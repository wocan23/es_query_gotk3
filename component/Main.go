package component

import (
	"github.com/gotk3/gotk3/gtk"
	"../common"
)



func CreateMain()*gtk.Box{
	box,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)

	box.Add(CreateMainDetail())

	box.ShowAll()
	return box
}

func CreateMainDetail() *gtk.Box{
	mainBox,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)

	// 输入区域
	text,_ := gtk.TextViewNew()

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

	return mainBox
}
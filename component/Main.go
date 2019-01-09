package component

import (
	"github.com/gotk3/gotk3/gtk"
	"../common"
	//"../helper"
)



func CreateMain()*gtk.Box{
	box,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)
	box.SetHExpand(true)
	box.SetMarginEnd(10)
	box.Add(CreateMainDetail())

	box.ShowAll()
	return box
}

func CreateMainDetail() *gtk.Box{

	// tab页
	

	mainBox,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)
	//helper.ChangeBgColor(mainBox,"#676767")

	// 输入区域
	text,_ := gtk.TextViewNew()
	text.SetSizeRequest(common.WindowWidth-common.WindowLeftWidth,common.MainInputHeight)

	// 展示区域
	adjust,_ := gtk.AdjustmentNew(common.LeftScrollInital,common.LeftScrollLower,common.LeftScrollUpper,common.LeftScrollStepIncrement,common.LeftScrollPageIncrement,common.LeftScrollPageSize)

	showWindow,_ := gtk.ScrolledWindowNew(nil,adjust)
	showBox,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)
	// 列表


	showWindow.Add(showBox)
	// 工具条
	bar,_ :=  gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,common.ShowBarSpace)

	btn,_ := gtk.ButtonNewWithLabel("查询")
	bar.Add(btn)


	showBox.Add(bar)

	mainBox.Add(text)
	mainBox.Add(showWindow)

	return mainBox
}
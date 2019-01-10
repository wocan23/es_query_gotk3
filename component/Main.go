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
	mainBox,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)

	// tab页
	tab := Tab{}



	//helper.ChangeBgColor(mainBox,"#676767")



	tab.AddTab("aaa",TabPage("aaa"))
	tab.AddTab("bbb",TabPage("bbb"))
	tab.AddTab("ccc",TabPage("ccc"))

	mainBox.Add(tab.ToTabBox())

	return mainBox
}

func TabPage(txt string) *gtk.Box{
	tabBox,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)
	// 输入区域
	text,_ := gtk.LabelNew(txt)
	text.SetSizeRequest(common.WindowWidth-common.WindowLeftWidth,common.MainInputHeight)
	text.SetTooltipText(txt)

	// 展示区域
	adjust,_ := gtk.AdjustmentNew(common.LeftScrollInital,common.LeftScrollLower,common.LeftScrollUpper,common.LeftScrollStepIncrement,common.LeftScrollPageIncrement,common.LeftScrollPageSize)

	// show box包含scroll和工具条+—
	showBox,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)
	showWindow,_ := gtk.ScrolledWindowNew(nil,adjust)
	// 列表
	dataBox,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)
	showWindow.Add(dataBox)

	// 工具条
	bar,_ :=  gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,common.ShowBarSpace)

	btn,_ := gtk.ButtonNewWithLabel(txt)
	bar.Add(btn)


	showBox.Add(showWindow)
	showBox.Add(bar)

	tabBox.Add(text)
	tabBox.Add(showBox)
	return tabBox
}
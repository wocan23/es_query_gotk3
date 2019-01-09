package component

import (
	"github.com/gotk3/gotk3/gtk"
	"../helper"
	"../common"
)

func CreateTabItem(imagePath string,text string) (*gtk.Box,*gtk.EventBox,*gtk.Image){
	box,_ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,0)

	eventBox,_ := gtk.EventBoxNew()
	subBox,_ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,0)

	box.Add(helper.CreateImage(5,5,imagePath))

	label,_ := gtk.LabelNew(text)
	label.SetSizeRequest(70,5)

	subBox.Add(label)

	eventBox.Add(subBox)

	box.Add(eventBox)
	image := helper.CreateImage(5,5,"")
	box.Add(image)


	return box,eventBox,image
}

type Tab struct{
	tabText []string
	tabBox []*gtk.Box
	curBoxIndex int
}

func (tab *Tab)CreateTab() *gtk.Box{
	box,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)

	// tab条
	adjust,_ := gtk.AdjustmentNew(common.LeftScrollInital,common.LeftScrollLower,common.LeftScrollUpper,common.LeftScrollStepIncrement,common.LeftScrollPageIncrement,common.LeftScrollPageSize)

	scroll,_ := gtk.ScrolledWindowNew(adjust,nil)

	for _,text := range tab.tabText{
		bar,eventBox,imageBox := CreateTabItem("",text)
		scroll.Add(bar)
		eventBox.Connect("", func(gtk.Box) {})
		imageBox.Connect("", func(gtk.Box) {})
	}

	// 展示区域
	scrollData,_ := gtk.ScrolledWindowNew(nil,adjust)
	scrollData.Add(tab.tabBox[tab.curBoxIndex])

	return box
}

func flushBar(){

}

func (tab *Tab)AddTab(text string, box *gtk.Box){

}
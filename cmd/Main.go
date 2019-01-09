package main

import (
	"../helper"
	"../common"
	"../component"
	"github.com/gotk3/gotk3/gtk"
)

func main(){
	mainFunc()
}

func mainFunc()  {
	win := helper.CreateWindow(common.WindowTitle,common.WindowWidth,common.WindowHeight)
	win.SetVAlign(gtk.ALIGN_CENTER)
	win.SetPosition(gtk.WIN_POS_CENTER)

	layout := Layout(win)
	layout.SetMarginTop(20)
	layout.SetMarginStart(20)
	layout.SetMarginBottom(20)
	layout.SetMarginEnd(20)
	layout.SetVExpand(true)


	win.Add(layout)


	win.Show()
	gtk.Main()
}


func Layout(win *gtk.Window) *gtk.Box{

	box,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)
	// heaer
	box.Add(component.CreateHeader())
	// left
	subBox,_ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,0)

	left := component.CreateLeft()
	subBox.Add(left)
	// main
	subBox.Add(component.CreateMain())

	subBox.SetMarginTop(10)
	subBox.SetVExpand(true)
	subBox.SetMarginBottom(10)

	box.Add(subBox)
	box.SetVExpand(true)
	box.SetMarginBottom(0)
	box.ShowAll()


	//win.Connect("configure_event", func(win *gtk.Window) {
	//
	//	_,h := win.GetSize()
	//	//_,height := subBox.GetSizeRequest()
	//	width,_ := left.GetSizeRequest()
	//	fmt.Println(h)
	//	//fmt.Println(height)
	//	sub := common.WindowHeight -22
	//	left.SetSizeRequest(width,sub)
	//})


	return box
}





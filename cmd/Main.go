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
	win.Add(Layout())
	win.SetVExpand(true)
	win.Show()
	gtk.Main()
}


func Layout() *gtk.Box{

	box,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)
	// heaer
	box.Add(component.CreateHeader())
	// left
	subBox,_ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,0)

	subBox.Add(component.CreateLeft())
	// main
	subBox.Add(component.CreateMain())

	subBox.SetMarginTop(10)
	subBox.SetVExpand(true)
	subBox.SetMarginBottom(0)

	box.Add(subBox)
	box.SetVExpand(true)
	box.SetMarginBottom(0)
	box.ShowAll()
	return box
}





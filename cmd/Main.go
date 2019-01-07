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
	win.Add(Layout())
	win.Show()
	gtk.Main()
}


func Layout() *gtk.Box{
	// heaer
	box,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)
	box.Add(CreateHeader())

	// left
	subBox,_ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,0)

	subBox.Add(CreateLeft())
	subBox.Add(CreateMain())

	// main

	box.Add(subBox)
	box.ShowAll()
	return box
}

func CreateHeader() *gtk.Box{
	box,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)

	menu,_ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,5)


	connBtn := component.CreateNavItem(common.NavItemWidth,common.NavItemHeight,"conn","/Users/zhaoshuai/Documents/go_workspace_wocan/es_query_gotk3/images/conn.png")
	indexBtn := component.CreateNavItem(common.NavItemWidth,common.NavItemHeight,"index","/Users/zhaoshuai/Documents/go_workspace_wocan/es_query_gotk3/images/index.png")
	docBtn := component.CreateNavItem(common.NavItemWidth,common.NavItemHeight,"doc","/Users/zhaoshuai/Documents/go_workspace_wocan/es_query_gotk3/images/doc.png")
	editBtn := component.CreateNavItem(common.NavItemWidth,common.NavItemHeight,"edit","/Users/zhaoshuai/Documents/go_workspace_wocan/es_query_gotk3/images/editDoc.png")
	addBtn := component.CreateNavItem(common.NavItemWidth,common.NavItemHeight,"add","/Users/zhaoshuai/Documents/go_workspace_wocan/es_query_gotk3/images/addDoc.png")
	searchBtn := component.CreateNavItem(common.NavItemWidth,common.NavItemHeight,"search","/Users/zhaoshuai/Documents/go_workspace_wocan/es_query_gotk3/images/search.png")

	menu.Add(connBtn)
	menu.Add(indexBtn)
	menu.Add(docBtn)
	menu.Add(editBtn)
	menu.Add(addBtn)
	menu.Add(searchBtn)

	box.Add(menu)
	box.SetMarginStart(20)
	box.SetMarginTop(20)
	box.ShowAll()
	return box
}

func CreateLeft()*gtk.Box{
	box,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)

	box.SetSizeRequest(common.WindowLeftWidth,common.WindowLeftHeight)

	box.SetBorderWidth(2)

	tree := component.CreateTree()
	box.Add(tree)

	//img := helper.CreateImage(common.WindowLeftWidth,common.WindowLeftHeight,"/Users/zhaoshuai/Documents/go_workspace_wocan/es_query_gotk3/images/index.png")
	//box.Add(img)

	box.SetMarginStart(20)

	box.ShowAll()
	return box
}

func CreateMain()*gtk.Box{
	box,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)

	box.SetSizeRequest(common.WindowWidth-common.WindowLeftWidth,common.WindowLeftHeight)

	box.SetBorderWidth(2)

	img := helper.CreateImage(common.WindowLeftWidth,common.WindowLeftHeight,"/Users/zhaoshuai/Documents/go_workspace_wocan/es_query_gotk3/images/doc.png")
	box.Add(img)

	box.SetMarginStart(20)

	box.ShowAll()
	return box
}

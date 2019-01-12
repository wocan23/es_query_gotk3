package component

import (
	"../common"
	//"../helper"
	"github.com/gotk3/gotk3/gtk"
)

var connFlag = false

func CreateHeader() *gtk.Box{
	box,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)

	menu,_ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,5)


	connBtn := CreateNavItem(common.NavItemWidth,common.NavItemHeight,"conn",common.ConnImagePath)
	indexBtn := CreateNavItem(common.NavItemWidth,common.NavItemHeight,"index",common.IndexImagePath)
	docBtn := CreateNavItem(common.NavItemWidth,common.NavItemHeight,"doc",common.DocImagePath)
	editBtn := CreateNavItem(common.NavItemWidth,common.NavItemHeight,"edit",common.EditDocImagePath)
	addBtn := CreateNavItem(common.NavItemWidth,common.NavItemHeight,"add",common.AddDocImagePath)
	searchBtn := CreateNavItem(common.NavItemWidth,common.NavItemHeight,"search",common.SearchDocImagePath)

	connBtn.Connect("button_press_event",connClickCallback)

	menu.Add(connBtn)
	menu.Add(indexBtn)
	menu.Add(docBtn)
	menu.Add(editBtn)
	menu.Add(addBtn)
	menu.Add(searchBtn)

	//helper.ChangeBgColor(menu,"#ff00f0")
	//helper.ChangeBgColor(box,"#00ff00")

	box.Add(menu)
	box.ShowAll()
	return box
}

func connClickCallback(btn *gtk.Button){

	connWin,_ :=  gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	connWin.SetKeepAbove(true)
	connWin.SetFocusVisible(true)
	connWin.SetPosition(gtk.WIN_POS_CENTER)
	connWin.SetTitle("create connection")
	connWin.SetTransientFor(common.GlobalWin)
	connWin.SetModal(true)
	connWin.SetSizeRequest(common.ConnWindowWidth,common.ConnWindowHeight)

	label,_ := gtk.LabelNew("test")
	box,_ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,0)
	box.Add(label)
	connWin.Add(box)

	connWin.ShowAll()


}


package component

import (
	"../common"
	//"../helper"
	"github.com/gotk3/gotk3/gtk"
	"fmt"
)

var connFlag = false

func CreateHeader() *gtk.Box{
	box,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)

	menu,_ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,5)


	connBtn := CreateNavItem(common.NavItemWidth,common.NavItemHeight,"conn","/Users/zhaoshuai/Documents/go_workspace_wocan/es_query_gotk3/images/conn.png")
	indexBtn := CreateNavItem(common.NavItemWidth,common.NavItemHeight,"index","/Users/zhaoshuai/Documents/go_workspace_wocan/es_query_gotk3/images/index.png")
	docBtn := CreateNavItem(common.NavItemWidth,common.NavItemHeight,"doc","/Users/zhaoshuai/Documents/go_workspace_wocan/es_query_gotk3/images/doc.png")
	editBtn := CreateNavItem(common.NavItemWidth,common.NavItemHeight,"edit","/Users/zhaoshuai/Documents/go_workspace_wocan/es_query_gotk3/images/editDoc.png")
	addBtn := CreateNavItem(common.NavItemWidth,common.NavItemHeight,"add","/Users/zhaoshuai/Documents/go_workspace_wocan/es_query_gotk3/images/addDoc.png")
	searchBtn := CreateNavItem(common.NavItemWidth,common.NavItemHeight,"search","/Users/zhaoshuai/Documents/go_workspace_wocan/es_query_gotk3/images/search.png")

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
	fmt.Println("click conn")

	//connWin,_ :=  gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	//connWin.SetKeepAbove(true)
	//connWin.SetPosition(gtk.WIN_POS_CENTER)
	//connWin.SetDestroyWithParent(true)
	//connWin.SetVisible(true)
	//connWin.Show()

	//dialog,_ := gtk.DialogNew()

	//common.GlobalWin.SetKeepAbove(true)
	//dialog.SetKeepAbove(true)
	//dialog.SetPosition(gtk.WIN_POS_CENTER)
	//dialog.SetSizeRequest(200,200)
	//dialog.SetTitle("新建连接")
	//dialog.SetVisible(true)
	//win,_ := btn.GetWindow()
	//dialog.SetSensitive(true)
	//dialog.SetParentWindow(win)
	//dialog.Show()

}


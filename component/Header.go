package component

import (
	"../common"
	"github.com/gotk3/gotk3/gtk"
)

func CreateHeader() *gtk.Box{
	box,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)

	menu,_ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,5)


	connBtn := CreateNavItem(common.NavItemWidth,common.NavItemHeight,"conn","/Users/zhaoshuai/Documents/go_workspace_wocan/es_query_gotk3/images/conn.png")
	indexBtn := CreateNavItem(common.NavItemWidth,common.NavItemHeight,"index","/Users/zhaoshuai/Documents/go_workspace_wocan/es_query_gotk3/images/index.png")
	docBtn := CreateNavItem(common.NavItemWidth,common.NavItemHeight,"doc","/Users/zhaoshuai/Documents/go_workspace_wocan/es_query_gotk3/images/doc.png")
	editBtn := CreateNavItem(common.NavItemWidth,common.NavItemHeight,"edit","/Users/zhaoshuai/Documents/go_workspace_wocan/es_query_gotk3/images/editDoc.png")
	addBtn := CreateNavItem(common.NavItemWidth,common.NavItemHeight,"add","/Users/zhaoshuai/Documents/go_workspace_wocan/es_query_gotk3/images/addDoc.png")
	searchBtn := CreateNavItem(common.NavItemWidth,common.NavItemHeight,"search","/Users/zhaoshuai/Documents/go_workspace_wocan/es_query_gotk3/images/search.png")

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

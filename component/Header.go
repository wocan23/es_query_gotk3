package component

import (
	"../common"
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

	// 新建连接
	connBtn.Connect("button_press_event",connClickCallback)

	menu.Add(connBtn)
	menu.Add(indexBtn)
	menu.Add(docBtn)
	menu.Add(editBtn)
	menu.Add(addBtn)
	menu.Add(searchBtn)



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
	connWin.SetTransientFor(GlobalWin)
	connWin.SetModal(true)
	connWin.SetSizeRequest(common.ConnWindowWidth,common.ConnWindowHeight)

	box,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)

	nameBox,nameEntry := createFormItem("name")
	ipBox,ipEntry := createFormItem("ip")
	portBox,portEntry := createFormItem("port")

	formItemBox,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,10)
	formItemBox.Add(nameBox)
	formItemBox.Add(ipBox)
	formItemBox.Add(portBox)
	formItemBox.SetMarginTop(20)
	formItemBox.SetMarginStart(20)


	commitBtn,_ := gtk.ButtonNewWithLabel("确定")
	commitBtn.SetMarginTop(10)
	commitBtn.SetSizeRequest(30,20)
	commitBtn.Connect("button_press_event", func() {
		ip,_ := ipEntry.GetText()
		port,_ := portEntry.GetText()
		name,_ := nameEntry.GetText()
		if ip =="" {
			dialog := gtk.MessageDialogNew(connWin,gtk.DIALOG_MODAL,gtk.MESSAGE_INFO,gtk.BUTTONS_CLOSE,"ip不能为空")
			dialog.Show()
			return
		}
		if port =="" {
			dialog := gtk.MessageDialogNew(connWin,gtk.DIALOG_MODAL,gtk.MESSAGE_INFO,gtk.BUTTONS_CLOSE,"port不能为空")
			dialog.Show()
			return
		}
		if name =="" {
			dialog := gtk.MessageDialogNew(connWin,gtk.DIALOG_MODAL,gtk.MESSAGE_INFO,gtk.BUTTONS_CLOSE,"name不能为空")
			dialog.ShowAll()
			return
		}
		CreateEsConn(name,ip+":"+port)
		// todo user action
		connWin.Destroy()
	})

	box.Add(formItemBox)
	box.Add(commitBtn)
	connWin.Add(box)

	connWin.ShowAll()

}

// 新建连接
func CreateEsConn(name string,url string){
	connNode := CreateNode(name,common.ConnImagePath)
	connNode.SetProp("conn",url)
	LeftTree.AddNode(connNode)
}

func createFormItem(labelText string)(*gtk.Box,*gtk.Entry){
	box,_ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,0)

	label,_ := gtk.LabelNew(labelText)
	label.SetWidthChars(10)
	labelValue,_ := gtk.EntryNew()
	box.Add(label)
	box.Add(labelValue)

	return box,labelValue
}




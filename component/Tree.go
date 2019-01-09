package component

import (
	"github.com/gotk3/gotk3/gtk"
	"../helper"
	"../common"
	"C"
)

type TreeData struct{
	item *TreeItemData // 当前数据
	subItems []*TreeData //子数据
	isSubShow bool // 是否展示子标签
	isRoot bool

	parent *TreeData
}

type TreeItemData struct{
	data map[string]string
}



func (treeData *TreeData)SetProperty(k string,v string){
	if treeData.item == nil{
		treeData.item = new(TreeItemData)
		treeData.item.data = make(map[string]string,0)
	}
	treeData.item.data[k] = v
}

func (treeData *TreeData)AddSubItems(subTreeData *TreeData){
	if treeData.subItems == nil{
		treeData.subItems = make([]*TreeData,0)
	}
	treeData.subItems = append(treeData.subItems, subTreeData)
	subTreeData.parent = treeData
}

func CreateTreeData()*TreeData{
	return new(TreeData)
}

func CreateTreeByData(data *TreeData,getWidget func(data *TreeData,root *gtk.Box,parent *gtk.Box)gtk.IWidget) *gtk.Box {
	box, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	//box.SetSizeRequest(common.WindowLeftWidth,common.WindowLeftHeight)
	box.SetBorderWidth(0)
	flushTree(box,data,box,getWidget)
	return box

}

func flushTree(root *gtk.Box,data *TreeData,parent *gtk.Box,getWidget func(data *TreeData,root *gtk.Box,parent *gtk.Box)gtk.IWidget) {
	// 第一层直接遍历子
	curBox,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)

	widget := getWidget(data,root,curBox)
	curBox.Add(widget)

	parent.Add(curBox)
}

func GetWidget(data *TreeData,root *gtk.Box,current *gtk.Box) gtk.IWidget{
	text := data.item.data["text"]
	imagePath := data.item.data["imagePath"]
	btn,_ := gtk.EventBoxNew()
	btn.SetSizeRequest(100,30)

	btn.SetHExpand(false)

	box,_ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,0)
	var level = checkLevel(data)
	for i := 0; i < level; i++{
		_,height := box.GetSizeRequest()
		label,_ := gtk.LabelNew(" ")
		label.SetSizeRequest(10,height)
		box.Add(label)
	}
	box.SetHExpand(false)

	image := helper.CreateImage(common.TreeItemWidth,common.TreeItemHeight,imagePath)


	textV,_ := gtk.LabelNew(text)
	textV.SetTooltipText(text)


	entity,_ := gtk.EntryNew()
	entity.SetTooltipText(text)
	entity.SetVisible(false)

	box.Add(image)
	box.Add(textV)


	// 点击事件
	btn.Connect("button_press_event", func(e *gtk.EventBox) {
		data.isSubShow = !data.isSubShow
		if data.isSubShow{
			showSubItems(data,root,current)
		}else{
			unShowSubItems(current)
		}
	})

	btn.SetCanFocus(true)

	btn.Connect("enter_notify_event", func(e *gtk.EventBox) {
		changeBgColor(e)
	})

	btn.Connect("leave_notify_event", func(e *gtk.EventBox) {
		clearBgColor(e)
	})

	btn.Add(box)
	btn.ShowAll()

	return btn
}

func changeBgColor(btn *gtk.EventBox){
	// 染色
	cssProvider,_ := gtk.CssProviderNew()
	cssProvider.LoadFromData(`.clicked{
			background-color:#87CEFA;
		}`)
	screen,_ := btn.GetScreen()
	style,_ := btn.GetStyleContext()
	gtk.AddProviderForScreen(screen,cssProvider,1)

	style.AddClass("clicked")
}

func clearBgColor(btn *gtk.EventBox){
	style,_ := btn.GetStyleContext()

	style.RemoveClass("clicked")
}

func showSubItems(data *TreeData,root *gtk.Box,current *gtk.Box){
	for _, e := range data.subItems {
		pBox,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)

		widget := GetWidget(e,root,pBox)

		pBox.Add(widget)
		current.Add(pBox)
	}
	current.ShowAll()
}

func unShowSubItems(current *gtk.Box){
	list := current.GetChildren()
	var length = list.Length()
	var i uint
	for i = 1 ; i < length; i ++{
		item := list.NthData(i)
		widget := item.(*gtk.Widget)
		widget.Destroy()
	}
	current.Show()
}

func checkLevel(data *TreeData)int{
	var i = 0
	parent := data.parent
	for parent != nil{
		parent = parent.parent
		i ++
	}
	return i
}



func TreeTest()*gtk.Box{

	// 数据
	d1 := CreateTreeData()

	d11 := CreateTreeData()
	d12 := CreateTreeData()
	d13 := CreateTreeData()
	d111 := CreateTreeData()

	d121 := CreateTreeData()
	d122 := CreateTreeData()

	d131 := CreateTreeData()
	d132 := CreateTreeData()
	d133 := CreateTreeData()

	d1.SetProperty("text","d1")
	d1.SetProperty("imagePath","/Users/zhaoshuai/Documents/go_workspace_wocan/es_query_gotk3/images/conn.png")
	d11.SetProperty("text","d11")
	d11.SetProperty("imagePath","/Users/zhaoshuai/Documents/go_workspace_wocan/es_query_gotk3/images/index.png")
	d12.SetProperty("text","d12")
	d12.SetProperty("imagePath","/Users/zhaoshuai/Documents/go_workspace_wocan/es_query_gotk3/images/index.png")
	d13.SetProperty("text","d13")
	d13.SetProperty("imagePath","/Users/zhaoshuai/Documents/go_workspace_wocan/es_query_gotk3/images/index.png")
	d111.SetProperty("text","d111")
	d111.SetProperty("imagePath","/Users/zhaoshuai/Documents/go_workspace_wocan/es_query_gotk3/images/doc.png")
	d121.SetProperty("text","d121")
	d121.SetProperty("imagePath","/Users/zhaoshuai/Documents/go_workspace_wocan/es_query_gotk3/images/doc.png")
	d122.SetProperty("text","d122")
	d122.SetProperty("imagePath","/Users/zhaoshuai/Documents/go_workspace_wocan/es_query_gotk3/images/doc.png")
	d131.SetProperty("text","d131")
	d131.SetProperty("imagePath","/Users/zhaoshuai/Documents/go_workspace_wocan/es_query_gotk3/images/doc.png")
	d132.SetProperty("text","d132")
	d132.SetProperty("imagePath","/Users/zhaoshuai/Documents/go_workspace_wocan/es_query_gotk3/images/doc.png")
	d133.SetProperty("text","d133")
	d133.SetProperty("imagePath","/Users/zhaoshuai/Documents/go_workspace_wocan/es_query_gotk3/images/doc.png")

	d11.AddSubItems(d111)

	d12.AddSubItems(d121)
	d12.AddSubItems(d122)

	d13.AddSubItems(d131)
	d13.AddSubItems(d132)
	d13.AddSubItems(d133)

	d1.AddSubItems(d11)
	d1.AddSubItems(d12)
	d1.AddSubItems(d13)
	d1.isRoot = true

	return CreateTreeByData(d1,GetWidget)

}
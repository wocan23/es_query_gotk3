package component

import (
	"github.com/gotk3/gotk3/gtk"
	"C"
	"../helper"
	"../common"
	"github.com/gotk3/gotk3/gdk"
	"fmt"
)
/**
 nodeBox为当前节点最外层box
里面包含：nodeDataBox
 */
type Node struct{
	nodeData *NodeData // 当前数据
	subNodes []*Node //子数据
	isSubShow bool // 是否展示子标签
	nodeBox *gtk.Box // 当前对应的最外层box
	nodeDataBox *gtk.Box // 当前数据box
	nodeSubBox []*gtk.Box // 当前子节点包含层级
	rootBox *gtk.Box // 根box
	nodeSubEditBox *gtk.Box // 编辑box

	parent *Node // 父节点

	nodeLevel int // 节点级别


	tree *Tree // 属于哪颗树

}

type Tree struct{
	nodes []*Node
	rootBox *gtk.Box

	showStatus int // 0默认，1全部展开，2全部关闭
}

type NodeData struct{
	data map[string]string
	labelBox *gtk.Label

	label string
	icon string
}



func (node *Node)SetProp(k string,v string){
	node.nodeData.data[k] = v
}


func CreateNode(label string,icon string)(*Node){
	// 数据区
	node := new(Node)
	node.subNodes = make([]*Node,0)
	node.nodeData = new(NodeData)
	node.nodeData.data = make(map[string]string,0)
	node.nodeData.icon = icon
	node.nodeData.label = label

	// 展示区
	nodeBox,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)
	node.nodeDataBox,_ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,-0)
	node.nodeDataBox.Add( node.CreateNodeDataBox(label,icon))

	nodeBox.Add(node.nodeDataBox)
	node.nodeBox = nodeBox
	return node
}

func CreateTree() (*Tree){
	tree := new(Tree)
	rootBox,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)
	tree.rootBox = rootBox
	tree.nodes = make([]*Node,0)
	return tree
}

// 增加节点
func (tree *Tree)AddNode(node *Node){
	node.isSubShow = false
	node.nodeLevel = 0
	tree.nodes = append(tree.nodes, node)
	tree.rootBox.Add(node.nodeBox)
	tree.rootBox.ShowAll()
}

// 展示子节点
func (node *Node)ShowSub(){
	node.isSubShow = true
	for _,subNodeBox := range node.nodeSubBox{
		node.nodeBox.Add(subNodeBox)
	}
	node.nodeBox.ShowAll()
}

// 关闭子节点
func (node *Node)UnShowSub(){
	node.isSubShow = false
	for _,subNodeBox := range node.nodeSubBox{
		node.nodeBox.Remove(subNodeBox)
	}
	node.nodeBox.ShowAll()
}

// 增加子节点
func (node *Node)AddSubNode(subNode *Node){
	subNode.parent = node
	subNode.rootBox = node.nodeBox
	subNode.nodeLevel = node.nodeLevel + 1

	node.subNodes = append(node.subNodes, subNode)

	// subBox带缩进
	subBox,_ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,0)
	subBox.Add(fullIncidentBox(subNode.nodeBox,subNode.nodeLevel))
	node.nodeSubBox = append(node.nodeSubBox, subBox)

	if node.isSubShow{
		// 缩紧
		node.nodeBox.Add(subBox)
	}
}

// 移除当前节点
func (node *Node)RemoveNode(){

	// 删除父节点存储的任何关联
	parent := node.parent

	// 查询子位置
	var subIndex = -1
	for i,e := range parent.subNodes{
		if e == node{
			subIndex = i
		}
	}

	if subIndex != -1{
		// 移除子节点视图
		parent.nodeSubBox = append(parent.nodeSubBox[:subIndex],parent.nodeSubBox[subIndex+1:]...)
		// 移除子节点
		parent.subNodes = append(parent.subNodes[:subIndex],parent.subNodes[subIndex+1:]...)
	}
	node.parent = nil
	node.tree = nil
	node.nodeData = nil

	node.nodeBox.Destroy()
}


// 编辑已有节点
func (node *Node)EditNodeBox(){
	editBox := node.CreateEditNodeBox(node.nodeData.label,node.nodeData.icon)
	RemoveChildren(node.nodeDataBox)
	node.nodeDataBox.Add(editBox)
}

// 编辑完替换
func (node *Node)EndEditNodeBox(label string,icon string){
	RemoveChildren(node.nodeDataBox)
	node.nodeDataBox.Add(node.CreateNodeDataBox(label,icon))

	// 数据区
	node.nodeData.data["label"] = label
	node.nodeData.label = label
}

// 添加编辑节点
func (node *Node)AddEditNodeBox(label string,icon string){

	editSubNodeBox := node.CreateEditNodeBox(label,icon)

	node.nodeSubEditBox = fullIncidentBox(editSubNodeBox,node.nodeLevel+1)

	node.nodeSubBox = append(node.nodeSubBox, node.nodeSubEditBox)

	node.nodeBox.Add(node.nodeSubEditBox)

	node.nodeBox.ShowAll()
}

// 移除编辑替换节点
func (node *Node)RemoveEditNodeBox(){
	node.nodeBox.Remove(node.nodeSubEditBox)
	node.nodeSubBox = node.nodeSubBox[:len(node.nodeSubBox)-1]
	node.nodeSubEditBox.Destroy()
	node.nodeSubEditBox = nil
}

func (node *Node)CreateEditNodeBox(label string,icon string) *gtk.Box{
	nodeBox,_ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,0)

	image := helper.CreateImage(common.TreeItemWidth,common.TreeItemHeight,icon)
	labelEntity,_ := gtk.EntryNew()
	//labelEntity.SetCanFocus(true)
	labelEntity.SetText(label)
	labelEntity.SetVisibility(true)

	nodeBox.Add(image)
	nodeBox.Add(labelEntity)
	nodeBox.ShowAll()

	//labelEntity.Connect("focus_out_event", func() {
	//	//fmt.Println("test entry event")
	//	newLabel,_ := labelEntity.GetText()
	//	node.EndEditNodeBox(newLabel,icon)
	//	fmt.Println("test")
	//})


	labelEntity.Connect("activate", func() {
		fmt.Println("active")
		newLabel,_ := labelEntity.GetText()
		node.EndEditNodeBox(newLabel,icon)
	})

	return nodeBox
}

func (node *Node)CreateNodeDataBox(label string,icon string)(*gtk.Box){

	nodeBox,_ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,0)
	eventBox,_ := gtk.EventBoxNew()
	nodeBox1,_ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,0)

	image := helper.CreateImage(common.TreeItemWidth,common.TreeItemHeight,icon)
	textLabel,_ := gtk.LabelNew(label)

	nodeBox1.Add(image)
	nodeBox1.Add(textLabel)

	eventBox.Add(nodeBox1)
	nodeBox.Add(eventBox)


	eventBox.Connect("button_press_event", func(box *gtk.EventBox ,event *gdk.Event) {

		eventButton := gdk.EventButtonNewFromEvent(event)

		/**
			1:左键
			2:滚轮
			3:右键
		 */
		switch eventButton.Button() {
		case 3:
			// 鼠标右键
			bindMouseRightButton(node,eventButton)
		case 1:
			// 左键展开/关闭
			if node.isSubShow{
				node.UnShowSub()
			}else{
				node.ShowSub()
			}
		}

	})

	nodeBox.ShowAll()
	return nodeBox
}

func RemoveChildren(box *gtk.Box){
	children := box.GetChildren()
	children.Foreach(func(subBoxI interface{}) {
		subBox := subBoxI.(*gtk.Widget)
		box.Remove(subBox)
	})

}

func fullIncidentBox(box *gtk.Box, num int) *gtk.Box{
	incidentBox,_ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,0)
	for i := 0; i < num; i++{
		ilabel,_ := gtk.LabelNew("  ")
		incidentBox.Add(ilabel)
	}
	incidentBox.Add(box)
	return incidentBox
}



func Tree2Test() *Tree{
	tree := CreateTree()

	nodeA := CreateNode("aaa",common.ConnImagePath)
	nodeB := CreateNode("bbb",common.ConnImagePath)
	nodeC := CreateNode("ccc",common.ConnImagePath)


	nodeA1 := CreateNode("aaa111",common.IndexImagePath)
	nodeA2 := CreateNode("aaa112",common.IndexImagePath)
	nodeA3 := CreateNode("aaa113",common.IndexImagePath)

	nodeB1 := CreateNode("bbb111",common.IndexImagePath)
	nodeB2 := CreateNode("bbb112",common.IndexImagePath)
	nodeB3 := CreateNode("bbb113",common.IndexImagePath)

	nodeA1a := CreateNode("aaa111aaa",common.DocImagePath)
	nodeA1a.SetProp("conn","/")
	nodeA2b := CreateNode("aaa112bbb",common.DocImagePath)
	nodeA3c := CreateNode("aaa113ccc",common.DocImagePath)

	nodeB1a := CreateNode("bbb111aaa",common.DocImagePath)
	nodeB2b := CreateNode("bbb112bbb",common.DocImagePath)
	nodeB3c := CreateNode("bbb113ccc",common.DocImagePath)


	tree.AddNode(nodeA)
	tree.AddNode(nodeB)
	tree.AddNode(nodeC)

	nodeA.AddSubNode(nodeA1)
	nodeA.AddSubNode(nodeA2)
	nodeA.AddSubNode(nodeA3)

	nodeB.AddSubNode(nodeB1)
	nodeB.AddSubNode(nodeB2)
	nodeB.AddSubNode(nodeB3)

	nodeA1.AddSubNode(nodeA1a)
	nodeA2.AddSubNode(nodeA2b)
	nodeA3.AddSubNode(nodeA3c)

	nodeB1.AddSubNode(nodeB1a)
	nodeB2.AddSubNode(nodeB2b)
	nodeB3.AddSubNode(nodeB3c)

	tree.rootBox.ShowAll()
	return tree
}

// 打开鼠标右键
func bindMouseRightButton(node *Node,eventButton *gdk.EventButton){

	rightMenu,_ := gtk.MenuNew()

	helper.ChangeMeunBgColor("menu",rightMenu,"#0f0")

	seeItem,_ := gtk.MenuItemNewWithLabel("see")
	addItem,_ := gtk.MenuItemNewWithLabel("add")
	editItem,_ := gtk.MenuItemNewWithLabel("edit")
	sepItem,_ := gtk.SeparatorMenuItemNew()

	seeItem.Connect("button_press_event", func() {

	})
	addItem.Connect("button_press_event", func() {

	})
	editItem.Connect("button_press_event", func() {
		node.EditNodeBox()
	})

	rightMenu.Append(seeItem)
	rightMenu.Append(addItem)
	rightMenu.Append(sepItem)
	rightMenu.Append(editItem)
	rightMenu.ShowAll()

	rightMenu.PopupAtPointer(eventButton.Event)

}




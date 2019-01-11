package component

import (
	"github.com/gotk3/gotk3/gtk"
	"C"
	"../helper"
	"../common"
)

type Node struct{
	nodeData *NodeData // 当前数据
	subNodes []*Node //子数据
	isSubShow bool // 是否展示子标签
	nodeBox *gtk.Box // 当前对应的最外层box
	nodeDataBox *gtk.Box // 当前数据box
	rootBox *gtk.Box // 根box

	parent *Node // 父节点

	nodeLevel int // 节点级别


	tree *Tree // 属于哪颗树

	isEditNode bool // 是否编辑
}

type Tree struct{
	nodes []*Node
	rootBox *gtk.Box

	showStatus int // 0默认，1全部展开，2全部关闭
}

type NodeData struct{
	data map[string]string
	labelBox *gtk.Label
}



func (node *Node)SetProp(k string,v string){
	node.nodeData.data[k] = v
}


func CreateNode(isEdit bool)(*Node){
	// 数据区
	node := new(Node)
	node.isSubShow = true
	node.subNodes = make([]*Node,0)
	node.nodeData = new(NodeData)
	node.isEditNode = isEdit

	// 展示区
	nodeBox,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)
	node.nodeBox = nodeBox

	return node
}

func CreateTree(nodeFunc func(text string)*gtk.Box) (*Tree){
	tree := new(Tree)
	rootBox,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)
	tree.rootBox = rootBox
	tree.nodes = make([]*Node,0)
	return tree
}

// 增加节点
func (tree *Tree)AddNode(node *Node){
	tree.nodes = append(tree.nodes, node)
	tree.rootBox.Add(node.nodeBox)
	tree.rootBox.ShowAll()
}

// 展示子节点
func (node *Node)ShowSub(){
	node.isSubShow = true
	for _,subNode := range node.subNodes{
		node.nodeBox.Add(subNode.nodeBox)
	}
	node.nodeBox.ShowAll()
}

// 关闭子节点
func (node *Node)UnShowSub(){
	node.isSubShow = false
	for _,subNode := range node.subNodes{
		node.nodeBox.Remove(subNode.nodeBox)
	}
	node.nodeBox.ShowAll()
}

// 增加子节点
func (node *Node)AddSubNode(subNode *Node){
	node.subNodes = append(node.subNodes, subNode)
	if node.isSubShow{
		node.nodeBox.Add(subNode.nodeBox)
	}

	subNode.parent = node
	subNode.rootBox = node.nodeBox
}

// 编辑已有节点
func (node *Node)EditNodeBox(){

}

// 添加编辑节点
func (node *Node)AddEditNodeBox(label string,icon string){

	editNodeBox := CreateEditNodeBox(label,icon)

	node.nodeDataBox.Add(editNodeBox)

	node.nodeBox.ShowAll()
}

func (node *Node)RemoveEditNodeBox(){
}

func CreateEditNodeBox(label string,icon string) *gtk.Box{
	nodeBox,_ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,0)

	image := helper.CreateImage(common.TreeItemWidth,common.TreeItemHeight,icon)
	labelEntity,_ := gtk.EntryNew()
	labelEntity.SetText(label)

	nodeBox.Add(image)
	nodeBox.Add(labelEntity)

	return nodeBox
}

func CreateNodeBox(label string,icon string)*gtk.Box{
	nodeBox,_ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,0)

	image := helper.CreateImage(common.TreeItemWidth,common.TreeItemHeight,icon)
	textLabel,_ := gtk.LabelNew(label)

	nodeBox.Add(image)
	nodeBox.Add(textLabel)

	return nodeBox
}



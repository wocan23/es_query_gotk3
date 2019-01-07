package component

import (
	"github.com/gotk3/gotk3/gtk"
	"../helper"
	"../common"
)

func CreateTreeC(){

}

func CreateTreeItemC(text,imagePath string) *gtk.Button{
	btn,_ := gtk.ButtonNew()

	box,_ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,5)

	image := helper.CreateImage(common.TreeItemWidth,common.TreeItemHeight,imagePath)

	entity,_ := gtk.EntryNew()
	entity.SetEditable(false)

	box.Add(image)

	return btn
}

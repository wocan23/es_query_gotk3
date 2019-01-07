package component

import (
	//"../common"
	//"../helper"
	"github.com/gotk3/gotk3/gtk"
	"../helper"
)

func CreateNavItem(width int,height int,title string,imagePath string) *gtk.Button{

	btn,_ := gtk.ButtonNew()

	box,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)
	box.SetHExpand(false)
	box.SetVExpand(false)

	image := helper.CreateImage(width,width,imagePath)
	box.Add(image)

	label,_ := gtk.LabelNew(title)
	box.Add(label)

	btn.Add(box)
	btn.ShowAll()
	//btn.SetOpacity(0.5)
	return btn
}



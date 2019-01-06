package component

import (
	"github.com/gotk3/gotk3/gtk"
)


/**
	创建一个带图标文字带按钮
 */
func GenerateNavItem(width int,height int,imagePath string,title string) *gtk.Button{
	button,_ := gtk.BuilderNew()

	//vBox := gtk.NewVBox(false,0)
	//vBox.SetSizeRequest(width,height)
	//
	////image := helper.CreateImage(width-5,width-5,imagePath)
	////image.SetSizeRequest(width,width)
	//
	//label := gtk.NewLabel(title)
	//label.SetSizeRequest(width,height-width)
	//label.Show()
	//
	////vBox.Add(image)
	//vBox.Add(label)
	//vBox.Show()
	//
	//button.Add(vBox)
	//button.SetBorderWidth(0)
	//button.SetVisible(true)

	return button
}
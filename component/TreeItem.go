package component

//import (
//	"github.com/mattn/go-gtk/gtk"
//	)
//
//type TreeData struct{
//	TreeUnit TreeUnit
//	SubTreeData *TreeData
//}
//
//type TreeUnit struct{
//	ImagePath string
//	Text string
//}
//
//func CreateTreeItem(width int,height int,imagePath string, text string) *gtk.Button{
//
//	button := gtk.NewButton()
//	button.SetSizeRequest(width,height)
//
//	hBox := gtk.NewHBox(false,5)
//
//	//image := helper.CreateImage(height,height,imagePath)
//	label := gtk.NewLabel(text)
//	//hBox.Add(image)
//	hBox.Add(label)
//
//	button.Add(hBox)
//
//	return button
//}
//
//func CreateTree(treeData TreeData) *gtk.VBox{
//	vBox := gtk.NewVBox(false,0)
//
//	return vBox
//}
package component

import (
	"github.com/mattn/go-gtk/gtk"
)

func CreateFormItem(width int, height int,text string) *gtk.HBox{
	hBox := gtk.NewHBox(false,0)
	hBox.SetVisible(true)
	hBox.SetUSize(width,height)

	label := gtk.NewLabel(text)
	label.SetUSize(width/2,height)

	entry := gtk.NewEntry()
	entry.SetText(text)
	entry.SetUSize(width/2,height)
	
	hBox.Add(label)
	hBox.Add(entry)

	return hBox
}

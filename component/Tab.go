package component

import (
	"github.com/gotk3/gotk3/gtk"
	"../helper"
	"../common"
	"fmt"
	"github.com/gotk3/gotk3/gdk"
	common2 "github.com/itgeniusshuai/go_common/common"
)

func CreateTabItem(imagePath string,text string) (*gtk.Box,*gtk.EventBox,*gtk.EventBox){
	box,_ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,0)

	eventBox,_ := gtk.EventBoxNew()
	subBox,_ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,0)

	box.Add(helper.CreateImage(5,5,imagePath))

	label,_ := gtk.LabelNew(text)
	label.SetSizeRequest(70,5)

	subBox.Add(label)

	eventBox.Add(subBox)

	box.Add(eventBox)

	imageEventBox,_ := gtk.EventBoxNew()
	image := helper.CreateImage(5,5,imagePath)
	imageEventBox.Add(image)
	box.Add(imageEventBox)


	return box,eventBox,imageEventBox
}

type Tab struct{
	tabBoxMap map[string]*gtk.Box
	tabBoxs []string
	curBox *gtk.Box
}


func (tab *Tab)ToTabBox() *gtk.Box{
	box,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)

	// tab条
	adjust,_ := gtk.AdjustmentNew(common.LeftScrollInital,common.LeftScrollLower,common.LeftScrollUpper,common.LeftScrollStepIncrement,common.LeftScrollPageIncrement,common.LeftScrollPageSize)

	scroll,_ := gtk.ScrolledWindowNew(adjust,nil)

	var curBox,_ = gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)
	tab.flushBar(scroll,curBox)
	// 展示区域
	box.Add(scroll)
	box.Add(curBox)

	box.ShowAll()
	return box
}

var i = 0
var colors = []string{"00f","f00","#0f0"}


func (tab *Tab)flushBar(scroll *gtk.ScrolledWindow,curBox *gtk.Box){
	box,_ :=  gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,0)
	keys := Keys(tab.tabBoxMap)
	//var lastKey string
	curBox.Add(tab.curBox)
	for _,k:= range keys{
		tab.bindEvent(box,curBox,k)
	}
	scroll.Add(box)
}

func (tab *Tab)bindEvent(box *gtk.Box,curBox *gtk.Box,k string){
	bar,eventBox,imageBox := CreateTabItem(common.BarImagePath,k)

	box.Add(bar)
	var k1 = k
	var bar1 = bar
	var box1 = box
	eventBox.Connect("button_press_event", func(box *gtk.EventBox) {
		fmt.Println(k1)
		curBox.GetChildren().Foreach(func(item interface{}) {
			ie := item.(*gtk.Widget)
			curBox.Remove(ie)
		})
		helper.ChangeBgColor(k,tab.tabBoxMap[k1],colors[i%3])
		i++
		curBox.Add(tab.tabBoxMap[k1])
		curBox.ShowAll()
	})
	imageBox.AddEvents(int(gdk.EVENT_BUTTON_PRESS))
	imageBox.Connect("button_press_event", func(box *gtk.EventBox) {
		tab.tabBoxMap[k1].Destroy()
		delete(tab.tabBoxMap, k)
		box1.Remove(bar1)

		curBox.GetChildren().Foreach(func(item interface{}) {
			ie := item.(*gtk.Widget)
			curBox.Remove(ie)
		})
		var keyIndex = common2.IndexOfStrArr(tab.tabBoxs,k1)
		tab.tabBoxs = common2.RemoveStrArr(tab.tabBoxs,keyIndex)
		tabBoxLength := len(tab.tabBoxs)
		if tabBoxLength > 0 {
			if keyIndex >= tabBoxLength{
				keyIndex = tabBoxLength-1
			}
			newKey := tab.tabBoxs[keyIndex]

			curBox.Add(tab.tabBoxMap[newKey])
		}

		curBox.ShowAll()
		box1.ShowAll()
	})
}

func (tab *Tab)AddTab(text string, box *gtk.Box){
	if tab.tabBoxMap == nil{
		tab.tabBoxMap = make(map[string]*gtk.Box,0)
		tab.curBox = box
		fmt.Println("nil")
	}
	tab.tabBoxMap[text] = box
	if tab.tabBoxs == nil{
		tab.tabBoxs = make([]string,0)
	}
	tab.tabBoxs = append(tab.tabBoxs, text)
}

func Keys(kv map[string]*gtk.Box)[]string{
	var keys = []string{}
	for k,_ := range kv{
		keys = append(keys, k)
	}
	return keys
}
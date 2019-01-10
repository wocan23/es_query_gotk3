package component

import (
	"github.com/gotk3/gotk3/gtk"
	"../helper"
	"../common"
	"fmt"
	"github.com/gotk3/gotk3/gdk"
	common2 "github.com/itgeniusshuai/go_common/common"
	"sync"
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
	barBox *gtk.Box
}

func CreateTabBox()(*Tab,*gtk.Box){
	tab := new(Tab)
	tabBox := tab.toTabBox()
	return tab,tabBox
}


func (tab *Tab)toTabBox() *gtk.Box{
	box,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)

	// tab条
	adjust,_ := gtk.AdjustmentNew(common.LeftScrollInital,common.LeftScrollLower,common.LeftScrollUpper,common.LeftScrollStepIncrement,common.LeftScrollPageIncrement,common.LeftScrollPageSize)

	scroll,_ := gtk.ScrolledWindowNew(adjust,nil)

	var curBox,_ = gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)
	tab.curBox = curBox
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
	barBox,_ :=  gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,0)
	keys := Keys(tab.tabBoxMap)
	for _,k:= range keys{
		tab.AddTabAndBind(barBox,curBox,k)
	}
	scroll.Add(barBox)
	tab.barBox = barBox
}



func (tab *Tab)AddTabAndBind(barBox *gtk.Box,curBox *gtk.Box,k string){
	bar,eventBox,imageBox := CreateTabItem(common.BarImagePath,k)

	barBox.Add(bar)
	eventBox.Connect("button_press_event", func(box *gtk.EventBox) {
		helper.RemoveAndAddNew(curBox,tab.tabBoxMap[k])
	})
	imageBox.AddEvents(int(gdk.EVENT_BUTTON_PRESS))
	imageBox.Connect("button_press_event", func(box *gtk.EventBox) {
		tab.tabBoxMap[k].Destroy()
		delete(tab.tabBoxMap, k)
		barBox.Remove(bar)


		var newCurBox *gtk.Box
		var keyIndex = common2.IndexOfStrArr(tab.tabBoxs,k)
		tab.tabBoxs = common2.RemoveStrArr(tab.tabBoxs,keyIndex)
		tabBoxLength := len(tab.tabBoxs)
		if tabBoxLength > 0 {
			if keyIndex >= tabBoxLength{
				keyIndex = tabBoxLength-1
			}
			newKey := tab.tabBoxs[keyIndex]

			newCurBox  = tab.tabBoxMap[newKey]
		}
		helper.RemoveAndAddNew(curBox,newCurBox)

		barBox.ShowAll()
		curBox.ShowAll()
	})
	helper.RemoveAndAddNew(curBox,tab.tabBoxMap[k])
	barBox.ShowAll()
	curBox.ShowAll()
}

var lock sync.Mutex

func AddTab(tab *Tab,text string, box *gtk.Box){
	defer lock.Unlock()
	lock.Lock()
	if tab.tabBoxMap == nil{
		tab.tabBoxMap = make(map[string]*gtk.Box,0)
		fmt.Println("nil")
	}
	tab.tabBoxMap[text] = box
	if tab.tabBoxs == nil{
		tab.tabBoxs = make([]string,0)
	}
	tab.tabBoxs = append(tab.tabBoxs, text)
	tab.AddTabAndBind(tab.barBox,tab.curBox,text)
}

func Keys(kv map[string]*gtk.Box)[]string{
	var keys = []string{}
	for k,_ := range kv{
		keys = append(keys, k)
	}
	return keys
}
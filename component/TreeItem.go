package component

import (
	"github.com/gotk3/gotk3/gtk"
	"log"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gdk"
	"../common"
)

const (
	COLUMN_ICON = iota
	COLUMN_TEXT
)

type TreeData struct{
	TreeUnit TreeUnit
	SubTreeData *TreeData
}

type TreeUnit struct{
	ImagePath string
	Text string
}

func CreateTree() *gtk.Box{
	vBox,_ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL,0)
	//gtk.ScrollbarNew(gtk.ORIENTATION_VERTICAL,nil)
	treeView, err := gtk.TreeViewNew()
	if err != nil {
		log.Fatal("Unable to create tree view:", err)
	}

	treeView.AppendColumn(createItemColumn("", COLUMN_ICON))
	treeView.SetHExpand(false)

	treeStore, err := gtk.TreeStoreNew(glib.TYPE_OBJECT, glib.TYPE_STRING)
	if err != nil {
		log.Fatal("Unable to create tree store:", err)
	}
	treeView.SetModel(treeStore)

	indexspixbuf,_ := gdk.PixbufNewFromFile("/Users/zhaoshuai/Documents/go_workspace_wocan/es_query_gotk3/images/index.png")
	indexpixbuf,_ := indexspixbuf.ScaleSimple(20,20,gdk.INTERP_TILES)

	conspixbuf,_ := gdk.PixbufNewFromFile("/Users/zhaoshuai/Documents/go_workspace_wocan/es_query_gotk3/images/conn.png")
	conpixbuf,_ := conspixbuf.ScaleSimple(20,20,gdk.INTERP_TILES)

	docspixbuf,_ := gdk.PixbufNewFromFile("/Users/zhaoshuai/Documents/go_workspace_wocan/es_query_gotk3/images/doc.png")
	docpixbuf,_ := docspixbuf.ScaleSimple(20,20,gdk.INTERP_TILES)

	// todo data implements after
	iter1 := addRow(treeStore, conpixbuf, "127.0.0.1                            ")
	iter2 := addSubRow(treeStore, iter1, indexpixbuf, "test1-1-------------------------------------1")
	iter2 = addSubRow(treeStore, iter1, indexpixbuf, "test1-2")
	addSubRow(treeStore, iter2, docpixbuf, "test1-2-1")
	addSubRow(treeStore, iter2, docpixbuf, "test1-2-2")
	addSubRow(treeStore, iter2, docpixbuf, "test1-2-3")
	iter2 = addSubRow(treeStore, iter1, indexpixbuf, "test1-3")
	iter1 = addRow(treeStore, conpixbuf, "10.70.93.52                           ")
	iter2 = addSubRow(treeStore, iter1, indexpixbuf, "test2-1")
	iter2 = addSubRow(treeStore, iter1, indexpixbuf, "test2-2")
	iter2 = addSubRow(treeStore, iter1, indexpixbuf, "test2-3")
	addSubRow(treeStore, iter2, docpixbuf, "test2-3-1")
	addSubRow(treeStore, iter2, docpixbuf, "test2-3-2")

	// 处理选择
	selection, err := treeView.GetSelection()
	if err != nil {
		log.Fatal("Could not get tree selection object.")
	}
	selection.SetMode(gtk.SELECTION_BROWSE)
	selection.Connect("changed", treeSelectionChangedCB)

	// treeview
	treeView.SetSizeRequest(common.WindowLeftWidth,common.WindowLeftHeight)
	treeView.SetMarginTop(5)
	treeView.SetEnableSearch(true)
	treeView.SetSearchColumn(0)
	treeView.SetHAlign(gtk.ALIGN_START)

	searchEntity,_ := gtk.EntryNew()
	searchEntity.SetSizeRequest(common.WindowLeftWidth,300)
	searchEntity.SetText("dfsdf")
	searchEntity.SetEditable(true)
	treeView.SetSearchEntry(searchEntity)
	treeView.ShowAll()

	vBox.Add(treeView)
	vBox.SetSizeRequest(common.WindowLeftWidth,common.WindowLeftHeight)
	return vBox
}

func createItemColumn(title string, id int) *gtk.TreeViewColumn {
	cellpb,_ := gtk.CellRendererPixbufNew()
	col,_ := gtk.TreeViewColumnNewWithAttribute(title, cellpb, "pixbuf",id)
	cell,_ := gtk.CellRendererTextNew()
	col.PackEnd(cell, false)
	col.AddAttribute(cell, "text", 1)
	col.SetFixedWidth(common.WindowLeftWidth)
	col.SetClickable(true)

	//cellpb,_ := gtk.CellRendererPixbufNew()
	//cell,_ := gtk.CellRendererTextNew()
	//col,_ := gtk.TreeViewColumnNewWithAttribute(title, cell, "text",1)
	//col.PackEnd(cellpb, true)
	//col.AddAttribute(cellpb, "pixbuf", id)
	return col
}

// Append a sub row to the tree store for the tree view
func addSubRow(treeStore *gtk.TreeStore, iter *gtk.TreeIter, icon *gdk.Pixbuf, text string) *gtk.TreeIter {
	// Get an iterator for a new row at the end of the list store
	i := treeStore.Append(iter)

	// Set the contents of the tree store row that the iterator represents
	err := treeStore.SetValue(i, COLUMN_ICON, icon)
	if err != nil {
		log.Fatal("Unable set value:", err)
	}
	err = treeStore.SetValue(i, COLUMN_TEXT, text)
	if err != nil {
		log.Fatal("Unable set value:", err)
	}
	return i
}

// Append a toplevel row to the tree store for the tree view
func addRow(treeStore *gtk.TreeStore, icon *gdk.Pixbuf, text string) *gtk.TreeIter {
	return addSubRow(treeStore, nil, icon, text)
}

// Handle selection
func treeSelectionChangedCB(selection *gtk.TreeSelection) {
	var iter *gtk.TreeIter
	var model gtk.ITreeModel
	var ok bool
	model, iter, ok = selection.GetSelected()
	if ok {
		tpath, err := model.(*gtk.TreeModel).GetPath(iter)
		if err != nil {
			log.Printf("treeSelectionChangedCB: Could not get path from model: %s\n", err)
			return
		}
		log.Printf("treeSelectionChangedCB: selected path: %s\n", tpath)
	}
}
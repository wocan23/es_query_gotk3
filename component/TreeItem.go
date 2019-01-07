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
	treeView, err := gtk.TreeViewNew()
	if err != nil {
		log.Fatal("Unable to create tree view:", err)
	}

	treeView.AppendColumn(createImageColumn("", COLUMN_ICON))
	treeView.AppendColumn(createTextColumn("", COLUMN_TEXT))

	// Creating a tree store. This is what holds the data that will be shown on our tree view.
	treeStore, err := gtk.TreeStoreNew(glib.TYPE_OBJECT, glib.TYPE_STRING)
	if err != nil {
		log.Fatal("Unable to create tree store:", err)
	}
	treeView.SetModel(treeStore)

	spixbuf,_ := gdk.PixbufNewFromFile("/Users/zhaoshuai/Documents/go_workspace_wocan/es_query_gotk3/images/index.png")
	pixbuf,_ := spixbuf.ScaleSimple(30,30,gdk.INTERP_HYPER)

	// todo data implements after
	iter1 := addRow(treeStore, pixbuf, "Testsuite 1")
	iter2 := addSubRow(treeStore, iter1, pixbuf, "test1-1")
	iter2 = addSubRow(treeStore, iter1, pixbuf, "test1-2")
	addSubRow(treeStore, iter2, pixbuf, "test1-2-1")
	addSubRow(treeStore, iter2, pixbuf, "test1-2-2")
	addSubRow(treeStore, iter2, pixbuf, "test1-2-3")
	iter2 = addSubRow(treeStore, iter1, pixbuf, "test1-3")
	iter1 = addRow(treeStore, pixbuf, "Testsuite 2")
	iter2 = addSubRow(treeStore, iter1, pixbuf, "test2-1")
	iter2 = addSubRow(treeStore, iter1, pixbuf, "test2-2")
	iter2 = addSubRow(treeStore, iter1, pixbuf, "test2-3")
	addSubRow(treeStore, iter2, pixbuf, "test2-3-1")
	addSubRow(treeStore, iter2, pixbuf, "test2-3-2")

	// 处理选择
	selection, err := treeView.GetSelection()
	if err != nil {
		log.Fatal("Could not get tree selection object.")
	}
	selection.SetMode(gtk.SELECTION_SINGLE)
	selection.Connect("changed", treeSelectionChangedCB)

	vBox.Add(treeView)
	vBox.SetSizeRequest(common.WindowLeftWidth,common.WindowLeftHeight)
	vBox.ShowAll()
	return vBox
}

// Add a column to the tree view (during the initialization of the tree view)
// We need to distinct the type of data shown in either column.
func createTextColumn(title string, id int) *gtk.TreeViewColumn {
	// In this column we want to show text, hence create a text renderer
	cellRenderer, err := gtk.CellRendererTextNew()
	if err != nil {
		log.Fatal("Unable to create text cell renderer:", err)
	}

	// Tell the renderer where to pick input from. Text renderer understands
	// the "text" property.
	column, err := gtk.TreeViewColumnNewWithAttribute(title, cellRenderer, "text", id)
	if err != nil {
		log.Fatal("Unable to create cell column:", err)
	}

	return column
}

func createImageColumn(title string, id int) *gtk.TreeViewColumn {
	// In this column we want to show image data from Pixbuf, hence
	// create a pixbuf renderer
	cellRenderer, err := gtk.CellRendererPixbufNew()
	if err != nil {
		log.Fatal("Unable to create pixbuf cell renderer:", err)
	}

	// Tell the renderer where to pick input from. Pixbuf renderer understands
	// the "pixbuf" property.
	column, err := gtk.TreeViewColumnNewWithAttribute(title, cellRenderer, "pixbuf", id)
	if err != nil {
		log.Fatal("Unable to create cell column:", err)
	}

	return column
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
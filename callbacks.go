package main

import (
	"github.com/PucklaMotzer09/gohomeengine/src/frameworks/GTK/gtk"
	"github.com/PucklaMotzer09/gohomeengine/src/gohome"
	"io/ioutil"
	"log"
	"os"
)

func updateResolution(widget gtk.Widget) {
	w, h := widget.GetSize()
	gohome.Render.SetNativeResolution(uint32(w), uint32(h))
}

func quitApplication(menuItem gtk.MenuItem) {
	gohome.MainLop.Quit()
}

func onToolPlace(toolButton gtk.ToolButton) {
}

func onToolMove(toolButton gtk.ToolButton) {
	lb := gtk.LabelNew("ToolMove")
	lb_assets.Insert(lb.ToWidget(), -1)
	lb.ToWidget().Show()
}

func onToolRotate(toolButton gtk.ToolButton) {
	lb := gtk.LabelNew("ToolRotate")
	lb_assets.Insert(lb.ToWidget(), -1)
	lb.ToWidget().Show()
}

func onToolScale(toolButton gtk.ToolButton) {
	lb := gtk.LabelNew("ToolPlace")
	lb_assets.Insert(lb.ToWidget(), -1)
	lb.ToWidget().Show()
}

func onToolLoadModel(toolButton gtk.ToolButton) {
	fileChooser := gtk.FileChooserDialogNew("Load Model", gtk.GetWindow(), gtk.FILE_CHOOSER_ACTION_OPEN)
	if fileChooser.ToDialog().Run() == gtk.RESPONSE_ACCEPT {
		filename := fileChooser.ToFileChooser().GetFilename()
		file, _ := os.Open(filename)
		contents, _ := ioutil.ReadAll(file)
		name := gohome.GetFileFromPath(filename)

		var lm LoadableModel
		lm.Name = name
		lm.FileContents = string(contents)
		lm.Filename = filename
		loadable_models = append(loadable_models, lm)
	} else {
		gohome.ErrorMgr.Error("Load", "Model", "An error acoured")
	}

	fileChooser.ToWidget().Destroy()
}

func onSelectAsset(listBox gtk.ListBox, listBoxRow gtk.ListBoxRow) {
	lbl := listBoxRow.ToContainer().GetChildren().Data().ToLabel()
	data := lbl.ToGObject().GetData("ID")
	id := stringToUint32(data)
	selected_model = id
}

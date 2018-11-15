package main

import (
	"fmt"
	"github.com/PucklaMotzer09/gohomeengine/src/frameworks/GTK/gtk"
	"github.com/PucklaMotzer09/gohomeengine/src/gohome"
	"io/ioutil"
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
	lb := gtk.LabelNew("ToolPlace")
	lb_assets.Insert(lb.ToWidget(), -1)
	lb.ToWidget().Show()
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
		fileChooser.ToWidget().Destroy()
		file, _ := os.Open(filename)
		contents, _ := ioutil.ReadAll(file)
		fmt.Println(string(contents))
		/*gohome.ErrorMgr.Log("Load", "Model", filename)
		name := gohome.GetFileFromPath(filename)
		gohome.ResourceMgr.LoadLevel(name, filename, true)

		level := gohome.ResourceMgr.GetLevel(name)
		if level != nil && len(level.LevelObjects) != 0 {
			model := level.LevelObjects[0].Model3D
			if model != nil {
				loaded_models = append(loaded_models, model)
				lbl := gtk.LabelNew(name)
				lb_assets.Insert(lbl.ToWidget(), -1)
				lbl.ToWidget().Show()
			}
		} else {
			gohome.ErrorMgr.Error("Load", "Model", "Loaded model is broken")
		}*/
	} /* else {
		gohome.ErrorMgr.Error("Load", "Model", "An error acoured")
	}*/
}

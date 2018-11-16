package main

import (
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
	current_mode = MODE_PLACE
}

func onToolMove(toolButton gtk.ToolButton) {
	current_mode = MODE_MOVE
}

func onToolRotate(toolButton gtk.ToolButton) {
	current_mode = MODE_ROTATE
}

func onToolScale(toolButton gtk.ToolButton) {
	current_mode = MODE_SCALE
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
		lm.Life = 1
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

func onLeftClick() {
	if current_mode == MODE_PLACE {
		pmodel, ok := placable_models[selected_model]
		if ok {
			if placed_models == nil {
				placed_models = make(map[uint32]*PlacedModel)
			}

			model := loaded_models[pmodel.ID]
			var pm PlacedModel
			pm.Entity3D.InitModel(model)
			pm.PlacedObject.Transform = pm.Entity3D.Transform
			pm.PlacedObject.AABB = model.AABB
			pm.PlacedObject.PlaceID = place_id
			pm.PlaceableModel = pmodel
			placed_models[place_id] = &pm

			place_id++

			gohome.RenderMgr.AddObject(&pm.Entity3D)
			pm.PlacedObject.Transform.Position = camera.Position.Add(camera.LookDirection.Mul(3.0))
		}
	}
}

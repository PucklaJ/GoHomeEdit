package main

import (
	"github.com/PucklaMotzer09/gohomeengine/src/frameworks/GTK/gtk"
	"github.com/PucklaMotzer09/gohomeengine/src/gohome"
	"io/ioutil"
)

func updateResolution(widget gtk.Widget) {
	w, h := widget.GetSize()
	gohome.Render.SetNativeResolution(uint32(w), uint32(h))
	pickable_texture.ChangeSize(uint32(w), uint32(h))
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

	filter := gtk.FileFilterNew()
	filter.AddPattern("*.obj")

	fileChooser.ToFileChooser().SetFilter(filter)
	fileChooser.ToFileChooser().SetSelectMultiple(true)

	if fileChooser.ToDialog().Run() == gtk.RESPONSE_ACCEPT {
		filenames := fileChooser.ToFileChooser().GetFilenames()
		for i := 0; i < len(filenames); i++ {
			filename := filenames[i]
			file, _ := gohome.Framew.OpenFile(filename)
			contents, _ := ioutil.ReadAll(file)
			file.Close()
			name := gohome.GetFileFromPath(filename)

			var lm LoadableModel
			lm.Name = name
			lm.FileContents = string(contents)
			lm.Filename = filename
			loadable_models = append(loadable_models, lm)
		}
	} else {
		gohome.ErrorMgr.Error("Load", "Model", "An error acoured")
	}
	fileChooser.ToWidget().Destroy()
}

var is_wireframe = false

func onMenuWireframe(menuItem gtk.MenuItem) {
	is_wireframe = !is_wireframe
	gohome.RenderMgr.WireFrameMode = is_wireframe
}

func onMenuPlaceOnGrid(menuItem gtk.MenuItem) {
	place_on_grid = !place_on_grid
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
			pm.Entity3D.SetType(gohome.TYPE_3D_NORMAL | PICKABLE_BIT)
			pm.PlacedObject.Transform = pm.Entity3D.Transform
			pm.PlacedObject.AABB = model.AABB
			pm.PlacedObject.PlaceID = place_id
			pm.PlaceableModel = pmodel
			placed_models[place_id] = &pm

			place_id++

			gohome.RenderMgr.AddObject(&pm.Entity3D)
			pm.PlacedObject.Transform.Position = placing_object.Transform.Position
			selected_placed_object = &pm.PlacedObject
		}
	} else if current_mode == MODE_MOVE {
		if !handleMoveArrowClick() {
			handlePickableClick()
		}
	} else if current_mode == MODE_SCALE {
		if !handleScaleArrowClick() {
			handlePickableClick()
		}
	}
}

func onLeftClickRelease() {
	is_transforming = false
	arrows.TransformAxis = 0
	arrows.ResetPosition()
	arrows.SetScale()
}

package main

import (
	"github.com/PucklaMotzer09/gohomeengine/src/frameworks/GTK/gtk"
	"github.com/PucklaMotzer09/gohomeengine/src/gohome"
	"github.com/PucklaMotzer09/mathgl/mgl32"
	"io/ioutil"
)

func updateResolution(widget gtk.Widget) {
	w, h := widget.GetSize()
	gohome.Render.SetNativeResolution(uint32(w), uint32(h))
	pickable_texture.ChangeSize(uint32(w), uint32(h))
	gohome.RenderMgr.ReRender = true
}

func quitApplication(menuItem gtk.MenuItem) {
	gohome.MainLop.Quit()
}

func onToolPlace(toolButton gtk.ToolButton) {
	current_mode = MODE_PLACE
	gohome.RenderMgr.ReRender = true
}

func onToolMove(toolButton gtk.ToolButton) {
	current_mode = MODE_MOVE
	gohome.RenderMgr.ReRender = true
}

func onToolRotate(toolButton gtk.ToolButton) {
	current_mode = MODE_ROTATE
	gohome.RenderMgr.ReRender = true
}

func onToolScale(toolButton gtk.ToolButton) {
	current_mode = MODE_SCALE
	gohome.RenderMgr.ReRender = true
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
	gohome.RenderMgr.ReRender = true
}

func onMenuPlaceOnGrid(menuItem gtk.MenuItem) {
	place_on_grid = !place_on_grid
	gohome.RenderMgr.ReRender = true
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
			if instanced_entities == nil {
				instanced_entities = make(map[PlaceableObject]*gohome.InstancedEntity3D)
			}
			if pickable_colors == nil {
				pickable_colors = make(map[*gohome.InstancedEntity3D][]mgl32.Vec4)
			}

			model := loaded_models[pmodel.ID]
			entity, entok := instanced_entities[pmodel.PlaceableObject]
			var pm PlacedModel
			if !entok {
				imodel := gohome.InstancedModel3DFromModel3D(model)
				entity = &gohome.InstancedEntity3D{}
				imodel.AddValue(gohome.VALUE_VEC4)
				imodel.SetName(0, gohome.VALUE_VEC4, "pickableColor")
				entity.InitModel(imodel, 10)
				entity.SetType(gohome.TYPE_3D_INSTANCED | PICKABLE_BIT)
				entity.SetNumUsedInstances(0)
				instanced_entities[pmodel.PlaceableObject] = entity
				gohome.RenderMgr.AddObject(entity)
				pickable_colors[entity] = make([]mgl32.Vec4, 10)
			}
			index := entity.Model3D.GetNumUsedInstances()
			entity.SetNumUsedInstances(entity.Model3D.GetNumUsedInstances() + 1)
			if entity.Model3D.GetNumUsedInstances() > entity.Model3D.GetNumInstances() {
				entity.SetNumInstances(entity.Model3D.GetNumInstances() + 10)
				entity.SetNumUsedInstances(entity.Model3D.GetNumInstances() - 9)
				pickable_colors[entity] = append(pickable_colors[entity], make([]mgl32.Vec4, 10)...)
			}
			pickable_colors[entity][index] = idToColor(place_id)
			pm.PlacedObject.Transform = &entity.Transforms[index].TransformableObject3D
			pm.PlacedObject.AABB = entity.Model3D.AABB
			pm.PlacedObject.PlaceID = place_id
			pm.PlaceableModel = pmodel
			placed_models[place_id] = &pm

			place_id++

			pm.PlacedObject.Transform.Position = placing_object.Transform.Position
			selected_placed_object = &pm.PlacedObject

			gohome.RenderMgr.ReRender = true
		}
	} else if current_mode == MODE_MOVE {
		if !handleMoveArrowClick() {
			handlePickableClick()
		} else {
			gohome.RenderMgr.ReRender = true
		}
	} else if current_mode == MODE_SCALE {
		if !handleScaleArrowClick() {
			handlePickableClick()
		} else {
			gohome.RenderMgr.ReRender = true
		}
	}
}

func onLeftClickRelease() {
	if is_transforming {
		is_transforming = false
		arrows.TransformAxis = 0
		arrows.ResetPosition()
		arrows.SetScale()
	}
	gohome.RenderMgr.ReRender = true
}

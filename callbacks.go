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
	current_mode = MODE_PLACE
	gohome.RenderMgr.ReRender = true
}

func onLeftClick() {
	if current_mode == MODE_PLACE {
		handlePlaceClick()
	} else if current_mode == MODE_MOVE {
		handleMoveClick()
	} else if current_mode == MODE_SCALE {
		handleScaleClick()
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

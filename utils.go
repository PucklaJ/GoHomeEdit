package main

import (
	"github.com/PucklaMotzer09/gohomeengine/src/frameworks/GTK/gtk"
	"github.com/PucklaMotzer09/gohomeengine/src/gohome"
)

func loadModel(name string, fileContents string) {
	gohome.ErrorMgr.Log("Load", "Model", name)
	gohome.ResourceMgr.LoadLevelString(name, string(fileContents), "obj", true)

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
	}
}

func loadLoadableModels() {
	for i := 0; i < len(loadable_models); i++ {
		loadModel(loadable_models[i].Name, loadable_models[i].FileContents)
	}
	loadable_models = loadable_models[:0]
}

package main

import (
	"encoding/binary"
	"github.com/PucklaMotzer09/gohomeengine/src/frameworks/GTK/gtk"
	"github.com/PucklaMotzer09/gohomeengine/src/gohome"
)

func uint32ToString(i uint32) string {
	var buffer [4]byte
	n := binary.PutUvarint(buffer[:], uint64(i))
	return string(buffer[:n])
}

func stringToUint32(str string) uint32 {
	i, _ := binary.Uvarint([]byte(str))
	return uint32(i)
}

func loadModel(name, fileContents, fileName string) {
	gohome.ErrorMgr.Log("Load", "Model", name)
	gohome.ResourceMgr.LoadLevelString(name, string(fileContents), "obj", true)

	level := gohome.ResourceMgr.GetLevel(name)
	if level != nil && len(level.LevelObjects) != 0 {
		model := level.LevelObjects[0].Model3D
		if model != nil {
			if loaded_models == nil {
				loaded_models = make(map[uint32]*gohome.Model3D)
			}
			if placable_models == nil {
				placable_models = make(map[uint32]*PlaceableModel)
			}
			loaded_models[object_id] = model

			lbl := gtk.LabelNew(name)
			lbl.ToGObject().SetData("ID", uint32ToString(object_id))
			lb_assets.Insert(lbl.ToWidget(), -1)
			lbl.ToWidget().Show()

			var pm PlaceableModel
			pm.Name = name
			pm.Filename = fileName
			pm.ID = object_id
			placable_models[object_id] = &pm
			object_id++
		}
	} else {
		gohome.ErrorMgr.Error("Load", "Model", "Loaded model is broken")
	}
}

func loadLoadableModels() {
	for i := 0; i < len(loadable_models); i++ {
		loadModel(loadable_models[i].Name, loadable_models[i].FileContents, loadable_models[i].Filename)
	}
	loadable_models = loadable_models[:0]
}

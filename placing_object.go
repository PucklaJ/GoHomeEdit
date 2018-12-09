package main

import (
	"github.com/PucklaMotzer09/gohomeengine/src/gohome"
)

type PlacingObject struct {
	gohome.Entity3D
	prev_selected_model uint32
}

func (this *PlacingObject) Init() {
	this.Visible = false
	this.RenderLast = true
	gohome.RenderMgr.AddObject(this)
	gohome.UpdateMgr.AddObject(this)

	this.prev_selected_model = 1000000
}

func (this *PlacingObject) Update(delta_time float32) {
	if this.prev_selected_model != selected_model || this.Model3D == nil {
		pmodel, ok := placable_models[selected_model]
		if ok {
			model := loaded_models[pmodel.ID]
			placemodel := model.Copy()
			mesh := placemodel.GetMeshIndex(0)
			for i := 0; mesh != nil; i++ {
				mesh.GetMaterial().Transparency = 0.5
				mesh = placemodel.GetMeshIndex(uint32(i + 1))
			}
			this.InitModel(placemodel)
			this.RenderLast = true
		}
		this.prev_selected_model = selected_model
	}

	this.Visible = current_mode == MODE_PLACE && this.Model3D != nil
}

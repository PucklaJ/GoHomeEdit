package main

import (
	"github.com/PucklaMotzer09/gohomeengine/src/gohome"
)

type PlacingObject struct {
	gohome.Entity3D
	prev_selected_model uint32
}

func (this *PlacingObject) Init() {
	gohome.ResourceMgr.LoadShaderSource("PlacingObject", gohome.ENTITY_3D_SHADER_VERTEX_SOURCE_OPENGL, PLACING_OBJECT_SHADER_FRAGMENT_SOURCE_OPENGL, "", "", "", "")
	if gohome.ResourceMgr.GetShader("PlacingObject") == nil {
		gohome.ResourceMgr.LoadShaderSource("PlacingObjectNoShadow", gohome.ENTITY_3D_NO_SHADOWS_SHADER_VERTEX_SOURCE_OPENGL, PLACING_OBJECT_NO_SHADOWS_SHADER_FRAGMENT_SOURCE_OPENGL, "", "", "", "")
		if gohome.ResourceMgr.GetShader("PlacingObjectNoShadow") != nil {
			gohome.ResourceMgr.SetShader("PlacingObject", "PlacingObjectNoShadow")
		}
	}
	gohome.ResourceMgr.LoadShaderSource("PlacingObjectNoUV", gohome.ENTITY_3D_NOUV_SHADER_VERTEX_SOURCE_OPENGL, PLACING_MODEL_NOUV_SHADER_FRAGMENT_SOURCE_OPENGL, "", "", "", "")
	if gohome.ResourceMgr.GetShader("PlacingObjectNoUV") == nil {
		gohome.ResourceMgr.LoadShaderSource("PlacingObjectNoUVNoShadow", gohome.ENTITY_3D_NOUV_SHADER_VERTEX_SOURCE_OPENGL, PLACING_OBJECT_NOUV_NO_SHADOWS_SHADER_FRAGMENT_SOURCE_OPENGL, "", "", "", "")
		if gohome.ResourceMgr.GetShader("PlacingObjectNoUVNoShadow") != nil {
			gohome.ResourceMgr.SetShader("PlacingObjectNoUV", "PlacingObjectNoUVNoShadow")
		}
	}

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
			this.InitModel(model)
			this.RenderLast = true
			if model.HasUV() {
				this.SetShader(gohome.ResourceMgr.GetShader("PlacingObject"))
			} else {
				this.SetShader(gohome.ResourceMgr.GetShader("PlacingObjectNoUV"))
			}
		}
		this.prev_selected_model = selected_model
	}

	this.Visible = current_mode == MODE_PLACE && this.Model3D != nil
	if this.Visible {
		this.Transform.Position = camera_center
	}
}

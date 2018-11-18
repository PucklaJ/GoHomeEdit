package main

import (
	"github.com/PucklaMotzer09/gohomeengine/src/gohome"
	"github.com/PucklaMotzer09/mathgl/mgl32"
	"golang.org/x/image/colornames"
)

const (
	ARROWS_SIZE float32 = 0.1
)

type Arrows struct {
	translateX gohome.Entity3D
	translateY gohome.Entity3D
	translateZ gohome.Entity3D

	scaleX gohome.Entity3D
	scaleY gohome.Entity3D
	scaleZ gohome.Entity3D

	rotateX gohome.Entity3D
	rotateY gohome.Entity3D
	rotateZ gohome.Entity3D
}

func (this *Arrows) Init() {
	gohome.ResourceMgr.LoadLevel("Arrows", "arrows.obj", true)

	this.translateX.InitModel(gohome.ResourceMgr.GetLevel("Arrows").GetModel("Arrow_Cone"))
	this.translateY.InitModel(gohome.ResourceMgr.GetLevel("Arrows").GetModel("Arrow_Cone").Copy())
	this.translateZ.InitModel(gohome.ResourceMgr.GetLevel("Arrows").GetModel("Arrow_Cone").Copy())

	this.scaleX.InitModel(gohome.ResourceMgr.GetLevel("Arrows").GetModel("Block_Cube"))
	this.scaleY.InitModel(gohome.ResourceMgr.GetLevel("Arrows").GetModel("Block_Cube").Copy())
	this.scaleZ.InitModel(gohome.ResourceMgr.GetLevel("Arrows").GetModel("Block_Cube").Copy())

	this.translateX.Model3D.GetMeshIndex(0).GetMaterial().DiffuseColor = colornames.Red
	this.translateY.Model3D.GetMeshIndex(0).GetMaterial().DiffuseColor = colornames.Lime
	this.translateZ.Model3D.GetMeshIndex(0).GetMaterial().DiffuseColor = colornames.Mediumblue

	this.scaleX.Model3D.GetMeshIndex(0).GetMaterial().DiffuseColor = colornames.Red
	this.scaleY.Model3D.GetMeshIndex(0).GetMaterial().DiffuseColor = colornames.Lime
	this.scaleZ.Model3D.GetMeshIndex(0).GetMaterial().DiffuseColor = colornames.Mediumblue

	this.translateX.Transform.Rotation = mgl32.QuatRotate(mgl32.DegToRad(-90.0), mgl32.Vec3{0.0, 1.0, 0.0})
	this.translateY.Transform.Rotation = mgl32.QuatRotate(mgl32.DegToRad(90.0), mgl32.Vec3{1.0, 0.0, 0.0})
	this.translateZ.Transform.Rotation = mgl32.QuatRotate(mgl32.DegToRad(180.0), mgl32.Vec3{1.0, 0.0, 0.0})

	this.scaleX.Transform.Rotation = mgl32.QuatRotate(mgl32.DegToRad(180.0), mgl32.Vec3{0.0, 0.0, 1.0})
	this.scaleY.Transform.Rotation = mgl32.QuatRotate(mgl32.DegToRad(-90.0), mgl32.Vec3{0.0, 0.0, 1.0})
	this.scaleZ.Transform.Rotation = mgl32.QuatRotate(mgl32.DegToRad(90.0), mgl32.Vec3{0.0, 1.0, 0.0})

	gohome.RenderMgr.AddObject(&this.translateX)
	gohome.RenderMgr.AddObject(&this.translateY)
	gohome.RenderMgr.AddObject(&this.translateZ)
	gohome.RenderMgr.AddObject(&this.scaleX)
	gohome.RenderMgr.AddObject(&this.scaleY)
	gohome.RenderMgr.AddObject(&this.scaleZ)
	gohome.UpdateMgr.AddObject(this)
}

func (this *Arrows) Update(detla_time float32) {
	cam := camera.Position
	if current_mode == MODE_MOVE {
		txs := this.translateX.Transform.Position.Sub(cam).Len() * ARROWS_SIZE
		tys := this.translateY.Transform.Position.Sub(cam).Len() * ARROWS_SIZE
		tzs := this.translateZ.Transform.Position.Sub(cam).Len() * ARROWS_SIZE

		this.translateX.Transform.Scale = [3]float32{txs, txs, txs}
		this.translateY.Transform.Scale = [3]float32{tys, tys, tys}
		this.translateZ.Transform.Scale = [3]float32{tzs, tzs, tzs}

		this.translateX.Visible = true
		this.translateY.Visible = true
		this.translateZ.Visible = true
	} else {
		this.translateX.Visible = false
		this.translateY.Visible = false
		this.translateZ.Visible = false
	}

	if current_mode == MODE_SCALE {
		sxs := this.scaleX.Transform.Position.Sub(cam).Len() * ARROWS_SIZE
		sys := this.scaleY.Transform.Position.Sub(cam).Len() * ARROWS_SIZE
		szs := this.scaleZ.Transform.Position.Sub(cam).Len() * ARROWS_SIZE

		this.scaleX.Transform.Scale = [3]float32{sxs, sxs, sxs}
		this.scaleY.Transform.Scale = [3]float32{sys, sys, sys}
		this.scaleZ.Transform.Scale = [3]float32{szs, szs, szs}

		this.scaleX.Visible = true
		this.scaleY.Visible = true
		this.scaleZ.Visible = true
	} else {
		this.scaleX.Visible = false
		this.scaleY.Visible = false
		this.scaleZ.Visible = false
	}

}

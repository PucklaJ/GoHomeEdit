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

	this.translateX.Model3D.GetMeshIndex(0).GetMaterial().DiffuseColor = colornames.Red
	this.translateY.Model3D.GetMeshIndex(0).GetMaterial().DiffuseColor = colornames.Lime
	this.translateZ.Model3D.GetMeshIndex(0).GetMaterial().DiffuseColor = colornames.Mediumblue

	this.translateX.Transform.Rotation = mgl32.QuatRotate(mgl32.DegToRad(-90.0), mgl32.Vec3{0.0, 1.0, 0.0})
	this.translateY.Transform.Rotation = mgl32.QuatRotate(mgl32.DegToRad(90.0), mgl32.Vec3{1.0, 0.0, 0.0})
	this.translateZ.Transform.Rotation = mgl32.QuatRotate(mgl32.DegToRad(180.0), mgl32.Vec3{1.0, 0.0, 0.0})

	gohome.RenderMgr.AddObject(&this.translateX)
	gohome.RenderMgr.AddObject(&this.translateY)
	gohome.RenderMgr.AddObject(&this.translateZ)
	gohome.UpdateMgr.AddObject(this)
}

func (this *Arrows) Update(detla_time float32) {
	cam := camera.Position

	txs := this.translateX.Transform.Position.Sub(cam).Len() * ARROWS_SIZE
	tys := this.translateY.Transform.Position.Sub(cam).Len() * ARROWS_SIZE
	tzs := this.translateZ.Transform.Position.Sub(cam).Len() * ARROWS_SIZE

	this.translateX.Transform.Scale = [3]float32{txs, txs, txs}
	this.translateY.Transform.Scale = [3]float32{tys, tys, tys}
	this.translateZ.Transform.Scale = [3]float32{tzs, tzs, tzs}
}

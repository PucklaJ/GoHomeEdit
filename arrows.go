package main

import (
	"github.com/PucklaMotzer09/gohomeengine/src/gohome"
	"github.com/PucklaMotzer09/mathgl/mgl32"
	"golang.org/x/image/colornames"
	"sync"
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

	/*this.rotateX.InitModel(gohome.ResourceMgr.GetLevel("Arrows").GetModel("Block_Cube"))
	this.rotateY.InitModel(gohome.ResourceMgr.GetLevel("Arrows").GetModel("Block_Cube").Copy())
	this.rotateZ.InitModel(gohome.ResourceMgr.GetLevel("Arrows").GetModel("Block_Cube").Copy())*/

	this.translateX.Model3D.GetMeshIndex(0).GetMaterial().DiffuseColor = colornames.Red
	this.translateY.Model3D.GetMeshIndex(0).GetMaterial().DiffuseColor = colornames.Lime
	this.translateZ.Model3D.GetMeshIndex(0).GetMaterial().DiffuseColor = colornames.Mediumblue

	this.scaleX.Model3D.GetMeshIndex(0).GetMaterial().DiffuseColor = colornames.Red
	this.scaleY.Model3D.GetMeshIndex(0).GetMaterial().DiffuseColor = colornames.Lime
	this.scaleZ.Model3D.GetMeshIndex(0).GetMaterial().DiffuseColor = colornames.Mediumblue

	/*this.rotateX.Model3D.GetMeshIndex(0).GetMaterial().DiffuseColor = colornames.Red
	this.rotateY.Model3D.GetMeshIndex(0).GetMaterial().DiffuseColor = colornames.Lime
	this.rotateZ.Model3D.GetMeshIndex(0).GetMaterial().DiffuseColor = colornames.Mediumblue*/

	this.translateX.Transform.Rotation = mgl32.QuatRotate(mgl32.DegToRad(-90.0), mgl32.Vec3{0.0, 1.0, 0.0})
	this.translateY.Transform.Rotation = mgl32.QuatRotate(mgl32.DegToRad(90.0), mgl32.Vec3{1.0, 0.0, 0.0})
	this.translateZ.Transform.Rotation = mgl32.QuatRotate(mgl32.DegToRad(180.0), mgl32.Vec3{1.0, 0.0, 0.0})

	this.scaleX.Transform.Rotation = mgl32.QuatRotate(mgl32.DegToRad(180.0), mgl32.Vec3{0.0, 0.0, 1.0})
	this.scaleY.Transform.Rotation = mgl32.QuatRotate(mgl32.DegToRad(-90.0), mgl32.Vec3{0.0, 0.0, 1.0})
	this.scaleZ.Transform.Rotation = mgl32.QuatRotate(mgl32.DegToRad(90.0), mgl32.Vec3{0.0, 1.0, 0.0})

	/*this.rotateX.Transform.Rotation = mgl32.QuatRotate(mgl32.DegToRad(180.0), mgl32.Vec3{0.0, 0.0, 1.0})
	this.rotateY.Transform.Rotation = mgl32.QuatRotate(mgl32.DegToRad(-90.0), mgl32.Vec3{0.0, 0.0, 1.0})
	this.rotateZ.Transform.Rotation = mgl32.QuatRotate(mgl32.DegToRad(90.0), mgl32.Vec3{0.0, 1.0, 0.0})*/

	this.translateX.Transform.IgnoreParentRotation = true
	this.translateX.Transform.IgnoreParentScale = true
	this.translateX.DepthTesting = false
	this.translateX.RenderLast = true
	this.translateY.Transform.IgnoreParentRotation = true
	this.translateY.Transform.IgnoreParentScale = true
	this.translateY.DepthTesting = false
	this.translateY.RenderLast = true
	this.translateZ.Transform.IgnoreParentRotation = true
	this.translateZ.Transform.IgnoreParentScale = true
	this.translateZ.DepthTesting = false
	this.translateZ.RenderLast = true

	this.scaleX.Transform.IgnoreParentRotation = true
	this.scaleX.Transform.IgnoreParentScale = true
	this.scaleX.DepthTesting = false
	this.scaleX.RenderLast = true
	this.scaleY.Transform.IgnoreParentRotation = true
	this.scaleY.Transform.IgnoreParentScale = true
	this.scaleY.DepthTesting = false
	this.scaleY.RenderLast = true
	this.scaleZ.Transform.IgnoreParentRotation = true
	this.scaleZ.Transform.IgnoreParentScale = true
	this.scaleZ.DepthTesting = false
	this.scaleZ.RenderLast = true

	/*this.rotateX.Transform.IgnoreParentRotation = true
	this.rotateX.Transform.IgnoreParentScale = true
	this.rotateX.DepthTesting = false
	this.rotateX.RenderLast = true
	this.rotateY.Transform.IgnoreParentRotation = true
	this.rotateY.Transform.IgnoreParentScale = true
	this.rotateY.DepthTesting = false
	this.rotateY.RenderLast = true
	this.rotateZ.Transform.IgnoreParentRotation = true
	this.rotateZ.Transform.IgnoreParentScale = true
	this.rotateZ.DepthTesting = false
	this.rotateZ.RenderLast = true*/

	gohome.RenderMgr.AddObject(&this.translateX)
	gohome.RenderMgr.AddObject(&this.translateY)
	gohome.RenderMgr.AddObject(&this.translateZ)
	gohome.RenderMgr.AddObject(&this.scaleX)
	gohome.RenderMgr.AddObject(&this.scaleY)
	gohome.RenderMgr.AddObject(&this.scaleZ)
	/*gohome.RenderMgr.AddObject(&this.rotateX)
	gohome.RenderMgr.AddObject(&this.rotateY)
	gohome.RenderMgr.AddObject(&this.rotateZ)*/
	gohome.UpdateMgr.AddObject(this)
}

var rotate float32 = 0.0

func (this *Arrows) Update(detla_time float32) {
	cam := camera.Position
	if current_mode == MODE_MOVE {
		txs := this.translateX.Transform.GetPosition().Sub(cam).Len() * ARROWS_SIZE
		tys := this.translateY.Transform.GetPosition().Sub(cam).Len() * ARROWS_SIZE
		tzs := this.translateZ.Transform.GetPosition().Sub(cam).Len() * ARROWS_SIZE

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
		sxs := this.scaleX.Transform.GetPosition().Sub(cam).Len() * ARROWS_SIZE
		sys := this.scaleY.Transform.GetPosition().Sub(cam).Len() * ARROWS_SIZE
		szs := this.scaleZ.Transform.GetPosition().Sub(cam).Len() * ARROWS_SIZE

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

	/*if current_mode == MODE_ROTATE {
		rxs := this.rotateX.Transform.GetPosition().Sub(cam).Len() * ARROWS_SIZE
		rys := this.rotateY.Transform.GetPosition().Sub(cam).Len() * ARROWS_SIZE
		rzs := this.rotateZ.Transform.GetPosition().Sub(cam).Len() * ARROWS_SIZE

		this.rotateX.Transform.Scale = [3]float32{rxs, rxs, rxs}
		this.rotateY.Transform.Scale = [3]float32{rys, rys, rys}
		this.rotateZ.Transform.Scale = [3]float32{rzs, rzs, rzs}

		this.rotateX.Visible = true
		this.rotateY.Visible = true
		this.rotateZ.Visible = true
	} else {
		this.rotateX.Visible = false
		this.rotateY.Visible = false
		this.rotateZ.Visible = false
	}*/

	if len(placed_models) != 0 {
		this.SetParent(&placed_models[place_id-1].Entity3D)
	}
}

func (this *Arrows) SetParent(parent interface{}) {
	this.translateX.SetParent(parent)
	this.translateY.SetParent(parent)
	this.translateZ.SetParent(parent)

	this.scaleX.SetParent(parent)
	this.scaleY.SetParent(parent)
	this.scaleZ.SetParent(parent)

	/*this.rotateX.SetParent(parent)
	this.rotateY.SetParent(parent)
	this.rotateZ.SetParent(parent)*/
}

func transformAABB(aabb *gohome.AxisAlignedBoundingBox, transform *gohome.TransformableObject3D, wg *sync.WaitGroup) {
	tmat := transform.GetTransformMatrix()
	vmat := camera.GetViewMatrix()
	pmat := gohome.RenderMgr.Projection3D.GetProjectionMatrix()
	mat := pmat.Mul4(vmat).Mul4(tmat)
	min4 := mat.Mul4x1(aabb.Min.Vec4(1))
	max4 := mat.Mul4x1(aabb.Max.Vec4(1))
	min3 := min4.Div(min4.W()).Vec3()
	max3 := max4.Div(max4.W()).Vec3()
	aabb.Min = min3.Div(min3.Z())
	aabb.Max = max3.Div(max3.Z())
	nres := gohome.Render.GetNativeResolution()

	min, max := aabb.Min.Vec2(), aabb.Max.Vec2()
	min = min.MulVec([2]float32{1.0, -1.0}).Add([2]float32{1.0, 1.0}).DivVec([2]float32{2.0, 2.0}).MulVec(nres)
	max = max.MulVec([2]float32{1.0, -1.0}).Add([2]float32{1.0, 1.0}).DivVec([2]float32{2.0, 2.0}).MulVec(nres)

	aabb.Min = min.Vec3(-1.0)
	aabb.Max = max.Vec3(-1.0)

	wg.Done()
}

func (this *Arrows) getMoveAABBs() (gohome.AxisAlignedBoundingBox, gohome.AxisAlignedBoundingBox, gohome.AxisAlignedBoundingBox) {
	aabbx := this.translateX.Model3D.AABB
	aabby := this.translateY.Model3D.AABB
	aabbz := this.translateZ.Model3D.AABB
	var wg sync.WaitGroup

	wg.Add(5)
	go func() {
		camera.CalculateViewMatrix()
		wg.Done()
	}()
	go func() {
		gohome.RenderMgr.Projection3D.CalculateProjectionMatrix()
		wg.Done()
	}()
	go func() {
		this.translateX.Transform.CalculateTransformMatrix(&gohome.RenderMgr, -1)
		wg.Done()
	}()
	go func() {
		this.translateY.Transform.CalculateTransformMatrix(&gohome.RenderMgr, -1)
		wg.Done()
	}()
	go func() {
		this.translateZ.Transform.CalculateTransformMatrix(&gohome.RenderMgr, -1)
		wg.Done()
	}()
	wg.Wait()

	wg.Add(3)
	go transformAABB(&aabbx, this.translateX.Transform, &wg)
	go transformAABB(&aabby, this.translateY.Transform, &wg)
	go transformAABB(&aabbz, this.translateZ.Transform, &wg)
	wg.Wait()
	return aabbx, aabby, aabbz
}

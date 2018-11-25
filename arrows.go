package main

import (
	"github.com/PucklaMotzer09/gohomeengine/src/gohome"
	"github.com/PucklaMotzer09/mathgl/mgl32"
	"golang.org/x/image/colornames"
	"sync"
)

const (
	ARROWS_SIZE float32 = 0.1

	X_AXIS uint8 = 1
	Y_AXIS uint8 = 2
	Z_AXIS uint8 = 3
)

var (
	X_PLANES = [4]gohome.PlaneMath3D{
		{
			mgl32.Vec3{0.0, 0.0, 1.0},
			mgl32.Vec3{1.0, 1.0, 0.0},
		},
		{
			mgl32.Vec3{0.0, 0.0, -1.0},
			mgl32.Vec3{1.0, 1.0, 0.0},
		},
		{
			mgl32.Vec3{0.0, 1.0, 0.0},
			mgl32.Vec3{1.0, 0.0, 1.0},
		},
		{
			mgl32.Vec3{0.0, -1.0, 0.0},
			mgl32.Vec3{1.0, 0.0, 1.0},
		},
	}

	Y_PLANES = [4]gohome.PlaneMath3D{
		{
			mgl32.Vec3{0.0, 0.0, 1.0},
			mgl32.Vec3{1.0, 1.0, 0.0},
		},
		{
			mgl32.Vec3{0.0, 0.0, -1.0},
			mgl32.Vec3{1.0, 1.0, 0.0},
		},
		{
			mgl32.Vec3{1.0, 0.0, 0.0},
			mgl32.Vec3{0.0, 1.0, 1.0},
		},
		{
			mgl32.Vec3{-1.0, 0.0, 0.0},
			mgl32.Vec3{0.0, 1.0, 1.0},
		},
	}

	Z_PLANES = [4]gohome.PlaneMath3D{
		{
			mgl32.Vec3{0.0, 1.0, 0.0},
			mgl32.Vec3{1.0, 0.0, 1.0},
		},
		{
			mgl32.Vec3{0.0, -1.0, 0.0},
			mgl32.Vec3{1.0, 0.0, 1.0},
		},
		{
			mgl32.Vec3{1.0, 0.0, 0.0},
			mgl32.Vec3{0.0, 1.0, 1.0},
		},
		{
			mgl32.Vec3{-1.0, 0.0, 0.0},
			mgl32.Vec3{0.0, 1.0, 1.0},
		},
	}
)

type Arrows struct {
	gohome.NilRenderObject
	translateX gohome.Entity3D
	translateY gohome.Entity3D
	translateZ gohome.Entity3D

	scaleX gohome.Entity3D
	scaleY gohome.Entity3D
	scaleZ gohome.Entity3D

	rotateX gohome.Entity3D
	rotateY gohome.Entity3D
	rotateZ gohome.Entity3D

	IsTransforming uint8

	points   [2]mgl32.Vec2
	points3D [2]mgl32.Vec3
}

func (this *Arrows) Init() {
	gohome.ResourceMgr.LoadLevel("Arrows", "arrows.obj", true)

	this.translateX.InitModel(gohome.ResourceMgr.GetLevel("Arrows").GetModel("Arrow_Cone"))
	this.translateY.InitModel(gohome.ResourceMgr.GetLevel("Arrows").GetModel("Arrow_Cone").Copy())
	this.translateZ.InitModel(gohome.ResourceMgr.GetLevel("Arrows").GetModel("Arrow_Cone").Copy())

	this.scaleX.InitModel(gohome.ResourceMgr.GetLevel("Arrows").GetModel("Block_Cube.001"))
	this.scaleY.InitModel(gohome.ResourceMgr.GetLevel("Arrows").GetModel("Block_Cube.001").Copy())
	this.scaleZ.InitModel(gohome.ResourceMgr.GetLevel("Arrows").GetModel("Block_Cube.001").Copy())

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
	gohome.RenderMgr.AddObject(this)

	this.IsTransforming = 0
	this.points[0] = [2]float32{1.0, 1.0}
	this.points[1] = [2]float32{2.0, 2.0}
	this.points3D[0] = [3]float32{1.0, 1.0, 1.0}
	this.points3D[1] = [3]float32{2.0, 2.0, 2.0}
}

func (this *Arrows) SetScale() {
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

}

func (this *Arrows) Update(detla_time float32) {
	this.SetScale()
	if !is_transforming && len(placed_models) != 0 {
		this.SetParent(&placed_models[place_id-1].Entity3D)
	} else {
		this.SetParent(nil)
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

	if parent != nil {
		this.translateX.Transform.Position = [3]float32{0.0, 0.0, 0.0}
		this.translateY.Transform.Position = [3]float32{0.0, 0.0, 0.0}
		this.translateZ.Transform.Position = [3]float32{0.0, 0.0, 0.0}

		this.scaleX.Transform.Position = [3]float32{0.0, 0.0, 0.0}
		this.scaleY.Transform.Position = [3]float32{0.0, 0.0, 0.0}
		this.scaleZ.Transform.Position = [3]float32{0.0, 0.0, 0.0}

		/*this.rotateX.Transform.Position = [3]float32{0.0, 0.0, 0.0}
		this.rotateY.Transform.Position = [3]float32{0.0, 0.0, 0.0}
		this.rotateZ.Transform.Position = [3]float32{0.0, 0.0, 0.0}*/
	}
}

func (this *Arrows) calculateAllMatrices() {
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
}

func convert3Dto2D(pos mgl32.Vec3, pos2 *mgl32.Vec2, wg *sync.WaitGroup) {
	vmat := camera.GetViewMatrix()
	pmat := gohome.RenderMgr.Projection3D.GetProjectionMatrix()
	mat := pmat.Mul4(vmat)
	pos4 := mat.Mul4x1(pos.Vec4(1))
	pos3 := pos4.Div(pos4.W()).Vec3()
	nres := gohome.Render.GetNativeResolution()

	*pos2 = pos3.Vec2()
	*pos2 = pos2.MulVec([2]float32{1.0, -1.0}).Add([2]float32{1.0, 1.0}).Div(2.0).MulVec(nres)
	wg.Done()
}

func calculateRectangle(pos, dir mgl32.Vec2, points *[4]mgl32.Vec2, wg *sync.WaitGroup) {
	point := dir.Sub(pos)
	point = mgl32.Rotate2D(mgl32.DegToRad(90.0)).Mul2x1(point).Normalize().Mul(ARROW_WIDTH / 2.0)

	min := pos.Add(point)
	max := dir.Add(point.Mul(-1))

	(*points)[0] = min
	(*points)[1] = min.Add(point.Mul(-2))
	(*points)[2] = max
	(*points)[3] = max.Add(point.Mul(2))
	wg.Done()
}

func calculateRectangles(pos, xdir, ydir, zdir mgl32.Vec2) (pointsx, pointsy, pointsz [4]mgl32.Vec2) {

	var wg sync.WaitGroup
	wg.Add(3)
	go calculateRectangle(pos, xdir, &pointsx, &wg)
	go calculateRectangle(pos, ydir, &pointsy, &wg)
	go calculateRectangle(pos, zdir, &pointsz, &wg)
	wg.Wait()
	return
}

func (this *Arrows) GetMoveHitboxes() (pointsx, pointsy, pointsz [4]mgl32.Vec2) {
	this.calculateAllMatrices()
	pos, xdir, ydir, zdir := this.getMovePosAndDirections()
	pointsx, pointsy, pointsz = calculateRectangles(pos, xdir, ydir, zdir)
	return
}

func (this *Arrows) getMovePosAndDirections() (pos, xdir, ydir, zdir mgl32.Vec2) {
	x := mgl32.Vec3{1.0, 0.0, 0.0}
	y := mgl32.Vec3{0.0, 1.0, 0.0}
	z := mgl32.Vec3{0.0, 0.0, 1.0}
	xpos := this.translateX.Transform.GetPosition()
	ypos := this.translateY.Transform.GetPosition()
	zpos := this.translateZ.Transform.GetPosition()
	scale := this.translateX.Transform.Scale[0]
	var wg sync.WaitGroup
	wg.Add(4)
	go convert3Dto2D(xpos, &pos, &wg)
	go convert3Dto2D(xpos.Add(x.Mul(ARROW_LENGTH*scale)), &xdir, &wg)
	go convert3Dto2D(ypos.Add(y.Mul(ARROW_LENGTH*scale)), &ydir, &wg)
	go convert3Dto2D(zpos.Add(z.Mul(ARROW_LENGTH*scale)), &zdir, &wg)
	wg.Wait()
	return
}

func (this *Arrows) SetPosition() {
	this.translateX.Transform.Position = transform_start_pos
	this.translateY.Transform.Position = transform_start_pos
	this.translateZ.Transform.Position = transform_start_pos

	this.scaleX.Transform.Position = transform_start_pos
	this.scaleY.Transform.Position = transform_start_pos
	this.scaleZ.Transform.Position = transform_start_pos

	/*this.rotateX.Transform.Position = transform_start_pos
	this.rotateY.Transform.Position = transform_start_pos
	this.rotateZ.Transform.Position = transform_start_pos*/
}

func (this *Arrows) ResetPosition() {
	if m, ok := placed_models[place_id-1]; ok {
		this.SetParent(&m.Entity3D)
	}
}

func (this *Arrows) CalculatePoints() {
	var point1, point2 mgl32.Vec3
	var point12D, point22D mgl32.Vec2

	switch this.IsTransforming {
	case X_AXIS:
		point1 = this.translateX.Transform.GetPosition()
		point2 = point1.Add(mgl32.Vec3{1.0, 0.0, 0.0}.Mul(ARROW_LENGTH))
		gohome.DrawColor = colornames.Red
	case Y_AXIS:
		point1 = this.translateY.Transform.GetPosition()
		point2 = point1.Add(mgl32.Vec3{0.0, 1.0, 0.0}.Mul(ARROW_LENGTH))
		gohome.DrawColor = colornames.Lime
	case Z_AXIS:
		point1 = this.translateZ.Transform.GetPosition()
		point2 = point1.Add(mgl32.Vec3{0.0, 0.0, 1.0}.Mul(ARROW_LENGTH))
		gohome.DrawColor = colornames.Mediumblue
	}

	mid := point1.Add(point2.Sub(point1).Mul(0.5))
	left := point1.Sub(mid).Normalize()
	right := point2.Sub(mid).Normalize()

	point1 = mid.Add(left.Mul(ARROW_LINE_LENGTH / 2.0))
	point2 = mid.Add(right.Mul(ARROW_LINE_LENGTH / 2.0))
	this.points3D[0] = point1
	this.points3D[1] = point2
	var wg sync.WaitGroup
	wg.Add(1)
	go convert3Dto2D(point1, &point12D, &wg)
	wg.Wait()
	wg.Add(1)
	go convert3Dto2D(point2, &point22D, &wg)
	wg.Wait()
	this.points[0] = point12D
	this.points[1] = point22D
}

func (this *Arrows) drawHitboxes() {
	pointsx, pointsy, pointsz := this.GetMoveHitboxes()

	gohome.Filled = false
	gohome.DrawColor = colornames.Red
	gohome.DrawRectangle2D(pointsx[0], pointsx[1], pointsx[2], pointsx[3])
	gohome.DrawColor = colornames.Lime
	gohome.DrawRectangle2D(pointsy[0], pointsy[1], pointsy[2], pointsy[3])
	gohome.DrawColor = colornames.Mediumblue
	gohome.DrawRectangle2D(pointsz[0], pointsz[1], pointsz[2], pointsz[3])
}

func (this *Arrows) Render() {

	if this.IsTransforming != 0 {
		this.translateX.Visible = false
		this.translateY.Visible = false
		this.translateZ.Visible = false

		this.scaleX.Visible = false
		this.scaleY.Visible = false
		this.scaleZ.Visible = false

		this.rotateX.Visible = false
		this.rotateY.Visible = false
		this.rotateZ.Visible = false

		switch this.IsTransforming {
		case X_AXIS:
			gohome.DrawColor = colornames.Red
		case Y_AXIS:
			gohome.DrawColor = colornames.Lime
		case Z_AXIS:
			gohome.DrawColor = colornames.Mediumblue
		}

		gohome.DrawLine3D(this.points3D[0], this.points3D[1])
	}
}

package main

import (
	"github.com/PucklaMotzer09/GoHomeEngine/src/gohome"
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
	X_COLOR = colornames.Red
	Y_COLOR = colornames.Lime
	Z_COLOR = colornames.Mediumblue
)

var (
	X_PLANES = [4]gohome.PlaneMath3D{
		{
			Normal: mgl32.Vec3{0.0, 0.0, 1.0},
			Point:  mgl32.Vec3{1.0, 1.0, 0.0},
		},
		{
			Normal: mgl32.Vec3{0.0, 0.0, -1.0},
			Point:  mgl32.Vec3{1.0, 1.0, 0.0},
		},
		{
			Normal: mgl32.Vec3{0.0, 1.0, 0.0},
			Point:  mgl32.Vec3{1.0, 0.0, 1.0},
		},
		{
			Normal: mgl32.Vec3{0.0, -1.0, 0.0},
			Point:  mgl32.Vec3{1.0, 0.0, 1.0},
		},
	}

	Y_PLANES = [4]gohome.PlaneMath3D{
		{
			Normal: mgl32.Vec3{0.0, 0.0, 1.0},
			Point:  mgl32.Vec3{1.0, 1.0, 0.0},
		},
		{
			Normal: mgl32.Vec3{0.0, 0.0, -1.0},
			Point:  mgl32.Vec3{1.0, 1.0, 0.0},
		},
		{
			Normal: mgl32.Vec3{1.0, 0.0, 0.0},
			Point:  mgl32.Vec3{0.0, 1.0, 1.0},
		},
		{
			Normal: mgl32.Vec3{-1.0, 0.0, 0.0},
			Point:  mgl32.Vec3{0.0, 1.0, 1.0},
		},
	}

	Z_PLANES = [4]gohome.PlaneMath3D{
		{
			Normal: mgl32.Vec3{0.0, 1.0, 0.0},
			Point:  mgl32.Vec3{1.0, 0.0, 1.0},
		},
		{
			Normal: mgl32.Vec3{0.0, -1.0, 0.0},
			Point:  mgl32.Vec3{1.0, 0.0, 1.0},
		},
		{
			Normal: mgl32.Vec3{1.0, 0.0, 0.0},
			Point:  mgl32.Vec3{0.0, 1.0, 1.0},
		},
		{
			Normal: mgl32.Vec3{-1.0, 0.0, 0.0},
			Point:  mgl32.Vec3{0.0, 1.0, 1.0},
		},
	}
)

type ArrowEntity3D struct {
	gohome.Entity3D
}

func (this *ArrowEntity3D) Render() {
	gohome.Render.SetWireFrame(false)
	this.Entity3D.Render()
	if gohome.RenderMgr.WireFrameMode {
		gohome.Render.SetWireFrame(true)
	}
}

type Arrows struct {
	gohome.NilRenderObject
	translateX ArrowEntity3D
	translateY ArrowEntity3D
	translateZ ArrowEntity3D

	scaleX ArrowEntity3D
	scaleY ArrowEntity3D
	scaleZ ArrowEntity3D

	rotateX ArrowEntity3D
	rotateY ArrowEntity3D
	rotateZ ArrowEntity3D

	TransformAxis uint8

	points3D [2]mgl32.Vec3
}

func (this *Arrows) Init() {
	gohome.ResourceMgr.LoadLevel("Arrows", "arrows.obj", true)

	this.initMove()
	this.initScale()
	this.initRotate()

	gohome.UpdateMgr.AddObject(this)
	gohome.RenderMgr.AddObject(this)

	this.TransformAxis = 0
}

func (this *Arrows) SetScale() {
	if selected_placed_object == nil {
		return
	}
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		if current_mode == MODE_MOVE {
			wg.Add(2)
			go func() {
				this.setScaleMove()
				wg.Done()
			}()
			go func() {
				this.setVisibleMove()
				wg.Done()
			}()
		} else {
			this.setInvisibleMove()
		}
		wg.Done()
	}()

	go func() {
		if current_mode == MODE_SCALE {
			wg.Add(2)
			go func() {
				this.setScaleScale()
				wg.Done()
			}()
			go func() {
				this.setVisibleScale()
				wg.Done()
			}()
		} else {
			this.setInvisibleScale()
		}
		wg.Done()
	}()
	wg.Done()
	/*go func() {
		if current_mode == MODE_ROTATE {
			wg.Add(2)
			go func() {
				this.setScaleRotate()
				wg.Done()
			}()
			go func() {
				this.setVisibleRotate()
				wg.Done()
			}()
		} else {
			this.setInvisibleRotate()
		}
		wg.Done()
	}()*/

	wg.Wait()
}

func (this *Arrows) Update(detla_time float32) {
	if selected_placed_object == nil {
		this.SetInvisible()
		return
	}
	this.SetScale()
	if !is_transforming && selected_placed_object != nil {
		this.SetParent(selected_placed_object)
	} else {
		this.SetParent(nil)
		this.SetInvisible()
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
		this.centerArrows()
	}
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
	pos, xdir, ydir, zdir := this.getPosAndDirections(&this.translateX)
	pointsx, pointsy, pointsz = calculateRectangles(pos, xdir, ydir, zdir)
	return
}

func (this *Arrows) GetScaleHitboxes() (pointsx, pointsy, pointsz [4]mgl32.Vec2) {
	this.calculateAllMatrices()
	pos, xdir, ydir, zdir := this.getPosAndDirections(&this.scaleX)
	pointsx, pointsy, pointsz = calculateRectangles(pos, xdir, ydir, zdir)
	return
}

func (this *Arrows) getPosAndDirections(arrow *ArrowEntity3D) (pos, xdir, ydir, zdir mgl32.Vec2) {
	x := mgl32.Vec3{1.0, 0.0, 0.0}
	y := mgl32.Vec3{0.0, 1.0, 0.0}
	z := mgl32.Vec3{0.0, 0.0, 1.0}
	xpos := arrow.Transform.GetPosition()
	ypos := xpos
	zpos := xpos
	scale := arrow.Transform.Scale[0]
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
	this.SetParent(selected_placed_object)
}

func (this *Arrows) CalculatePoints() {
	var point1, point2 mgl32.Vec3

	switch this.TransformAxis {
	case X_AXIS:
		point1 = this.translateX.Transform.GetPosition()
		point2 = point1.Add(mgl32.Vec3{1.0, 0.0, 0.0}.Mul(ARROW_LENGTH))
		gohome.DrawColor = X_COLOR
	case Y_AXIS:
		point1 = this.translateY.Transform.GetPosition()
		point2 = point1.Add(mgl32.Vec3{0.0, 1.0, 0.0}.Mul(ARROW_LENGTH))
		gohome.DrawColor = Y_COLOR
	case Z_AXIS:
		point1 = this.translateZ.Transform.GetPosition()
		point2 = point1.Add(mgl32.Vec3{0.0, 0.0, 1.0}.Mul(ARROW_LENGTH))
		gohome.DrawColor = Z_COLOR
	}

	mid := point1.Add(point2.Sub(point1).Mul(0.5))
	left := point1.Sub(mid).Normalize()
	right := point2.Sub(mid).Normalize()

	point1 = mid.Add(left.Mul(ARROW_LINE_LENGTH / 2.0))
	point2 = mid.Add(right.Mul(ARROW_LINE_LENGTH / 2.0))
	this.points3D[0] = point1
	this.points3D[1] = point2
}

func (this *Arrows) drawHitboxes() {
	pointsx, pointsy, pointsz := this.GetMoveHitboxes()

	gohome.Filled = false
	gohome.DrawColor = X_COLOR
	gohome.DrawRectangle2D(pointsx[0], pointsx[1], pointsx[2], pointsx[3])
	gohome.DrawColor = Y_COLOR
	gohome.DrawRectangle2D(pointsy[0], pointsy[1], pointsy[2], pointsy[3])
	gohome.DrawColor = Z_COLOR
	gohome.DrawRectangle2D(pointsz[0], pointsz[1], pointsz[2], pointsz[3])
}

func (this *Arrows) Render() {

	if this.TransformAxis != 0 {
		this.SetInvisible()

		switch this.TransformAxis {
		case X_AXIS:
			gohome.DrawColor = X_COLOR
		case Y_AXIS:
			gohome.DrawColor = Y_COLOR
		case Z_AXIS:
			gohome.DrawColor = Z_COLOR
		}

		gohome.DrawLine3D(this.points3D[0], this.points3D[1])
		if gohome.RenderMgr.WireFrameMode {
			gohome.Render.SetWireFrame(true)
		}
	}
}

func (this *Arrows) SetVisible() {
	this.setVisibleMove()
	this.setVisibleScale()
	this.setVisibleRotate()
}

func (this *Arrows) SetInvisible() {
	this.setInvisibleMove()
	this.setInvisibleScale()
	this.setInvisibleRotate()
}

func (this *Arrows) GetType() gohome.RenderType {
	return gohome.TYPE_2D_NORMAL
}

func (this *Arrows) HasDepthTesting() bool {
	return false
}

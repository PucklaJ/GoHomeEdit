package main

import (
	"encoding/binary"
	"github.com/PucklaMotzer09/gohomeengine/src/frameworks/GTK/gtk"
	"github.com/PucklaMotzer09/gohomeengine/src/gohome"
	"github.com/PucklaMotzer09/mathgl/mgl32"
	"math"
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
	gohome.ResourceMgr.LoadLevelString(name, string(fileContents), fileName, true)

	level := gohome.ResourceMgr.GetLevel(name)
	if level != nil && len(level.LevelObjects) != 0 {
		for i := 0; i < len(level.LevelObjects); i++ {
			model := level.LevelObjects[i].Model3D
			if model != nil {
				if loaded_models == nil {
					loaded_models = make(map[uint32]*gohome.Model3D)
				}
				if placable_models == nil {
					placable_models = make(map[uint32]*PlaceableModel)
				}
				loaded_models[object_id] = model

				lbl := gtk.LabelNew(model.Name)
				lbl.ToGObject().SetData("ID", uint32ToString(object_id))
				lb_assets.Insert(lbl.ToWidget(), -1)
				lbl.ToWidget().Show()

				var pm PlaceableModel
				pm.Name = model.Name
				pm.Filename = fileName
				pm.ID = object_id
				placable_models[object_id] = &pm
				object_id++
			}
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

func handleMoveArrowClick() {
	if selected_placed_object != nil {
		transform_start_pos = selected_placed_object.Transform.GetPosition()
	} else {
		return
	}

	pointsx, pointsy, pointsz := arrows.GetMoveHitboxes()
	quadx, quady, quadz := gohome.QuadMath2D(pointsx), gohome.QuadMath2D(pointsy), gohome.QuadMath2D(pointsz)

	checkMouseIntersections(quadx, quady, quadz)
}

func handleScaleArrowClick() {
	if selected_placed_object != nil {
		transform_start_pos = selected_placed_object.Transform.GetPosition()
		transform_start_scale = selected_placed_object.Transform.Scale
	} else {
		return
	}

	pointsx, pointsy, pointsz := arrows.GetScaleHitboxes()
	quadx, quady, quadz := gohome.QuadMath2D(pointsx), gohome.QuadMath2D(pointsy), gohome.QuadMath2D(pointsz)

	checkMouseIntersections(quadx, quady, quadz)
}

func getPlaneForIntersection(axis uint8) gohome.PlaneMath3D {
	switch axis {
	case X_AXIS:
		return getBestPlane(X_PLANES)
	case Y_AXIS:
		return getBestPlane(Y_PLANES)
	case Z_AXIS:
		return getBestPlane(Z_PLANES)
	}

	return X_PLANES[0]
}

func getAxisProjectedPos(screenPos mgl32.Vec2, axis uint8, m *PlacedObject) mgl32.Vec3 {
	plane := getPlaneForIntersection(axis)
	changePlanePoint(&plane, m.Transform.Position)

	mray := gohome.ScreenPositionToRay(screenPos)
	planePos := mray.PlaneIntersect(camera.Position, plane.Normal, plane.Point)
	planePos = planePos.Project(arrows.points3D[0], arrows.points3D[1])

	return planePos
}

func handleTransforming() {
	if !is_transforming {
		return
	}
	if selected_placed_object == nil {
		return
	}

	arrows.SetPosition()
	arrows.SetInvisible()
	planePos := getAxisProjectedPos(gohome.InputMgr.Mouse.ToScreenPosition(), arrows.TransformAxis, selected_placed_object)

	if current_mode == MODE_MOVE {
		selected_placed_object.Transform.Position = planePos.Sub(transform_start).Add(transform_start_pos)
	} else if current_mode == MODE_SCALE {
		len_all := arrows.points3D[1].Sub(arrows.points3D[0]).Len()
		len_start := transform_start.Sub(arrows.points3D[0]).Len() / len_all
		len_end := planePos.Sub(arrows.points3D[0]).Len() / len_all
		len_dif := len_end - len_start
		cam_dist := camera.Position.Sub(transform_start_pos).Len()
		selected_placed_object.Transform.Scale[arrows.TransformAxis-1] = transform_start_scale[arrows.TransformAxis-1] + len_dif*TRANSFORM_SCALE_SPEED/cam_dist
	}
}

func getBestPlane(planes [4]gohome.PlaneMath3D) gohome.PlaneMath3D {
	minDot := float32(math.Acos(float64(planes[0].Normal.Dot(camera.LookDirection))))
	minIndex := 0

	for i := 1; i < 4; i++ {
		dot := float32(math.Acos(float64(planes[i].Normal.Dot(camera.LookDirection))))
		if dot < minDot {
			minDot = dot
			minIndex = i
		}
	}

	return planes[minIndex]
}

func changePlanePoint(plane *gohome.PlaneMath3D, position mgl32.Vec3) {
	if plane.Normal.X() != 0.0 {
		plane.Point[0] = position.X()
	} else if plane.Normal.Y() != 0.0 {
		plane.Point[1] = position.Y()
	} else {
		plane.Point[2] = position.Z()
	}
}

func handlePlacing() {
	if current_mode != MODE_PLACE || placing_object.Model3D == nil {
		return
	}

	mray := gohome.InputMgr.Mouse.ToRay()

	var plane gohome.PlaneMath3D
	if !place_on_grid {
		plane.Normal = camera.LookDirection.Mul(-1)
		plane.Point = camera.Position.Add(camera.LookDirection.Mul(PLACE_PLANE_DIST))
	} else {
		plane.Normal = mgl32.Vec3{0.0, 1.0, 0.0}
		plane.Point = mgl32.Vec3{1.0, 0.0, 1.0}
	}

	placePoint := mray.PlaneIntersect(camera.Position, plane.Normal, plane.Point)

	placing_object.Transform.Position = placePoint
}

package main

import (
	"encoding/binary"
	"github.com/PucklaMotzer09/gohomeengine/src/frameworks/GTK/gtk"
	"github.com/PucklaMotzer09/gohomeengine/src/gohome"
	"github.com/PucklaMotzer09/mathgl/mgl32"
	"golang.org/x/image/colornames"
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
	level := gohome.ResourceMgr.LoadLevelString(name, string(fileContents), fileName, true)

	if level != nil && len(level.LevelObjects) != 0 {
		for i := 0; i < len(level.LevelObjects); i++ {
			model := level.LevelObjects[i].Model3D
			if model != nil {
				loaded_models[object_id] = model

				lbl := gtk.LabelNew(model.Name)
				lbl.ToGObject().SetData("ID", uint32ToString(object_id))
				lb_assets.Insert(lbl.ToWidget(), -1)
				lbl.ToWidget().Show()

				var pm PlaceableModel
				pm.Name = model.Name
				pm.Filename = fileName
				pm.ID = object_id
				placeable_models[object_id] = &pm
				object_id++

				addEntityModel(pm.PlaceableObject, &pm)
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

func handleMoveArrowClick() bool {
	if selected_placed_object != nil {
		transform_start_pos = selected_placed_object.Transform.GetPosition()
	} else {
		return false
	}

	pointsx, pointsy, pointsz := arrows.GetMoveHitboxes()
	quadx, quady, quadz := gohome.QuadMath2D(pointsx), gohome.QuadMath2D(pointsy), gohome.QuadMath2D(pointsz)

	return checkMouseIntersections(quadx, quady, quadz)
}

func handleScaleArrowClick() bool {
	if selected_placed_object != nil {
		transform_start_pos = selected_placed_object.Transform.GetPosition()
		transform_start_scale = selected_placed_object.Transform.Scale
	} else {
		return false
	}

	pointsx, pointsy, pointsz := arrows.GetScaleHitboxes()
	quadx, quady, quadz := gohome.QuadMath2D(pointsx), gohome.QuadMath2D(pointsy), gohome.QuadMath2D(pointsz)

	return checkMouseIntersections(quadx, quady, quadz)
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

	if gohome.InputMgr.Mouse.DPos[0] != 0 || gohome.InputMgr.Mouse.DPos[1] != 0 {
		gohome.RenderMgr.ReRender = true
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

	if gohome.InputMgr.Mouse.DPos[0] != 0 || gohome.InputMgr.Mouse.DPos[1] != 0 {
		gohome.RenderMgr.ReRender = true
	}
}

func initPickableTexture() {
	pw, ph := gohome.RenderMgr.BackBufferMS.GetWidth(), gohome.RenderMgr.BackBufferMS.GetHeight()
	pickable_texture = gohome.Render.CreateRenderTexture("Pickable Texture", uint32(pw), uint32(ph), 1, true, false, false, false)

	if gohome.Render.HasFunctionAvailable("INSTANCED") {
		gohome.ResourceMgr.LoadShaderSource("Pickable", PICKABLE_INSTANCED_VERTEX_SHADER, PICKABLE_INSTANCED_FRAGMENT_SHADER, "", "", "", "")
	} else {
		gohome.ResourceMgr.LoadShaderSource("Pickable", PICKABLE_VERTEX_SHADER, PICKABLE_FRAGMENT_SHADER, "", "", "", "")
	}
}

func renderPickableTexture() {
	pickable_texture.SetAsTarget()
	gohome.RenderMgr.ForceShader3D = gohome.ResourceMgr.GetShader("Pickable")

	gohome.Render.ClearScreen(colornames.White)
	for _, ent := range instanced_entities {
		ent.Model3D.SetV4(0, pickable_colors[ent])
	}
	gohome.RenderMgr.Render(PICKABLE_BIT, 0, 0, 0)

	gohome.RenderMgr.ForceShader3D = nil
	pickable_texture.UnsetAsTarget()
}

func idToColor(id uint32) (col mgl32.Vec4) {
	col[3] = 1.0 - (float32((id&0xFF000000)>>24) / 255.0)
	col[2] = float32((id&0x00FF0000)>>16) / 255.0
	col[1] = float32((id&0x0000FF00)>>8) / 255.0
	col[0] = float32((id&0x000000FF)>>0) / 255.0

	return
}

func colorToID(col mgl32.Vec4) (id uint32) {
	for i := 0; i < 4; i++ {
		if i == 3 {
			id |= uint32((1.0-col[i])*255.0) << uint32(i*8)
		} else {
			id |= uint32(col[i]*255.0) << uint32(i*8)
		}
	}
	return
}

func handlePickableClick() {
	mpos := gohome.InputMgr.Mouse.ToScreenPosition()
	mpos[1] = gohome.Render.GetNativeResolution()[1] - mpos[1]

	pixel, w, h := pickable_texture.GetData()

	if !(int(mpos[0]) < w && int(mpos[1]) < h && mpos[0] >= 0.0 && mpos[1] >= 0.0) {
		return
	}

	arrayIndex := (int(mpos[0]) + int(mpos[1])*w) * 4

	r := pixel[arrayIndex+0]
	g := pixel[arrayIndex+1]
	b := pixel[arrayIndex+2]
	a := pixel[arrayIndex+3]

	id := colorToID([4]float32{
		float32(r) / 255.0,
		float32(g) / 255.0,
		float32(b) / 255.0,
		float32(a) / 255.0,
	})

	prev := selected_placed_object
	placed_model, ok := placed_models[id]
	if ok {
		selected_placed_object = &placed_model.PlacedObject
		arrows.SetParent(selected_placed_object)
		arrows.SetVisible()
		arrows.SetScale()
	} else {
		selected_placed_object = nil
		arrows.SetInvisible()
	}

	if prev != selected_placed_object {
		gohome.RenderMgr.ReRender = true
	}
}

func addEntityModel(object PlaceableObject, model *PlaceableModel) {
	imodel := gohome.InstancedModel3DFromModel3D(loaded_models[object.ID])
	entity := &gohome.InstancedEntity3D{}
	imodel.AddValue(gohome.VALUE_VEC4)
	imodel.SetName(0, gohome.VALUE_VEC4, "pickableColor")
	entity.InitModel(imodel, 10)
	entity.SetType(gohome.TYPE_3D_INSTANCED | PICKABLE_BIT)
	entity.SetNumUsedInstances(0)
	instanced_entities[object] = entity
	gohome.RenderMgr.AddObject(entity)
	pickable_colors[entity] = make([]mgl32.Vec4, 10)
}

func addEntityInstance(object PlaceableObject) {
	entity := instanced_entities[object]
	index := entity.Model3D.GetNumUsedInstances()
	entity.SetNumUsedInstances(index + 1)
	if entity.Model3D.GetNumUsedInstances() > entity.Model3D.GetNumInstances() {
		entity.SetNumInstances(entity.Model3D.GetNumInstances() + 10)
		entity.SetNumUsedInstances(entity.Model3D.GetNumInstances() - 9)
		pickable_colors[entity] = append(pickable_colors[entity], make([]mgl32.Vec4, 10)...)
	}
	pickable_colors[entity][index] = idToColor(place_id)
}

func handlePlaceClick() {
	pmodel, ok := placeable_models[selected_model]
	if ok {
		entity := instanced_entities[pmodel.PlaceableObject]
		var pm PlacedModel
		index := entity.Model3D.GetNumUsedInstances()
		addEntityInstance(pmodel.PlaceableObject)
		pm.PlacedObject.Transform = &entity.Transforms[index].TransformableObject3D
		pm.PlacedObject.AABB = entity.Model3D.AABB
		pm.PlacedObject.PlaceID = place_id
		pm.PlaceableModel = pmodel
		placed_models[place_id] = &pm

		place_id++

		pm.PlacedObject.Transform.Position = placing_object.Transform.Position
		selected_placed_object = &pm.PlacedObject

		gohome.RenderMgr.ReRender = true
	}
}

func handleMoveClick() {
	if !handleMoveArrowClick() {
		handlePickableClick()
	} else {
		gohome.RenderMgr.ReRender = true
	}
}

func handleScaleClick() {
	if !handleScaleArrowClick() {
		handlePickableClick()
	} else {
		gohome.RenderMgr.ReRender = true
	}
}

func initMaps() {
	placed_models = make(map[uint32]*PlacedModel)
	instanced_entities = make(map[PlaceableObject]*gohome.InstancedEntity3D)
	pickable_colors = make(map[*gohome.InstancedEntity3D][]mgl32.Vec4)
	loaded_models = make(map[uint32]*gohome.Model3D)
	placeable_models = make(map[uint32]*PlaceableModel)
}

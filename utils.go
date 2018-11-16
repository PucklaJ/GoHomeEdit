package main

import (
	"encoding/binary"
	"github.com/PucklaMotzer09/gohomeengine/src/frameworks/GTK/gtk"
	"github.com/PucklaMotzer09/gohomeengine/src/gohome"
	"github.com/PucklaMotzer09/mathgl/mgl32"
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
		loadable_models[i].Life--
		if loadable_models[i].Life == 0 {
			loadModel(loadable_models[i].Name, loadable_models[i].FileContents, loadable_models[i].Filename)
			if len(loadable_models) == 1 {
				loadable_models = loadable_models[:0]
			} else {
				loadable_models = append(loadable_models[:i], loadable_models[i+1:]...)
			}
			i--
		}
	}
}

func updateCamera() {
	updateCameraZoom()
	updateCameraRotation()
}

var smooth_deltas [NUM_SMOOTH_DELTAS][2]float32
var current_smooth_deltas int = NUM_SMOOTH_DELTAS
var current_smooth_delta int = 0

func smoothDeltas() (dx float32, dy float32) {
	var sumx, sumy float32 = 0.0, 0.0
	for i := 0; i < current_smooth_deltas; i++ {
		sumx += smooth_deltas[i][0]
		sumy += smooth_deltas[i][1]
	}
	sumx /= float32(current_smooth_deltas)
	sumy /= float32(current_smooth_deltas)
	dx = sumx
	dy = sumy

	current_smooth_deltas++
	if current_smooth_deltas > NUM_SMOOTH_DELTAS {
		current_smooth_deltas = NUM_SMOOTH_DELTAS
	}
	current_smooth_delta++
	if current_smooth_delta == NUM_SMOOTH_DELTAS {
		current_smooth_delta = 0
	}

	return
}

func resetSmoothDeltas() {
	for i := 0; i < NUM_SMOOTH_DELTAS; i++ {
		smooth_deltas[i][0] = 0.0
		smooth_deltas[i][1] = 0.0
	}
	current_smooth_deltas = NUM_SMOOTH_DELTAS
	current_smooth_delta = 0
}

func updateCameraRotation() {

	dx, dy := float32(gohome.InputMgr.Mouse.DPos[0]), float32(gohome.InputMgr.Mouse.DPos[1])
	smooth_deltas[current_smooth_delta][0] = dx
	smooth_deltas[current_smooth_delta][1] = dy
	dx, dy = smoothDeltas()
	if mgl32.Abs(dx) > MAX_DELTA || mgl32.Abs(dy) > MAX_DELTA {
		return
	}

	yaw, pitch := mgl32.DegToRad(-dx*CAM_ROTATE_VELOCITY), mgl32.DegToRad(dy*CAM_ROTATE_VELOCITY)

	if camera_pitch+pitch > mgl32.DegToRad(88.0) || camera_pitch+pitch < mgl32.DegToRad(-85.0) {
		pitch = 0.0
	}

	if !gohome.InputMgr.IsPressed(gohome.MouseButtonRight) {
		resetSmoothDeltas()
		yaw = 0.0
		pitch = 0.0
	}

	pos := mgl32.Vec3{0.0, 0.0, 1.0}
	look := mgl32.Vec3{0.0, 0.0, -1.0}
	up := mgl32.Vec3{0.0, 1.0, 0.0}
	relVec := pos.Sub(camera_center).Normalize()

	rotateAxisPitch := up.Cross(look).Normalize()
	rotatePitch := mgl32.HomogRotate3D(camera_pitch, rotateAxisPitch)

	rotateAxisYaw := mgl32.Vec3{0.0, 1.0, 0.0}
	rotateYaw := mgl32.HomogRotate3D(camera_yaw, rotateAxisYaw)

	rotate := rotateYaw.Mul4(rotatePitch)

	relVec = rotate.Mul4x1(relVec.Vec4(0.0)).Vec3()

	camera.Position = camera_center.Add(relVec.Mul(camera_zoom))
	camera.LookDirection = camera_center.Sub(camera.Position).Normalize()

	camera_yaw += yaw
	camera_pitch += pitch
}

func updateCameraZoom() {
	wy := float32(gohome.InputMgr.Mouse.Wheel[1])
	camera_zoom -= wy * CAM_ZOOM_VELOCITY
}

package main

import (
	"github.com/PucklaMotzer09/gohomeengine/src/frameworks/GTK/gtk"
	"github.com/PucklaMotzer09/gohomeengine/src/gohome"
	"github.com/PucklaMotzer09/mathgl/mgl32"
)

var lb_assets gtk.ListBox

var camera gohome.Camera3D
var camera_center mgl32.Vec3 = [3]float32{0.0, 0.0, 0.0}
var camera_zoom float32 = MID_ZOOM
var camera_yaw, camera_pitch float32 = 0.0, 0.0

var loaded_models map[uint32]*gohome.Model3D
var loadable_models []LoadableModel
var placable_models map[uint32]*PlaceableModel
var object_id uint32 = 0
var place_id uint32 = 0
var placed_models map[uint32]*PlacedModel

var selected_model uint32

var current_mode Mode = MODE_PLACE

var placing_object PlacingObject

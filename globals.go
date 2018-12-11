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
var camera_yaw, camera_pitch float32 = 3.1415 / 4.0, 3.1415 / 4.0

var loaded_models map[uint32]*gohome.Model3D
var loadable_models []LoadableModel
var placeable_models map[uint32]*PlaceableModel
var object_id uint32 = 0
var place_id uint32 = 0
var placed_models map[uint32]*PlacedModel

var selected_model uint32
var selected_placed_object *PlacedObject

var current_mode Mode = MODE_PLACE

var placing_object PlacingObject
var arrows Arrows

var transform_start mgl32.Vec3
var is_transforming = false

var transform_start_pos mgl32.Vec3
var transform_start_scale mgl32.Vec3

var grid Grid
var place_on_grid = false

var pickable_texture gohome.RenderTexture

var instanced_entities map[PlaceableObject]*gohome.InstancedEntity3D
var pickable_colors map[*gohome.InstancedEntity3D][]mgl32.Vec4

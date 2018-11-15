package main

import (
	"github.com/PucklaMotzer09/gohomeengine/src/frameworks/GTK/gtk"
	"github.com/PucklaMotzer09/gohomeengine/src/gohome"
)

var lb_assets gtk.ListBox

var camera gohome.Camera3D

var loaded_models []*gohome.Model3D
var loadable_models []LoadableModel

package main

import (
	"github.com/PucklaMotzer09/gohomeengine/src/gohome"
)

type CameraUpdater struct {
}

func (this *CameraUpdater) Init() {
	gohome.UpdateMgr.AddObject(this)
}

func (this *CameraUpdater) Update(delta_time float32) {
	updateCamera()
}

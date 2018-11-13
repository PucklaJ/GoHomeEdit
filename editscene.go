package main

import (
	"github.com/PucklaMotzer09/gohomeengine/src/gohome"
)

type EditScene struct {
}

func (this *EditScene) Init() {
	gohome.ErrorMgr.Log("Scene", "EditScene", "Initialised!")
}

func (this *EditScene) Update(delta_time float32) {

}

func (this *EditScene) Terminate() {

}

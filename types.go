package main

import (
	"github.com/PucklaMotzer09/GoHomeEngine/src/gohome"
)

type Mode uint8

const (
	MODE_PLACE  Mode = iota
	MODE_MOVE   Mode = iota
	MODE_ROTATE Mode = iota
	MODE_SCALE  Mode = iota
)

type LoadableModel struct {
	Name         string
	FileContents string
	Filename     string
}

type PlaceableObject struct {
	Name string
	ID   uint32
}

type PlaceableModel struct {
	PlaceableObject
	Filename string
}

type PlacedObject struct {
	PlaceID   uint32
	Transform *gohome.TransformableObject3D
	AABB      gohome.AxisAlignedBoundingBox
}

func (this *PlacedObject) GetTransform3D() *gohome.TransformableObject3D {
	return this.Transform
}

func (this *PlacedObject) SetChildChannel(channel chan bool, tobj *gohome.TransformableObject3D) {
	this.Transform.SetChildChannel(channel, tobj)
}

type PlacedModel struct {
	PlacedObject
	*PlaceableModel
}

package main

import (
	"github.com/PucklaMotzer09/gohomeengine/src/gohome"
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

type PlacedModel struct {
	gohome.Entity3D
	PlacedObject
	*PlaceableModel
	Filename string
}

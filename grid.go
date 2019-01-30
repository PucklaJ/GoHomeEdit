package main

import (
	"github.com/PucklaMotzer09/GoHomeEngine/src/gohome"
	"golang.org/x/image/colornames"
)

const (
	NUM_GRID_X  = 20
	NUM_GRID_Z  = 20
	GRID_SIZE_X = 1
	GRID_SIZE_Z = 1

	GRID_START_X = -NUM_GRID_X / 2 * GRID_SIZE_X
	GRID_START_Z = -NUM_GRID_Z / 2 * GRID_SIZE_Z
	GRID_END_X   = -GRID_START_X
	GRID_END_Z   = -GRID_START_Z

	NUM_LINES = NUM_GRID_X * NUM_GRID_Z
)

var (
	GRID_LINE_COLOR = colornames.Darkgray
)

type Grid struct {
	gohome.Shape3D
}

func (this *Grid) Init() {
	this.Shape3D.Init()

	lines := make([]gohome.Line3D, NUM_LINES)

	for x := GRID_START_X + GRID_SIZE_X; x < GRID_END_X; x += GRID_SIZE_X {
		var line gohome.Line3D
		line[0][0] = float32(x)
		line[0][1] = 0.0
		line[0][2] = GRID_START_Z

		line[1][0] = float32(x)
		line[1][1] = 0.0
		line[1][2] = GRID_END_Z

		line.SetColor(GRID_LINE_COLOR)
		lines = append(lines, line)
	}

	for z := GRID_START_Z + GRID_SIZE_Z; z < GRID_END_Z; z += GRID_SIZE_Z {
		var line gohome.Line3D
		line[0][0] = GRID_START_X
		line[0][1] = 0.0
		line[0][2] = float32(z)

		line[1][0] = GRID_END_X
		line[1][1] = 0.0
		line[1][2] = float32(z)

		line.SetColor(GRID_LINE_COLOR)
		lines = append(lines, line)
	}

	this.AddLines(lines)
	this.Load()
	this.SetDrawMode(gohome.DRAW_MODE_LINES)
	gohome.RenderMgr.AddObject(this)
}

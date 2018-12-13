package main

import (
	"github.com/PucklaMotzer09/GoHomeEngine/src/gohome"
	"github.com/PucklaMotzer09/mathgl/mgl32"
)

type intersect_data struct {
	axis       uint8
	intersects bool
}

func checkIntersection(quad gohome.QuadMath2D, mpos mgl32.Vec2, axis uint8, intersect_channel chan intersect_data) {
	intersect := quad.IntersectsPoint(mpos)
	intersect_channel <- intersect_data{
		axis,
		intersect,
	}
}

func finishIntersection(axis uint8, mpos mgl32.Vec2, m *PlacedObject) {
	arrows.TransformAxis = axis
	is_transforming = true
	arrows.CalculatePoints()
	transform_start = getAxisProjectedPos(mpos, axis, m)

}

func allIntersected(intersected [3]uint8) bool {
	for i := 0; i < 3; i++ {
		if intersected[i] == 2 {
			return false
		}
	}
	return true
}

func checkMouseIntersections(quadx, quady, quadz gohome.QuadMath2D) bool {
	intersect_channel := make(chan intersect_data)
	var intersected = [3]uint8{2, 2, 2}
	m := selected_placed_object
	mpos := gohome.InputMgr.Mouse.ToScreenPosition()

	go checkIntersection(quadx, mpos, X_AXIS, intersect_channel)
	go checkIntersection(quady, mpos, Y_AXIS, intersect_channel)
	go checkIntersection(quadz, mpos, Z_AXIS, intersect_channel)

	for !allIntersected(intersected) {
		data := <-intersect_channel
		if data.intersects {
			intersected[data.axis-1] = 1
		} else {
			intersected[data.axis-1] = 0
		}
	}

	close(intersect_channel)

	for i := 0; i < 3; i++ {
		if intersected[i] == 1 {
			finishIntersection(uint8(i+1), mpos, m)
			return true
		}
	}

	return false
}

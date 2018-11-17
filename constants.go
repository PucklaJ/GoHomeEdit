package main

const (
	MIN_ZOOM float32 = 1.0
	MAX_ZOOM float32 = 50.0
	MID_ZOOM float32 = (MAX_ZOOM-MIN_ZOOM)/2.0 + MIN_ZOOM

	CAM_ROTATE_VELOCITY float32 = 0.5
	CAM_ZOOM_VELOCITY   float32 = 0.5
	CAM_PAN_VELOCITY    float32 = 0.025
	MAX_DELTA           float32 = 200.0

	NUM_SMOOTH_DELTAS int = 5
	NUM_SMOOTH_ZOOM   int = 10
	NUM_SMOOTH_PAN    int = 5
)

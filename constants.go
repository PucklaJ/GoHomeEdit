package main

import (
	"github.com/PucklaMotzer09/GoHomeEngine/src/gohome"
)

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

	ARROW_LENGTH      float32 = 2.3
	ARROW_WIDTH       float32 = 10.0
	ARROW_LINE_LENGTH float32 = 200.0

	TRANSFORM_SCALE_SPEED float32 = 1000.0

	PICKABLE_BIT gohome.RenderType = (1 << 5)
)

var PLACE_PLANE_DIST float32 = MID_ZOOM

const (
	PICKABLE_VERTEX_SHADER = `
	#version 110
	
	attribute vec3 vertex;
	
	uniform mat4 transformMatrix3D;
	uniform mat4 viewMatrix3D;
	uniform mat4 projectionMatrix3D;
	
	void main()
	{
		gl_Position = projectionMatrix3D*viewMatrix3D*transformMatrix3D*vec4(vertex,1.0);
	}`

	PICKABLE_FRAGMENT_SHADER = `
	#version 110
	
	uniform vec4 pickableColor;
	
	void main()
	{
		gl_FragColor = pickableColor;
	}`

	PICKABLE_INSTANCED_VERTEX_SHADER = `
	#version 110
	
	attribute vec3 vertex;
	attribute vec3 normal;
	attribute vec2 texCoord;
	attribute vec3 tangent;
	attribute mat4 transformMatrix3D;
	attribute vec4 pickableColor;

	varying vec4 fragPickableColor;
	
	uniform mat4 viewMatrix3D;
	uniform mat4 projectionMatrix3D;
	
	void main()
	{
		gl_Position = projectionMatrix3D*viewMatrix3D*transformMatrix3D*vec4(vertex,1.0);
		fragPickableColor = pickableColor;
	}`

	PICKABLE_INSTANCED_FRAGMENT_SHADER = `
	#version 110
	
	varying vec4 fragPickableColor;
	
	void main()
	{
		gl_FragColor = fragPickableColor;
	}`
)

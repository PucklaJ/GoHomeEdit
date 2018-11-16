package main

import (
	framework "github.com/PucklaMotzer09/gohomeengine/src/frameworks/GTK"
	"github.com/PucklaMotzer09/gohomeengine/src/frameworks/GTK/gtk"
	"github.com/PucklaMotzer09/gohomeengine/src/gohome"
	"github.com/PucklaMotzer09/mathgl/mgl32"
	"golang.org/x/image/colornames"
)

type EditScene struct {
	cube gohome.Entity3D // Test
}

func (this *EditScene) InitGUI() {
	builder := gtk.BuilderNew()
	if err := builder.AddFromFile("editscene.ui"); err != nil {
		gohome.ErrorMgr.MessageError(gohome.ERROR_LEVEL_ERROR, "Builder", "builder.ui", err)
		return
	}

	window := builder.GetObject("window").ToWidget().ToWindow()
	glarea := builder.GetObject("glarea").ToGLArea()
	lb_assets = builder.GetObject("lb_assets").ToListBox()

	gohome.Framew.(*framework.GTKFramework).InitExternalDefault(&window, &glarea)

	glarea.ToWidget().SignalConnect("size-allocate", updateResolution)
	lb_assets.SignalConnect("row-selected", onSelectAsset)
	builder.GetObject("menu_quit").ToMenuItem().SignalConnect("activate", quitApplication)
	builder.GetObject("tool_place").ToToolButton().SignalConnect(onToolPlace)
	builder.GetObject("tool_move").ToToolButton().SignalConnect(onToolMove)
	builder.GetObject("tool_rotate").ToToolButton().SignalConnect(onToolRotate)
	builder.GetObject("tool_scale").ToToolButton().SignalConnect(onToolScale)
	builder.GetObject("tool_load_model").ToToolButton().SignalConnect(onToolLoadModel)
}

func (this *EditScene) InitGraphics() {
	gohome.Render.SetBackgroundColor(colornames.Lime)
	gohome.Init3DShaders()
	gohome.RenderMgr.UpdateProjectionWithViewport = true
	gohome.LightMgr.DisableLighting()

	camera.Init()
	gohome.RenderMgr.SetCamera3D(&camera, 0)
}

func (this *EditScene) InitTest() {
	gohome.ResourceMgr.LoadTexture("CubeImage", "cube.png")

	mesh := gohome.Box("Cube", [3]float32{1.0, 1.0, 1.0})
	mesh.GetMaterial().SetTextures("CubeImage", "", "")
	this.cube.InitMesh(mesh)
	this.cube.Transform.Position = [3]float32{0.0, 0.0, -3.0}

	gohome.RenderMgr.AddObject(&this.cube)
}

func (this *EditScene) Init() {
	this.InitGUI()
	this.InitGraphics()
	this.InitTest()
}

func (this *EditScene) Update(delta_time float32) {
	this.cube.Transform.Rotation = this.cube.Transform.Rotation.Mul(mgl32.QuatRotate(mgl32.DegToRad(30.0)*delta_time, mgl32.Vec3{0.0, 1.0, 0.0})).Mul(mgl32.QuatRotate(mgl32.DegToRad(30.0)*delta_time, mgl32.Vec3{1.0, 0.0, 0.0}))
	loadLoadableModels()

	if gohome.InputMgr.JustPressed(gohome.MouseButtonLeft) {
		if !lb_assets.ToWidget().HasFocus() {
			onLeftClick()
		}
	}
}

func (this *EditScene) Terminate() {

}

package main

import (
	"strconv"

	framework "github.com/PucklaMotzer09/gohomeengine/src/frameworks/GTK"
	"github.com/PucklaMotzer09/gohomeengine/src/frameworks/GTK/gtk"
	"github.com/PucklaMotzer09/gohomeengine/src/gohome"
	"github.com/PucklaMotzer09/mathgl/mgl32"
	"golang.org/x/image/colornames"
)

type EditScene struct {
	cube      gohome.Entity3D // Test
	lb_assets gtk.ListBox
}

func (this *EditScene) InitGUI() {
	builder := gtk.BuilderNew()
	if err := builder.AddFromFile("editscene.ui"); err != nil {
		gohome.ErrorMgr.MessageError(gohome.ERROR_LEVEL_ERROR, "Builder", "builder.ui", err)
		return
	}

	window := builder.GetObject("window").ToWidget().ToWindow()
	glarea := builder.GetObject("glarea").ToGLArea()
	this.lb_assets = builder.GetObject("lb_assets").ToListBox()
	menu_quit := builder.GetObject("menu_quit").ToMenuItem()

	gohome.Framew.(*framework.GTKFramework).InitExternalDefault(&window, &glarea)
	glarea.ToWidget().Show()

	glarea.ToWidget().SignalConnect("size-allocate", func(widget gtk.Widget) {
		w, h := widget.GetSize()
		gohome.Render.SetNativeResolution(uint32(w), uint32(h))
	})
	menu_quit.SignalConnect("activate", func(menuItem gtk.MenuItem) {
		gohome.MainLop.Quit()
	})
}

func (this *EditScene) InitGraphics() {
	gohome.Render.SetBackgroundColor(colornames.Lime)
	gohome.Init3DShaders()
	gohome.RenderMgr.UpdateProjectionWithViewport = true
	gohome.LightMgr.DisableLighting()
}

func (this *EditScene) InitTest() {
	gohome.ResourceMgr.LoadTexture("CubeImage", "cube.png")

	mesh := gohome.Box("Cube", [3]float32{1.0, 1.0, 1.0})
	mesh.GetMaterial().SetTextures("CubeImage", "", "")
	this.cube.InitMesh(mesh)
	this.cube.Transform.Position = [3]float32{0.0, 0.0, -3.0}

	gohome.RenderMgr.AddObject(&this.cube)
	for i := 0; i < 20; i++ {
		lbl := gtk.LabelNew("Test Asset " + strconv.FormatInt(int64(i), 2))
		this.lb_assets.Insert(lbl.ToWidget(), -1)
		lbl.ToWidget().Show()
	}

}

func (this *EditScene) Init() {
	this.InitGUI()
	this.InitGraphics()
	this.InitTest()
}

func (this *EditScene) Update(delta_time float32) {
	this.cube.Transform.Rotation = this.cube.Transform.Rotation.Mul(mgl32.QuatRotate(mgl32.DegToRad(30.0)*delta_time, mgl32.Vec3{0.0, 1.0, 0.0})).Mul(mgl32.QuatRotate(mgl32.DegToRad(30.0)*delta_time, mgl32.Vec3{1.0, 0.0, 0.0}))
}

func (this *EditScene) Terminate() {

}

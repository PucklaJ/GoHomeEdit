package main

import (
	"fmt"

	"github.com/PucklaMotzer09/gohomeengine/src/frameworks/GTK"
	"github.com/PucklaMotzer09/gohomeengine/src/frameworks/GTK/gtk"
	"github.com/PucklaMotzer09/gohomeengine/src/gohome"
	"github.com/PucklaMotzer09/mathgl/mgl32"
	"golang.org/x/image/colornames"
)

type EditScene struct {
	cube gohome.Entity3D
}

func (this *EditScene) Init() {
	gohome.ErrorMgr.Log("Scene", "EditScene", "Initialised!")

	builder := gtk.BuilderNew()
	if err := builder.AddFromFile("builder.ui"); err != nil {
		gohome.ErrorMgr.MessageError(gohome.ERROR_LEVEL_ERROR, "Builder", "builder.ui", err)
		return
	}

	window := builder.GetObject("window").ToWidget().ToWindow()
	window.ConfigureParameters()
	window.ConnectSignals()
	glarea := builder.GetObject("glarea").ToGLArea()
	glarea.Configure()
	gtk.SetGLArea(glarea)
	/*gtk.CreateGLArea()
	window.ToContainer().GetChildren().Data().ToWidget().ToGrid().Attach(gtk.GetGLArea().ToWidget(), 1, 1, 1, 1)*/
	gtk.SetWindow(window)
	gtkf := gohome.Framew.(*framework.GTKFramework)
	window.ToWidget().ShowAll()
	gtkf.AfterWindowCreation(&gohome.MainLop)

	gohome.Render.SetBackgroundColor(colornames.Lime)
	gohome.Init3DShaders()
	gohome.ResourceMgr.LoadTexture("CubeImage", "cube.png")

	mesh := gohome.Box("Cube", [3]float32{1.0, 1.0, 1.0})
	mesh.GetMaterial().SetTextures("CubeImage", "", "")
	this.cube.InitMesh(mesh)
	this.cube.Transform.Position = [3]float32{0.0, 0.0, -3.0}

	gohome.RenderMgr.AddObject(&this.cube)
	gohome.LightMgr.DisableLighting()
}

func (this *EditScene) Update(delta_time float32) {
	if gohome.InputMgr.IsPressed(gohome.MouseButtonLeft) {
		fmt.Println("Pressed")
	}
	this.cube.Transform.Rotation = this.cube.Transform.Rotation.Mul(mgl32.QuatRotate(mgl32.DegToRad(30.0)*delta_time, mgl32.Vec3{0.0, 1.0, 0.0})).Mul(mgl32.QuatRotate(mgl32.DegToRad(30.0)*delta_time, mgl32.Vec3{1.0, 0.0, 0.0}))
}

func (this *EditScene) Terminate() {

}

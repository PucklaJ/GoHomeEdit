package main

import (
	framework "github.com/PucklaMotzer09/gohomeengine/src/frameworks/GTK"
	"github.com/PucklaMotzer09/gohomeengine/src/frameworks/GTK/gtk"
	"github.com/PucklaMotzer09/gohomeengine/src/gohome"
	"github.com/PucklaMotzer09/mathgl/mgl32"
	"golang.org/x/image/colornames"
	"io/ioutil"
	"os"
)

type EditScene struct {
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
	builder.GetObject("menu_wireframe").ToMenuItem().SignalConnect("activate", onMenuWireframe)
	builder.GetObject("tool_place").ToToolButton().SignalConnect(onToolPlace)
	builder.GetObject("tool_move").ToToolButton().SignalConnect(onToolMove)
	builder.GetObject("tool_rotate").ToToolButton().SignalConnect(onToolRotate)
	builder.GetObject("tool_scale").ToToolButton().SignalConnect(onToolScale)
	builder.GetObject("tool_load_model").ToToolButton().SignalConnect(onToolLoadModel)
}

func (this *EditScene) InitGraphics() {
	gohome.Render.SetBackgroundColor(colornames.Lightgray)
	gohome.Init3DShaders()
	gohome.RenderMgr.UpdateProjectionWithViewport = true
	gohome.LightMgr.DisableLighting()

	camera.Init()
	camera.LookAt(mgl32.Vec3{0.0, 0.0, MID_ZOOM}, camera_center, mgl32.Vec3{0.0, 1.0, 0.0})
	gohome.RenderMgr.SetCamera3D(&camera, 0)
	updateResolution(gtk.GetGLArea().ToWidget())

	var cameraUpdater CameraUpdater
	cameraUpdater.Init()
	arrows.Init()
	placing_object.Init()
}

func (this *EditScene) InitTest() {
	coin, _ := os.Open("files/Coin.obj")
	gopher, _ := os.Open("files/gopher.obj")
	hammer, _ := os.Open("files/Hammer.obj")
	meat, _ := os.Open("files/Meat.obj")
	sword, _ := os.Open("files/Sword.obj")

	coinc, _ := ioutil.ReadAll(coin)
	gopherc, _ := ioutil.ReadAll(gopher)
	hammerc, _ := ioutil.ReadAll(hammer)
	meatc, _ := ioutil.ReadAll(meat)
	swordc, _ := ioutil.ReadAll(sword)

	coin.Close()
	gopher.Close()
	hammer.Close()
	meat.Close()
	sword.Close()

	loadable_models = append(loadable_models, []LoadableModel{
		{
			"Coin.obj", string(coinc), "files/Coin.obj",
		},
		{
			"gopher.obj", string(gopherc), "files/gopher.obj",
		},
		{
			"Hammer.obj", string(hammerc), "files/Hammer.obj",
		},
		{
			"Meat.obj", string(meatc), "files/Meat.obj",
		},
		{
			"Sword.obj", string(swordc), "files/Sword.obj",
		},
	}...)
}

func (this *EditScene) Init() {
	this.InitGUI()
	this.InitGraphics()
	this.InitTest()
}

func (this *EditScene) Update(delta_time float32) {
	loadLoadableModels()

	if gohome.InputMgr.JustPressed(gohome.MouseButtonLeft) {
		if !lb_assets.ToWidget().HasFocus() {
			onLeftClick()
		}
	}

}

func (this *EditScene) Terminate() {

}

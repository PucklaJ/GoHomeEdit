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

	gohome.ResourceMgr.LoadShaderSource("PlacingObject", gohome.ENTITY_3D_SHADER_VERTEX_SOURCE_OPENGL, PLACING_OBJECT_SHADER_FRAGMENT_SOURCE_OPENGL, "", "", "", "")
	if gohome.ResourceMgr.GetShader("PlacingObject") == nil {
		gohome.ResourceMgr.LoadShaderSource("PlacingObjectNoShadow", gohome.ENTITY_3D_NO_SHADOWS_SHADER_VERTEX_SOURCE_OPENGL, PLACING_OBJECT_NO_SHADOWS_SHADER_FRAGMENT_SOURCE_OPENGL, "", "", "", "")
		if gohome.ResourceMgr.GetShader("PlacingObjectNoShadow") != nil {
			gohome.ResourceMgr.SetShader("PlacingObject", "PlacingObjectNoShadow")
		}
	}
	gohome.ResourceMgr.LoadShaderSource("PlacingObjectNoUV", gohome.ENTITY_3D_NOUV_SHADER_VERTEX_SOURCE_OPENGL, PLACING_MODEL_NOUV_SHADER_FRAGMENT_SOURCE_OPENGL, "", "", "", "")
	if gohome.ResourceMgr.GetShader("PlacingObjectNoUV") == nil {
		gohome.ResourceMgr.LoadShaderSource("PlacingObjectNoUVNoShadow", gohome.ENTITY_3D_NOUV_SHADER_VERTEX_SOURCE_OPENGL, PLACING_OBJECT_NOUV_NO_SHADOWS_SHADER_FRAGMENT_SOURCE_OPENGL, "", "", "", "")
		if gohome.ResourceMgr.GetShader("PlacingObjectNoUVNoShadow") != nil {
			gohome.ResourceMgr.SetShader("PlacingObjectNoUV", "PlacingObjectNoUVNoShadow")
		}
	}

	placing_object.Visible = false
	placing_object.RenderLast = true
	gohome.RenderMgr.AddObject(&placing_object)

	var arrows Arrows
	arrows.Init()
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
	updateCamera()

	if gohome.InputMgr.JustPressed(gohome.MouseButtonLeft) {
		if !lb_assets.ToWidget().HasFocus() {
			onLeftClick()
		}
	}

	updatePlacingObject()
}

func (this *EditScene) Terminate() {

}

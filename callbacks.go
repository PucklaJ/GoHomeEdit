package main

import (
	"github.com/PucklaMotzer09/gohomeengine/src/frameworks/GTK/gtk"
	"github.com/PucklaMotzer09/gohomeengine/src/gohome"
)

func updateResolution(widget gtk.Widget) {
	w, h := widget.GetSize()
	gohome.Render.SetNativeResolution(uint32(w), uint32(h))
}

func quitApplication(menuItem gtk.MenuItem) {
	gohome.MainLop.Quit()
}

func onToolPlace(toolButton gtk.ToolButton) {
	lb := gtk.LabelNew("ToolPlace")
	lb_assets.Insert(lb.ToWidget(), -1)
	lb.ToWidget().Show()
}

func onToolMove(toolButton gtk.ToolButton) {
	lb := gtk.LabelNew("ToolMove")
	lb_assets.Insert(lb.ToWidget(), -1)
	lb.ToWidget().Show()
}

func onToolRotate(toolButton gtk.ToolButton) {
	lb := gtk.LabelNew("ToolRotate")
	lb_assets.Insert(lb.ToWidget(), -1)
	lb.ToWidget().Show()
}

func onToolScale(toolButton gtk.ToolButton) {
	lb := gtk.LabelNew("ToolPlace")
	lb_assets.Insert(lb.ToWidget(), -1)
	lb.ToWidget().Show()
}

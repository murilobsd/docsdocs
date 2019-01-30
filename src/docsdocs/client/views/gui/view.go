package gui

import (
	"docsdocs/log"
	"os"

	"github.com/gotk3/gotk3/gtk"
)

const windowName = "window"
const uiMain = "assets/client/glade/main.glade"

// GuiView construct view
type GuiView struct {
	Window *gtk.Window
	Menu   *gtk.MenuItem
	About  *gtk.AboutDialog
	log.Logger
}

func NewGuiView() *GuiView {
	gtk.Init(&os.Args)
	b, _ := gtk.BuilderNewFromFile(uiMain)
	winObj, _ := b.GetObject(windowName)
	window, _ := winObj.(*gtk.Window)
	abtObj, _ := b.GetObject("about")
	about, _ := abtObj.(*gtk.AboutDialog)
	menuObj, _ := b.GetObject("btn_about")
	menu, _ := menuObj.(*gtk.MenuItem)
	return &GuiView{
		Window: window,
		About:  about,
		Menu:   menu,
		Logger: log.NewDocsLogger(),
	}
}

func (g *GuiView) ShowAbout() {
	g.About.Show()
	if g.About.Run() == -4 {
		g.About.Hide()
	}
}

func (g *GuiView) Run() {
	g.Window.Connect("destroy", destroy)
	g.Menu.Connect("activate", g.ShowAbout)
	g.Window.ShowAll()
	// g.ShowAbout()
	gtk.Main()
}

func destroy() {
	gtk.MainQuit()
}

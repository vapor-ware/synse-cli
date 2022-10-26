package ui

import (
	"github.com/rivo/tview"
	"sync"
)

type App struct {
	*tview.Application

	Main    *tview.Pages
	views   map[string]tview.Primitive
	mx      sync.RWMutex
	running bool
}

func NewApp() *App {
	a := &App{
		Application: tview.NewApplication(),
		Main:        tview.NewPages(),
	}
	a.views = map[string]tview.Primitive{
		"logo": NewLogo(),
	}
	return a
}

func (a *App) Run() error {
	if err := a.Application.Run(); err != nil {
		return err
	}
	return nil
}

func (a *App) Init() {
	a.SetRoot(a.Main, true).EnableMouse(true)
}

func (a *App) SetRunning(r bool) {
	a.mx.Lock()
	defer a.mx.Unlock()
	a.running = r
}

func (a *App) Views() map[string]tview.Primitive {
	return a.views
}

func (a *App) QueueUpdate(f func()) {
	if a.Application == nil {
		return
	}
	go func() {
		a.Application.QueueUpdate(f)
	}()
}

func (a *App) QueueUpdateDraw(f func()) {
	if a.Application == nil {
		return
	}
	go func() {
		a.Application.QueueUpdateDraw(f)
	}()
}

func (a *App) Logo() *Logo {
	return a.views["logo"].(*Logo)
}

func (a *App) Conn() {

}

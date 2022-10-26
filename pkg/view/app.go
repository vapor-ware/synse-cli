package view

import (
	"context"
	"github.com/rivo/tview"
	"github.com/vapor-ware/synse-cli/pkg/client"
	"github.com/vapor-ware/synse-cli/pkg/ui"
	"time"
)

const (
	splashDelay     = 1 * time.Second
	refreshInterval = 5 * time.Second
)

type Instance struct {
	*ui.App
	Content   *ui.Pages
	APIClient *client.APIClient

	cancelFunc context.CancelFunc
}

func NewInstance() *Instance {
	i := &Instance{
		App:     ui.NewApp(),
		Content: ui.NewPages(),
	}
	i.Views()["server"] = NewServerInfo(i)
	i.Views()["actions"] = NewActions(i)
	return i
}

func (i *Instance) Init() error {
	ctx := context.Background()
	i.App.Init()

	// TODO: Check if server is active first
	i.serverInfo().Init()
	i.actions().Init()

	i.layout(ctx)
	return nil
}

func (i *Instance) Run() error {
	i.Resume()

	go func() {
		<-time.After(splashDelay)
		i.QueueUpdateDraw(func() {
			i.Main.SwitchToPage("main")
		})
	}()

	// set commands
	i.SetRunning(true)
	if err := i.Application.Run(); err != nil {
		return err
	}
	return nil
}

func (i *Instance) layout(ctx context.Context) {
	main := tview.NewFlex().SetDirection(tview.FlexRow)

	main.AddItem(i.header(), 0, 1, false)

	t := ui.NewTable(i.APIClient)
	go i.QueueUpdateDraw(func() {
		err := t.Watch(ctx)
		if err != nil {
			return
		}
	})

	main.AddItem(t, 0, 2, false)

	i.Main.AddPage("main", main, true, false)
	i.Main.AddPage("splash", ui.NewSplash(), true, true)
}

func (i *Instance) header() tview.Primitive {
	header := tview.NewFlex()
	header.SetDirection(tview.FlexColumn)

	header.SetBorder(true)
	header.SetDirection(tview.FlexColumn)
	header.AddItem(i.serverInfo(), 0, 1, false)
	header.AddItem(i.actions(), 0, 1, false)
	header.AddItem(i.Logo(), 0, 1, false)

	// header.AddItem(i.Menu(), 0, 1, false)
	return header
}

func (i *Instance) serverInfo() *ServerInfo {
	return i.Views()["server"].(*ServerInfo)
}

func (i *Instance) actions() *Actions {
	return i.Views()["actions"].(*Actions)
}

func (i *Instance) Stop() {
	if i.cancelFunc != nil {
		i.cancelFunc()
		i.cancelFunc = nil
	}
}

func (i *Instance) Resume() {
	// var ctx context.Context
	// ctx, i.cancelFunc = context.WithCancel(context.Background())

	//poller := NewSynsePoller(ctx, i.APIClient, refreshInterval)
	//go poller.Update()
	// TODO: Config Watcher
}

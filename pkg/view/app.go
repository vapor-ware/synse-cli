package view

import (
	"context"
	"fmt"
	"github.com/rivo/tview"
	"github.com/vapor-ware/synse-cli/pkg/client"
	"github.com/vapor-ware/synse-cli/pkg/cmd/plugin"
	"github.com/vapor-ware/synse-cli/pkg/ui"
	synse "github.com/vapor-ware/synse-server-grpc/go"
	"sort"
	"time"
)

const (
	splashDelay = 0 * time.Second
)

type Instance struct {
	*ui.App
	Content   *ui.Pages
	APIClient *client.APIClient
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

func (i *Instance) layout(_ context.Context) {
	main := tview.NewFlex().SetDirection(tview.FlexRow)

	main.AddItem(i.header(), 0, 1, false)
	//main.AddItem(tview.NewBox().SetBorder(true).SetTitle("Devices"), 0, 2, false)

	t := tview.NewTable()
	headers := []string{"ID", "Value", "Unit", "Type", "Timestamp"}
	for n, header := range headers {
		c := tview.NewTableCell(header)
		c.NotSelectable = true
		c.SetExpansion(1)
		t.SetCell(0, n, c)
		t.SetFixed(1, 0)
	}

	readings, _ := i.APIClient.Readings()
	sort.Sort(plugin.Readings(readings))
	for j, r := range readings {
		for k, _ := range headers {
			var id *tview.TableCell
			switch k {
			case 0:
				id = tview.NewTableCell(r.Id)
			case 1:
				id = tview.NewTableCell(convertToString(r))
			case 2:
				id = tview.NewTableCell(r.Unit.Symbol)
			case 3:
				id = tview.NewTableCell(r.Type)
			case 4:
				id = tview.NewTableCell(r.Timestamp)
			}
			t.SetCell(j+1, k, id)
		}
	}

	t.SetBorder(true)
	t.SetTitle(fmt.Sprintf(" Readings [%d] ", len(readings)))
	t.SetSelectable(true, false)
	t.Select(1, 0)

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

func convertToString(data interface{}) string {
	i, ok := data.(*synse.V3Reading)
	if !ok {
		return ""
	}
	var value interface{}
	switch i.Value.(type) {
	case *synse.V3Reading_StringValue:
		value = i.GetStringValue()
	case *synse.V3Reading_BoolValue:
		value = i.GetBoolValue()
	case *synse.V3Reading_Float32Value:
		value = i.GetFloat32Value()
	case *synse.V3Reading_Float64Value:
		value = i.GetFloat64Value()
	case *synse.V3Reading_Int32Value:
		value = i.GetInt32Value()
	case *synse.V3Reading_Int64Value:
		value = i.GetInt64Value()
	case *synse.V3Reading_BytesValue:
		value = i.GetBytesValue()
	case *synse.V3Reading_Uint32Value:
		value = i.GetUint32Value()
	case *synse.V3Reading_Uint64Value:
		value = i.GetUint64Value()
	}
	return fmt.Sprintf("%v", value)
}

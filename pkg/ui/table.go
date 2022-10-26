package ui

import (
	"context"
	"fmt"
	"github.com/rivo/tview"
	log "github.com/sirupsen/logrus"
	"github.com/vapor-ware/synse-cli/pkg/client"
	"github.com/vapor-ware/synse-cli/pkg/cmd/plugin"
	synse "github.com/vapor-ware/synse-server-grpc/go"
	"sort"
	"sync"
	"time"
)

type Table struct {
	*tview.Table

	refreshRate time.Duration
	mx          sync.RWMutex

	client *client.APIClient
	data   []*synse.V3Reading
}

func NewTable(c *client.APIClient) *Table {
	return &Table{
		Table:       tview.NewTable(),
		refreshRate: 2 * time.Second,
		client:      c,
	}
}

func (t *Table) Watch(ctx context.Context) error {
	go t.updater(ctx)
	return nil
}

func (t *Table) updater(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(t.refreshRate):
			err := t.update(ctx)
			if err != nil {
				log.Errorf("cannot update: %v", err)
				return
			}
		}
	}
}

func (t *Table) update(_ context.Context) error {
	t.mx.Lock()
	defer t.mx.Unlock()

	var err error
	t.data, err = t.client.Readings()
	if err != nil {
		return err
	}

	headers := []string{"ID", "Value", "Unit", "Type", "Timestamp"}
	for n, header := range headers {
		c := tview.NewTableCell(header)
		c.NotSelectable = true
		c.SetExpansion(1)
		t.SetCell(0, n, c)
		t.SetFixed(1, 0)
	}

	if t.data == nil {
		return nil
	}
	sort.Sort(plugin.Readings(t.data))
	for j, r := range t.data {
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
	t.SetTitle(fmt.Sprintf(" Readings [%d] ", len(t.data)))
	t.SetSelectable(true, false)
	t.Select(1, 0)

	return nil
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

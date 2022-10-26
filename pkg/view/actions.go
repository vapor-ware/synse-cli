package view

import (
	"fmt"
	"github.com/rivo/tview"
	"strings"
)

type Actions struct {
	*tview.Table

	instance *Instance
	data     map[string]string
}

func NewActions(i *Instance) *Actions {
	return &Actions{
		Table:    tview.NewTable(),
		instance: i,
		data:     map[string]string{},
	}
}

func (a *Actions) Init() {
	a.SetBorderPadding(0, 0, 1, 0)
	a.loadData()
	a.layout()
}

func (a *Actions) layout() {
	for i, pluginName := range strings.Split(a.data["Plugins"], ",") {
		c := tview.NewTableCell(fmt.Sprintf("<%d>  %s", i, pluginName))
		a.SetCell(i, 0, c)
	}
}

func (a *Actions) loadData() {
	plugins := make([]string, 0)
	for _, c := range a.instance.APIClient.Context() {
		plugins = append(plugins, c.Name)
	}
	a.data["Plugins"] = strings.Join(plugins, ",")
}

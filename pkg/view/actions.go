package view

import (
	"fmt"
	"github.com/rivo/tview"
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
	for i := 0; i < 3; i++ {
		c := tview.NewTableCell(fmt.Sprintf("<%d>  %s", i, "plugin"))
		a.SetCell(i, 0, c)
	}
}

func (a *Actions) loadData() {

}

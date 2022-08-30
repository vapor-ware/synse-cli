package ui

import (
	"github.com/rivo/tview"
	"sync"
	"time"
)

type Table struct {
	*tview.Table

	refreshRate time.Duration
	mx          sync.RWMutex
}

func NewTable() *Table {
	return &Table{
		Table: tview.NewTable(),
	}
}

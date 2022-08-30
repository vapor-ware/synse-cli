package ui

import "github.com/rivo/tview"

type Pages struct {
	*tview.Pages
}

func NewPages() *Pages {
	return &Pages{
		Pages: tview.NewPages(),
	}
}

func (p *Pages) Show() {
	// p.SwitchToPage()
}

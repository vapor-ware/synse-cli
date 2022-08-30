package ui

import (
	"fmt"
	"github.com/rivo/tview"
)

type Logo struct {
	*tview.Flex
	logo *tview.TextView
}

func NewLogo() *Logo {
	l := tview.NewTextView()
	l.SetWordWrap(false)
	l.SetWrap(false)
	l.SetTextAlign(tview.AlignCenter)
	l.SetDynamicColors(true)

	logo := &Logo{
		Flex: tview.NewFlex(),
		logo: l,
	}
	logo.SetDirection(tview.FlexRow)
	logo.AddItem(logo.logo, 0, 1, false)
	logo.refresh()
	return logo
}

func (l *Logo) refresh() {
	for i, s := range LogoSmall {
		fmt.Fprintf(l.logo, "[%s::b]%s", "green", s)
		if i+1 < len(LogoSmall) {
			fmt.Fprintf(l.logo, "\n")
		}
	}
}

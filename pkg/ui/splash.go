package ui

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strings"
)

var LogoSmall = []string{
	`    =+: :=-	`,
	`.:  =+: :=-  ..`,
	`**: =+: :=- .=-`,
	`**: =+: :=- .=-`,
	`**: =+: :=- .=-`,
	`**: =+: :=- .=-`,
	`+*: =+: :=- .--`,
	`    =+: :=-	`,
	`    :-. .-:	`,
}

var LogoBig = []string{
	`                                                                                     `,
	`         ::::    .:::                                                                `,
	`        .++++.   ====:                                                               `,
	`        .++++.   ====:                                                               `,
	`        .++++.   ====:                                                               `,
	`:+++-   .++++.   ====:   .---:       :---:    ::    ::   :..---.     :---:     :---: `,
	`+****   .++++.   ====:   :==--      =+=:=+=   =+:  :+=  .++=:-++    ++=-=+=   =+=-=+-`,
	`+****   .++++.   ====:   :==--      ++   ++   :+=  -+-  .++   ++.  .++   ++   ++   ++`,
	`+****   .++++.   ====:   :==--      ++   ++   .++  ++.  .++   ++.   ++   +=   ++   ++`,
	`+****   .++++.   ====:   :==--      -+=.       ++..++   .++   ++.   =+=       ++  .++`,
	`+****   .++++.   ====:   :==--       .=+=.     -+::+-   .++   ++.    :=+-.    ++=----`,
	`+****   .++++.   ====:   :==--         .=+-    .+--+.   .++   ++.      .++:   ++   ::`,
	`+****   .++++.   ====:   :==--      ==   ++     ++=+    .++   ++.   =-  .++   ++   ++`,
	`+****   .++++.   ====:   :==--      ++   ++.    =++=    .++   ++.  .+=  .++   ++   ++`,
	`+****   .++++.   ====:   :==--      ++-:-++     :++:    .++   ++.   ++::-+=   ++-:-+=`,
	`-+++=   .++++.   ====:   .---:       -===-      .++     .--   --     -==--     :==-: `,
	`        .++++.   ====:                        ..=+=                                  `,
	`        .++++.   ====:                        =+=-                                   `,
	`        .++++.   ====:                                                               `,
	`         .::.    ....                                                                `,
}

type Splash struct {
	*tview.Flex
}

func NewSplash() *Splash {
	s := Splash{
		Flex: tview.NewFlex(),
	}
	s.SetBackgroundColor(tcell.Color47)

	logo := tview.NewTextView()
	logo.SetDynamicColors(true)
	logo.SetTextAlign(tview.AlignCenter)
	s.layoutLogo(logo)

	version := tview.NewTextView()
	version.SetDynamicColors(true)
	version.SetTextAlign(tview.AlignCenter)
	s.layoutRev(version, "v4.0.0")

	s.SetDirection(tview.FlexRow)
	s.AddItem(logo, 24, 1, false)
	s.AddItem(version, 1, 1, false)
	return &s
}

func (s *Splash) layoutLogo(t *tview.TextView) {
	logo := strings.Join(LogoBig, fmt.Sprintf("\n[%s::b]", "green"))
	fmt.Fprintf(t, "%s[%s::b]%s\n",
		strings.Repeat("\n", 2),
		"green",
		logo)
}

func (s *Splash) layoutRev(t *tview.TextView, rev string) {
	fmt.Fprintf(t, "[%s::b]Revision [red::b]%s", "blue", rev)
}

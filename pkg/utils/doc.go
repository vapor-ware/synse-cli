package utils

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/gookit/color"
)

// Doc normalizes CLI documentation.
func Doc(raw string) string {
	return color.Sprint(heredoc.Doc(raw))
}

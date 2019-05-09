// Synse CLI
// Copyright (c) 2019 Vapor IO
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package utils

import (
	"io"

	"github.com/liggitt/tabwriter"
)

const (
	tabwriterFlags    = tabwriter.RememberWidths
	tabwriterMinWidth = 6
	tabwriterPadding  = 3
	tabwriterWidth    = 4
	tabwriterPadChar  = ' '
)

// NewTabWriter creates a tabwriter with default configurations to align
// input text into tab-spaced columns.
func NewTabWriter(out io.Writer) *tabwriter.Writer {
	return tabwriter.NewWriter(out, tabwriterMinWidth, tabwriterWidth, tabwriterPadding, tabwriterPadChar, tabwriterFlags)
}

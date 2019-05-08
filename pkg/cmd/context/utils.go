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

package context

import (
	"fmt"
	"io"
	"strings"

	"github.com/vapor-ware/synse-cli/pkg/config"
)

func printContextHeader(out io.Writer, full bool) error {
	columns := []string{"CURRENT", "NAME", "TYPE", "ADDRESS"}
	if !full {
		columns = columns[:3]
	}
	_, err := fmt.Fprintf(out, "%s\n", strings.Join(columns, "\t"))
	return err
}

func printContext(out io.Writer, ctx *config.ContextRecord, full bool) error {

	isCurrent := " "
	if config.IsCurrentContext(ctx) {
		isCurrent = "*"
	}

	var row string
	if full {
		row = fmt.Sprintf("%s\t%s\t%s\t%s\n", isCurrent, ctx.Name, ctx.Type, ctx.Context.Address)
	} else {
		row = fmt.Sprintf("%s\t%s\t%s\n", isCurrent, ctx.Name, ctx.Type)
	}

	_, err := fmt.Fprintf(out, row)
	return err
}

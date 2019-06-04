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
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// Printer-specific errors.
var (
	ErrNoOutputMode = errors.New("no output mode set for printer")
	ErrNoRowFunc    = errors.New("table output requires a row function")
)

// Printer defines the printing capabilities for CLI output.
type Printer struct {
	table    bool
	json     bool
	yaml     bool
	noHeader bool

	out io.Writer

	rowFunc func(data interface{}) ([]interface{}, error)
	header  []string
}

// NewPrinter creates a new printer to use for output formatting.
func NewPrinter(out io.Writer, useJSON, useYaml, noHeader bool) *Printer {
	useTable := true
	if useJSON || useYaml {
		useTable = false
	}

	return &Printer{
		table:    useTable,
		json:     useJSON,
		yaml:     useYaml,
		noHeader: noHeader,
		out:      out,
	}
}

// Write writes the data to the Printer's specified output.
func (p *Printer) Write(data interface{}) error {
	if p.table {
		return p.toTable(data)

	} else if p.json {
		return p.toJSON(data)

	} else if p.yaml {
		return p.toYAML(data)
	}

	return ErrNoOutputMode
}

// SetRowFunc sets the table row printer function, which specifies which
// data gets printed in a row of the table.
func (p *Printer) SetRowFunc(f func(data interface{}) ([]interface{}, error)) {
	p.rowFunc = f
}

// SetHeader sets the column header row for tabular formatting.
func (p *Printer) SetHeader(header ...string) {
	p.header = header
}

// toTable prints the data out in tabular format.
func (p *Printer) toTable(data interface{}) error {
	if p.rowFunc == nil {
		return ErrNoRowFunc
	}

	w := NewTabWriter(p.out)
	defer w.Flush()

	if err := p.writeHeader(w); err != nil {
		return err
	}

	var rows [][]interface{}
	switch reflect.TypeOf(data).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(data)

		for i := 0; i < s.Len(); i++ {
			row, err := p.rowFunc(s.Index(i).Interface())
			if err != nil {
				return err
			}
			rows = append(rows, row)
		}
	default:
		row, err := p.rowFunc(data)
		if err != nil {
			return err
		}
		rows = append(rows, row)
	}

	for _, row := range rows {
		fmtstr := "%v" + strings.Repeat("\t%v", len(row)-1) + "\n"
		_, err := fmt.Fprintf(w, fmtstr, row...)
		if err != nil {
			return err
		}
	}
	return nil
}

// toJSON prints the data out in JSON format.
func (p *Printer) toJSON(data interface{}) error {
	output, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	_, err = p.out.Write(append(output, '\n'))
	return err
}

// toYAML prints the data out in YAML format.
func (p *Printer) toYAML(data interface{}) error {
	output, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	_, err = p.out.Write(output)
	return err
}

// writeHeader is a helper function to write the header row out
// when headers are enabled in tabular output.
func (p *Printer) writeHeader(out io.Writer) error {
	if !p.noHeader {
		_, err := fmt.Fprintf(out, "%s\n", strings.Join(p.header, "\t"))
		return err
	}
	return nil
}

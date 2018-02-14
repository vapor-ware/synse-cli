package formatters

import (
	"fmt"
	"html/template"
	"io"
	"strings"
	"text/tabwriter"
)

const (
	// TableFormat is the key used by a Format string to denote that the
	// output should be a table.
	TableFormat = "table"

	// RawFormat is the key used by a Format string to denote that the
	// output should be in its raw form.
	RawFormat = "raw"
)

// Format is the output format for a command.
type Format string

// IsTable returns true if the format string has the "table" prefix.
func (f Format) IsTable() bool {
	return strings.HasPrefix(string(f), TableFormat)
}

// Formatter holds the context necessary to accumulate and format output
// based on a given format string.
type Formatter struct {
	Format Format
	Output io.Writer

	handler      func(interface{}) (interface{}, error)
	header       interface{}
	data         []interface{}
	parsedFormat string
}

// pre is the pre-action called before the writing of the formatted data.
func (f *Formatter) pre() {
	if f.Format.IsTable() {
		// if the format string is for a table, set the parsedFormat as the
		// format string without the 'table ' prefix (note the space included
		// after the 'table' keyword)
		f.parsedFormat = string(f.Format)[len(TableFormat)+1:]
	}
}

// SetHandler sets the Formatter handler used by the `Add` function which
// allows for command-specific formatting.
func (f *Formatter) SetHandler(fn func(interface{}) (interface{}, error)) {
	f.handler = fn
}

// SetHeader sets the header of the formatted output. The header is only
// output if the Format is a table.
func (f *Formatter) SetHeader(header interface{}) {
	f.header = header
}

// Add adds an additional piece of data to output on `Write`. For a Formatter
// processing for table, this can be thought of as adding a new row.
func (f *Formatter) Add(data interface{}) error {
	if f.handler == nil {
		return fmt.Errorf("no handler set for the formatter")
	}

	d, err := f.handler(data)
	if err != nil {
		return err
	}

	l, ok := d.([]interface{})
	if ok {
		f.data = append(f.data, l...)
	} else {
		f.data = append(f.data, d)
	}
	return nil
}

// Write writes out the data and headers (if applicable) accumulated by the
// Formatter to the output used by the CLI.
func (f *Formatter) Write() error {
	f.pre()

	if f.Format.IsTable() {
		w := tabwriter.NewWriter(f.Output, 10, 1, 3, ' ', 0)
		tmpl, err := template.New("").Parse(f.parsedFormat)
		if err != nil {
			return err
		}

		err = tmpl.Execute(w, f.header)
		if err != nil {
			return err
		}
		for _, d := range f.data {
			err := tmpl.Execute(w, d)
			if err != nil {
				return nil
			}
		}
		return w.Flush()
	}
	return nil
}

// NewFormatter creates a new instance of a Formatter.
func NewFormatter(format Format, out io.Writer) *Formatter {
	return &Formatter{
		Format: format,
		Output: out,
	}
}

package formatters

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

// Formats specifies the formats that are supported for a given formatter. The
// Pretty format should be a Go Template string that will be used to format the
// output into a table. The Yaml and JSON fields should specify a struct with
// exported fields that can be marshaled to YAML and JSON, respectively.
type Formats struct {
	// Pretty will hold the format string for pretty printing the info into a table.
	Pretty string

	// Yaml will hold the struct for the data to marshal to YAML format.
	Yaml interface{}

	// JSON will hold the struct for the data to marshal to JSON format.
	JSON interface{}
}

// Formatter holds the context necessary to accumulate and format command output.
type Formatter struct {
	Formats *Formats
	Output  io.Writer
	Context *cli.Context

	// Formatter state for pretty printing
	prettyHandler func(interface{}) (interface{}, error)
	header        interface{}
	data          []interface{}
}

// NewFormatter creates a new instance of a Formatter.
func NewFormatter(c *cli.Context, formats *Formats) *Formatter {
	return &Formatter{
		Formats: formats,
		Context: c,
		Output:  c.App.Writer,
	}
}

// SetHandler sets the Formatter handler used by the `Add` function which
// allows for command-specific "pretty" formatting.
func (f *Formatter) SetHandler(fn func(interface{}) (interface{}, error)) {
	f.prettyHandler = fn
}

// SetHeader sets the header of the formatted output. The header is only
// output if pretty formatting.
func (f *Formatter) SetHeader(header interface{}) {
	f.header = header
}

// Add adds an additional piece of data to output on `Write` for pretty
// formatting. This is equivalent to adding a row.
func (f *Formatter) Add(data interface{}) error {
	if f.prettyHandler == nil {
		return fmt.Errorf("no handler set for the formatter")
	}

	d, err := f.prettyHandler(data)
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
	format := f.getFormat()

	switch format {
	case "pretty":
		return f.writePretty()
	case "yaml", "yml":
		return f.writeYaml()
	case "json":
		return f.writeJSON()
	default:
		return fmt.Errorf("%s is an unsupported format flag. it should be one of [pretty|yaml|yml|json]", format)
	}
}

// HasPretty is a convenience function to check whether or not the given formatter
// supports "pretty" output.
func (f *Formatter) HasPretty() bool {
	return f.Formats.Pretty != ""
}

// HasYaml is a convenience function to check whether or not the given formatter
// supports "yaml" output.
func (f *Formatter) HasYaml() bool {
	return f.Formats.Yaml != nil
}

// HasJSON is a convenience function to check whether or not the given formatter
// supports "json" output.
func (f *Formatter) HasJSON() bool {
	return f.Formats.JSON != nil
}

// getFormat gets the format set by the CLI via the --format flag. By default this
// value is "pretty". If the default is used and the given formatter does not support
// a "pretty" format, it will fallback to Yaml, if it is supported.
func (f *Formatter) getFormat() string {
	var format string

	// If the --format flag is not set, it will use the default "pretty", but if
	// the formatter does not support pretty formatting, use YAML instead.
	if !f.Context.GlobalIsSet("format") && !f.HasPretty() {
		return "yaml"
	}
	format = strings.ToLower(f.Context.GlobalString("format"))
	return format
}

// writePretty writes out data in a pretty table format.
func (f *Formatter) writePretty() error {
	w := tabwriter.NewWriter(f.Output, 10, 1, 3, ' ', 0)
	tmpl, err := template.New("").Parse(f.Formats.Pretty)
	if err != nil {
		return err
	}

	if f.header != nil {
		err = tmpl.Execute(w, f.header)
		if err != nil {
			return err
		}
	}

	for _, d := range f.data {
		err := tmpl.Execute(w, d)
		if err != nil {
			return err
		}
	}
	return w.Flush()
}

// writeJSON writes out data in JSON format.
func (f *Formatter) writeJSON() error {
	o, err := json.MarshalIndent(f.Formats.JSON, "", "  ")
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	f.Output.Write(o) // nolint
	return nil
}

// writeYaml writes out data in YAML format.
func (f *Formatter) writeYaml() error {
	o, err := yaml.Marshal(f.Formats.Yaml)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	f.Output.Write(o) // nolint
	return nil
}

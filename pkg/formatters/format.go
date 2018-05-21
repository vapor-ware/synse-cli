package formatters

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"strings"
	"text/tabwriter"

	"reflect"

	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

// templateFns are custom functions that we can call in templates.
var templateFns = template.FuncMap{
	// plus1 adds one to the input value
	"plus1": func(x int) int {
		return x + 1
	},
}

// DataHandler is a function that is called when adding a row to the
// Formatter Data field. The formatter for each command should set their
// own handler, since they may all need to manipulate data a bit differently.
type DataHandler func(interface{}) (interface{}, error)

// PassthroughHandler is a built-in handler that just pass the input through
// without any modification. This can be used by various simple Formatters.
func PassthroughHandler(data interface{}) (interface{}, error) {
	return data, nil
}

// Formatter holds the context necessary to accumulate and format command output.
type Formatter struct {
	// Template is the tab-separated template that defines the output for pretty
	// formatting. If not specified for the Formatter, it will not support pretty
	// printing and will return an error when writing output.
	Template string

	// Decoder is a struct that defines how data should be output. This is used
	// for JSON and YAML output formatting. If not specified for the Formatter,
	// it will not support JSON or YAML printing and will return an error when
	// writing output. The Decoder struct can also have the tag "pretty", which
	// will specify the header info for pretty formatting. If no "pretty" tag is
	// set, it will use the field name as the header.
	Decoder interface{}

	// Output is the io.Writer that the formatter will write to when it's
	// Write() function is called.
	Output io.Writer

	// Context is the CLI context for a command.
	Context *cli.Context

	// Data is the data that will be output on Write(). Data is stored as
	// a slice to allow multiple data points to be stored (e.g. rows). If
	// multiple rows exist, all will be output (e.g. for JSON/YAML, as a
	// list). If a single value exists, it will be output individually
	// (e.g. for JSON/YAML, as an object).
	Data []interface{}

	// dataHandler is the function that is called when adding a row to the
	// Formatter Data field. The formatter for each command should set their
	// own handler, since they may all need to manipulate data a bit differently.
	dataHandler DataHandler
}

// NewFormatter creates a new instance of a Formatter.
func NewFormatter(c *cli.Context, handler DataHandler) *Formatter {
	return &Formatter{
		Context:     c,
		Output:      c.App.Writer,
		dataHandler: handler,
	}
}

// Add adds a row of data to the Formatter's Data field. This is the data that
// will be output on Write().
func (formatter *Formatter) Add(data interface{}) error {
	d, err := formatter.dataHandler(data)
	if err != nil {
		return err
	}

	dataList, ok := d.([]interface{})
	if ok {
		// If the returned data is a list, append all elements of the
		// list to the data.
		formatter.Data = append(formatter.Data, dataList...)
	} else {
		// otherwise, we just append the single instance to the data.
		formatter.Data = append(formatter.Data, d)
	}
	return nil
}

// Write writes out the data and headers (if applicable) accumulated by the
// Formatter to the output used by the CLI.
func (formatter *Formatter) Write() error {
	format := formatter.getFormat()

	switch format {
	case "pretty":
		return formatter.writePretty()
	case "yaml", "yml":
		return formatter.writeYaml()
	case "json":
		return formatter.writeJSON()
	default:
		return fmt.Errorf("%s is an unsupported format flag. it should be one of [pretty|yaml|yml|json]", format)
	}
}

// HasPretty is a convenience function to check whether or not the given formatter
// supports "pretty" output.
func (formatter *Formatter) HasPretty() bool {
	return formatter.Template != ""
}

// HasYaml is a convenience function to check whether or not the given formatter
// supports "yaml" output.
func (formatter *Formatter) HasYaml() bool {
	return formatter.Decoder != nil
}

// HasJSON is a convenience function to check whether or not the given formatter
// supports "json" output.
func (formatter *Formatter) HasJSON() bool {
	return formatter.Decoder != nil
}

// getFormat gets the format set by the CLI via the --format flag. By default this
// value is "pretty". If the default is used and the given formatter does not support
// a "pretty" format, it will fallback to YAML, if it is supported.
func (formatter *Formatter) getFormat() string {
	// If the --format flag is not set, it will use the default "pretty", but if
	// the formatter does not support pretty formatting, use YAML instead.
	if !formatter.Context.GlobalIsSet("format") && !formatter.HasPretty() {
		return "yaml"
	}
	return strings.ToLower(formatter.Context.GlobalString("format"))
}

// makeHeader populates the header info for pretty output into the formatter's
// Decoder. This does not check if pretty formatting is handled or if headers
// are enabled -- that is the responsibility of the caller.
func (formatter *Formatter) makeHeader() error {
	v := reflect.ValueOf(formatter.Decoder)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("formatter decoder must be specified as a pointer")
	}

	e := v.Elem()
	t := e.Type()
	for i := 0; i < t.NumField(); i++ {
		field := e.Field(i)
		typeField := t.Field(i)

		tag := typeField.Tag.Get("pretty")
		// If there is no "pretty" tag specified, use the field name
		if tag == "" {
			tag = typeField.Name
		}

		// A tag of "-" means do not include the field as a header.
		if tag != "-" {
			if field.IsValid() && field.CanSet() {
				if field.Kind() == reflect.String {
					field.SetString(strings.ToUpper(tag))
				}
			}
		}
	}
	return nil
}

// writePretty writes out data in a pretty table format.
func (formatter *Formatter) writePretty() error {
	if !formatter.HasPretty() {
		return fmt.Errorf("'pretty' formatting not supported for %s", formatter.Context.Command.Name)
	}

	w := tabwriter.NewWriter(formatter.Output, 10, 1, 3, ' ', 0)
	tmpl, err := template.New("").Funcs(templateFns).Parse(formatter.Template)
	if err != nil {
		return err
	}

	// todo: other PR adds in no-header option, this will be checked here.
	err = formatter.makeHeader()
	if err != nil {
		return err
	}
	err = tmpl.Execute(w, formatter.Decoder)
	if err != nil {
		return err
	}

	for _, d := range formatter.Data {
		err := tmpl.Execute(w, d)
		if err != nil {
			return err
		}
	}
	return w.Flush()
}

// writeJSON writes out data in JSON format.
func (formatter *Formatter) writeJSON() error {
	if !formatter.HasJSON() {
		return fmt.Errorf("'json' formatting not supported for %s", formatter.Context.Command.Name)
	}

	o, err := json.MarshalIndent(formatter.Data, "", "  ")
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	formatter.Output.Write(o) // nolint
	return nil
}

// writeYaml writes out data in YAML format.
func (formatter *Formatter) writeYaml() error {
	if !formatter.HasYaml() {
		return fmt.Errorf("'yaml' formatting not supported for %s", formatter.Context.Command.Name)
	}

	o, err := yaml.Marshal(formatter.Data)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	formatter.Output.Write(o) // nolint
	return nil
}

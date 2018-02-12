package formatters

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

// FormatOutput formats the given scheme struct into a supported output
// format (JSON, YAML) as determined by the value of the --format flag
// associated with the command.
func FormatOutput(c *cli.Context, out interface{}) error {
	val := c.String("output")
	switch strings.ToLower(val) {
	case "yaml", "yml", "y":
		return AsYAML(out, c.App.Writer)
	case "json", "j":
		return AsJSON(out, c.App.Writer)
	default:
		return cli.NewExitError(
			fmt.Sprintf("unsupported output flag '%s' (must be on of [y|yml|yaml|j|json])", val),
			1,
		)
	}
}

// AsYAML prints out the given interface as YAML. Here, the interfaces
// are expected to be Synse Server response schemes.
func AsYAML(out interface{}, writer io.Writer) error {
	o, err := yaml.Marshal(out)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	writer.Write(o) // nolint
	return nil
}

// AsJSON prints out the given interface as JSON. Here the interfaces
// are expected to be Synse Server response schemes.
func AsJSON(out interface{}, writer io.Writer) error {
	o, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	writer.Write(o) // nolint
	return nil
}

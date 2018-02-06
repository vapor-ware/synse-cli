package utils

import (
	"encoding/json"
	"fmt"

	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

// AsYAML prints out the given interface as YAML. Here, the interfaces
// are expected to be Synse Server response schemes.
func AsYAML(out interface{}) error {
	o, err := yaml.Marshal(out)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	fmt.Printf("%s", o)
	return nil
}

// AsJSON prints out the given interface as JSON. Here the interfaces
// are expected to be Synse Server response schemes.
func AsJSON(out interface{}) error {
	o, err := json.Marshal(out)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	fmt.Printf("%s", o)
	return nil
}

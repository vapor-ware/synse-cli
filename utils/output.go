package utils

import (
	"encoding/json"
	"fmt"

	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

// AsYAML
func AsYAML(out interface{}) error {
	o, err := yaml.Marshal(out)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	fmt.Printf("%s", o)
	return nil
}

// AsJSON
func AsJSON(out interface{}) error {
	o, err := json.Marshal(out)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	fmt.Printf("%s", o)
	return nil
}

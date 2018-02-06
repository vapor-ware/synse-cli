package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"github.com/urfave/cli"
	"encoding/json"
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

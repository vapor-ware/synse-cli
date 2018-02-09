package utils

import (
	"flag"
	"testing"

	"github.com/urfave/cli"
)

// test struct for serializing data to json/yaml
type testData struct {
	Name  string `json:"name" yaml:"name"`
	Value int    `json:"value" yaml:"value"`
}

// string flag type
type stringValue string

func (s *stringValue) Set(val string) error {
	*s = stringValue(val)
	return nil
}

func (s *stringValue) Get() interface{} { return string(*s) }

func (s *stringValue) String() string { return string(*s) }

func TestAsJSONOk(t *testing.T) {
	td := testData{
		Name:  "test",
		Value: 1,
	}
	err := AsJSON(td)
	if err != nil {
		t.Errorf("expected nil, but got: %v", err)
	}
}

func TestAsYAMLOk(t *testing.T) {
	td := testData{
		Name:  "test",
		Value: 1,
	}
	err := AsYAML(td)
	if err != nil {
		t.Errorf("expected nil, but got: %v", err)
	}
}

func TestFormatOutputJSON1(t *testing.T) {
	fs := &flag.FlagSet{}
	var v stringValue
	fs.Var(&v, "output", "")
	err := fs.Parse([]string{"--output", "json"})
	if err != nil {
		t.Error(err)
	}
	ctx := cli.NewContext(nil, fs, nil)

	td := testData{
		Name:  "test",
		Value: 1,
	}

	err = FormatOutput(ctx, td)
	if err != nil {
		t.Errorf("expected nil, but get: %v", err)
	}
}

func TestFormatOutputJSON2(t *testing.T) {
	fs := &flag.FlagSet{}
	var v stringValue
	fs.Var(&v, "output", "")
	err := fs.Parse([]string{"--output=j"})
	if err != nil {
		t.Error(err)
	}
	ctx := cli.NewContext(nil, fs, nil)

	td := testData{
		Name:  "test",
		Value: 1,
	}

	err = FormatOutput(ctx, td)
	if err != nil {
		t.Errorf("expected nil, but get: %v", err)
	}
}

func TestFormatOutputYAML1(t *testing.T) {
	fs := &flag.FlagSet{}
	var v stringValue
	fs.Var(&v, "output", "")
	err := fs.Parse([]string{"--output=yaml"})
	if err != nil {
		t.Error(err)
	}
	ctx := cli.NewContext(nil, fs, nil)

	td := testData{
		Name:  "test",
		Value: 1,
	}

	err = FormatOutput(ctx, td)
	if err != nil {
		t.Errorf("expected nil, but get: %v", err)
	}
}

func TestFormatOutputYAML2(t *testing.T) {
	fs := &flag.FlagSet{}
	var v stringValue
	fs.Var(&v, "output", "")
	err := fs.Parse([]string{"--output=yml"})
	if err != nil {
		t.Error(err)
	}
	ctx := cli.NewContext(nil, fs, nil)

	td := testData{
		Name:  "test",
		Value: 1,
	}

	err = FormatOutput(ctx, td)
	if err != nil {
		t.Errorf("expected nil, but get: %v", err)
	}
}

func TestFormatOutputYAML3(t *testing.T) {
	fs := &flag.FlagSet{}
	var v stringValue
	fs.Var(&v, "output", "")
	err := fs.Parse([]string{"--output=y"})
	if err != nil {
		t.Error(err)
	}
	ctx := cli.NewContext(nil, fs, nil)

	td := testData{
		Name:  "test",
		Value: 1,
	}

	err = FormatOutput(ctx, td)
	if err != nil {
		t.Errorf("expected nil, but get: %v", err)
	}
}

func TestFormatOutputInvalid(t *testing.T) {
	fs := &flag.FlagSet{}
	var v stringValue
	fs.Var(&v, "output", "")
	err := fs.Parse([]string{"--output=invalid"})
	if err != nil {
		t.Error(err)
	}
	ctx := cli.NewContext(nil, fs, nil)

	td := testData{
		Name:  "test",
		Value: 1,
	}

	err = FormatOutput(ctx, td)
	if err == nil {
		t.Error("expected error, but got nil")
	}
	_, ok := err.(cli.ExitCoder)
	if !ok {
		t.Error("expected error to be of fulfill ExitCoder interface but does not")
	}
}

package main

import (
	"reflect"
	"testing"

	"github.com/vapor-ware/synse-cli/commands"

	"github.com/urfave/cli"
)

func TestMain(t *testing.T) {
	app := cli.NewApp()
	app.Name = "synse"
	app.Usage = "Destroy The World"
	app.Version = "97"
	app.Authors = []cli.Author{{Name: "Tiny Rick", Email: "rick@tiny.com"}}

	app.Commands = commands.Commands

	if reflect.TypeOf(app) != reflect.TypeOf(cli.NewApp()) {
		t.Errorf("App is type %s, not %s", reflect.TypeOf(app), reflect.TypeOf(cli.NewApp()))
	}
}

package test

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/config"
)


func Setup() {
	config.Config = config.CliConfig{
		Debug: false,
		ActiveHost: nil,
		Hosts: make(map[string]*config.HostConfig, 0),
	}
}


func NewFakeApp() *cli.App {
	app := &cli.App{
		Name: "test app",
	}

	// prevent the fake app from calling os.Exit() on failure
	cli.OsExiter = func(code int) {}
	return app
}
package commands

import (
	"testing"

	"github.com/urfave/cli"
)

func constructTestCli() *cli.App {
	app := cli.NewApp()
	app.Name = "vesh"
	app.Usage = "Destory The World"
	app.Version = "97"
	app.Authors = []cli.Author{{Name: "Tiny Rick", Email: "rick@tiny.com"}}

	app.Commands = Commands

	return app
}

func TestCommands_args(t *testing.T) {
	tests := []struct {
		purpose       string
		command       string
		args          []string
		expectedError bool
	}{
		{
			purpose:       "no args",
			args:          []string{""},
			expectedError: true,
		},
		{
			purpose:       "too many args",
			args:          make([]string, 10),
			expectedError: true,
		},
		// {
		// 	purpose:       "just right",
		// 	args:          []string{"", ""},
		// 	expectedError: false,
		// },
	}

	cli := constructTestCli()
	commands := cli.Commands

	for _, cmd := range commands {
		for _, condition := range tests {
			t.Log(cmd.FullName(), condition)
			switch err := cli.Run(condition.args); err != nil {
			case condition.expectedError:
				t.Errorf("Command %s failed while trying %s with error: %v", cmd.Name, condition.purpose, err)
			}
		}
	}
}

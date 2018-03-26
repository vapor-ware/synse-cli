package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

var (
	autocompletePath = "/etc/bash_completion.d"

	// bash
	bashProfile = "$HOME/.bash_profile"
	bashrc      = "$HOME/.bashrc"

	// zsh
	zshrc = "$HOME/.zshrc"
)

// bashCompletion is the bash completion function for Synse. This is taken
// from: https://github.com/urfave/cli/blob/master/autocomplete/bash_autocomplete
const bashCompletion = `
_cli_bash_autocomplete() {
    local cur opts base
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    opts=$( ${COMP_WORDS[@]:0:$COMP_CWORD} --generate-bash-completion )
    COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )
    return 0
}

complete -F _cli_bash_autocomplete synse
`

// zshCompletion is the zsh completion function for Synse. This is taken
// from: https://github.com/urfave/cli/blob/master/autocomplete/zsh_autocomplete
const zshCompletion = `
_cli_zsh_autocomplete() {

  local -a opts
  opts=("${(@f)$(_CLI_ZSH_AUTOCOMPLETE_HACK=1 ${words[@]:0:#words[@]-1} --generate-bash-completion)}")

  _describe 'values' opts

  return
}

compdef _cli_zsh_autocomplete synse
`

// completionCommand is the CLI command for generating shell completion scripts.
var completionCommand = cli.Command{
	Name:  "completion",
	Usage: "Generate shell completion scripts for bash or zsh",

	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdCompletion(c))
	},

	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:   "bash-completion",
			Usage:  "output bash completion (to be eval'd)",
			Hidden: true,
		},
		cli.BoolFlag{
			Name:   "zsh-completion",
			Usage:  "output zsh completion (to be eval'd)",
			Hidden: true,
		},
	},

	Subcommands: []cli.Command{
		{
			Name:  "bash",
			Usage: "enable bash completion",

			Action: func(c *cli.Context) error {
				return utils.CmdHandler(cmdEnable(c, "bash"))
			},
		},
		{
			Name:  "zsh",
			Usage: "enable zsh completion",

			Action: func(c *cli.Context) error {
				return utils.CmdHandler(cmdEnable(c, "zsh"))
			},
		},
	},
}

// cmdCompletion is the action for the completionCommand.
func cmdCompletion(c *cli.Context) error {
	switch {
	case c.IsSet("bash-completion") && c.IsSet("zsh-completion"):
		return fmt.Errorf("cannot specify both 'bash' and 'zsh' at once")
	case c.IsSet("bash-completion"):
		_, err := c.App.Writer.Write([]byte(bashCompletion))
		return err
	case c.IsSet("zsh-completion"):
		_, err := c.App.Writer.Write([]byte(zshCompletion))
		return err
	default:
		return cli.ShowSubcommandHelp(c)
	}
}

// cmdEnable is the action for enabling a shell completion.
func cmdEnable(c *cli.Context, shell string) (err error) {

	switch o := runtime.GOOS; o {
	case "darwin":
		return darwinCompletion(c, shell)
	case "linux":
		return linuxCompletion(c, shell)
	default:
		return fmt.Errorf("synse CLI does not yet support %s systems", o)
	}
}

// getFileOrDir returns the expanded path for the first path (file or
// dir) found to exist out of the ones passed in as parameters. If none
// of the given paths exist, an error is returned.
func getFileOrDir(paths ...string) (string, error) {
	for _, p := range paths {
		exp := os.ExpandEnv(p)
		_, err := os.Stat(exp)
		// if the path exists, return its expanded name
		if err == nil {
			return exp, nil
		}
	}
	return "", fmt.Errorf("none of the expected files/dirs were found: %v", paths)
}

// darwinCompletion sets up Synse CLI completion on darwin systems.
func darwinCompletion(c *cli.Context, shell string) (err error) {
	var file string
	switch shell {
	case "bash":
		file, err = getFileOrDir(bashProfile, bashrc)
		if err != nil {
			return err
		}
	case "zsh":
		file, err = getFileOrDir(zshrc)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported shell specified: %v", shell)
	}

	// the line that will be added to the file
	complete := `
# synse completion
eval '$(synse completion --%s-completion)'
`

	// append the completion to the rc file
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close() // nolint

	_, err = f.Write([]byte(fmt.Sprintf(complete, shell)))
	if err != nil {
		return err
	}

	_, err = c.App.Writer.Write([]byte(
		fmt.Sprintf("You can now run `source %s`\n", file)),
	)
	return
}

// linuxCompletion sets up Synse CLI completion on linux systems.
func linuxCompletion(c *cli.Context, shell string) (err error) {
	var file string
	var contents string
	switch shell {
	case "bash":
		file, err = getFileOrDir(autocompletePath)
		if err != nil {
			return err
		}
		file = filepath.Join(file, "synse") // /etc/bash_completion.d/synse
		contents = bashCompletion
	case "zsh":
		// FIXME (etd) - I don't know how zsh is setup on linux, so for now
		// will not do this.
		return fmt.Errorf("zsh completion is not yet supported on linux")
	default:
		return fmt.Errorf("unsupported shell specified: %v", shell)
	}

	// Create the completion file
	out, err := os.Create(file)
	if err != nil {
		return err
	}
	defer out.Close() // nolint

	_, err = out.Write([]byte(contents))
	if err != nil {
		return err
	}

	_, err = c.App.Writer.Write([]byte(
		fmt.Sprintf("You can now run `source %s`\n or reload your shell", file)),
	)
	return err
}

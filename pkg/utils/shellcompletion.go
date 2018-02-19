package utils

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

var autocompletionPath = "/etc/bash_completion.d/"

const autocompleteURL = "https://raw.githubusercontent.com/urfave/cli/master/autocomplete/"

// GenerateShellCompletion generates the correct completion file for `bash` and
// `zsh` shells. These files can then be sources to provide command and flag
// autocompletion.
func GenerateShellCompletion(c *cli.Context, shell string) error {
	var err error
	var path string
	switch shell {
	case "bash":
		path, err = downloadCompletionFile(c, "bash")
	case "zsh":
		path, err = downloadCompletionFile(c, "zsh")
	}
	if err != nil {
		return err
	}
	fmt.Printf("You can now run `%s %s`\n", "source", path)
	return nil
}

// downloadCompletionFile downloads the required bash completion file from
// https://github.com/urfave/cli .
func downloadCompletionFile(c *cli.Context, shell string) (string, error) {
	var err error
	var dirOut string

	// If the --path or -p flag is set for the completion command, use the
	// value provided for the output directory, otherwise use the default
	// output directory.
	if c.IsSet("path") {
		dirOut = c.String("path")
	} else {
		dirOut = autocompletionPath
	}

	// Check if the directory exists - if it does not, try to create it.
	path := filepath.Join(dirOut, c.App.Name)
	_, err = os.Stat(dirOut)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dirOut, os.ModePerm)
		if err != nil {
			return "", fmt.Errorf("unable to create %s (try running with sudo)", dirOut)
		}
	}

	// Create the completion file
	out, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer out.Close() // nolint

	// Download the completion file from GitHub and copy it into the created
	// completion file locally.
	shellPath := autocompleteURL + shell + "_autocomplete"
	resp, err := http.Get(shellPath)
	if err != nil {
		return path, err
	}
	defer resp.Body.Close() // nolint

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	inFile, err := ioutil.ReadFile(path) // FIXME: There's a better way to do this
	if err != nil {
		return "", err
	}
	output := bytes.Replace(inFile, []byte("$(basename ${BASH_SOURCE})"), []byte("synse"), -1)

	err = ioutil.WriteFile(path, output, 0666)
	return path, err
}

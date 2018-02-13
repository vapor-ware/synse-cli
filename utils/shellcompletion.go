package utils

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
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
		path, err = downloadCompletionFile(c.App.Name, "bash")
	case "zsh":
		path, err = downloadCompletionFile(c.App.Name, "zsh")
	}
	if err != nil {
		return err
	}
	fmt.Printf("You can now run `%s %s`\n", "source", path)
	return nil
}

// downloadCompletionFile downloads the required bash completion file from
// https://github.com/urfave/cli .
func downloadCompletionFile(appName, shell string) (string, error) {
	var err error
	path := autocompletionPath + appName

	_, err = os.Stat(autocompletionPath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(autocompletionPath, os.ModePerm)
		if err != nil {
			return "", fmt.Errorf("unable to create %s (try running with sudo)", autocompletionPath)
		}
	}

	out, err := os.Create(path)
	if err != nil {
		return path, err
	}
	defer out.Close() // nolint

	shellPath := autocompleteURL + shell + "_autocomplete"
	resp, err := http.Get(shellPath)
	if err != nil {
		return path, err
	}
	defer resp.Body.Close() // nolint

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Debug(err)
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

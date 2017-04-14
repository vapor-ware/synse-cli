package utils

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

var appName = "vesh" // TODO: Dynamically source this from app.Name
var autocompletionPath = "/etc/bash_completion.d/"

const autocompleteUrl = "https://raw.githubusercontent.com/urfave/cli/master/autocomplete/"

// GenerateShellCompletion generates the correct completion file for `bash` and
// `zsh` shells. These files can then be sources to provide command and flag
// autocompletion.
func GenerateShellCompletion(shell string) error {
	var err error
	var path string
	switch shell {
	case "bash":
		path, err = downloadCompletionFile("bash")
	case "zsh":
		path, err = downloadCompletionFile("zsh")
	}
	fmt.Printf("You can now run `%s %s`\n", "source", path)
	return err
}

// downloadCompletionFile downloads the required bash completion file from
// https://github.com/urfave/cli .
func downloadCompletionFile(shell string) (string, error) {
	var err error
	path := autocompletionPath + appName
	out, err := os.Create(path)
	defer out.Close()
	switch {
	case os.IsExist(err):
		return path, err
	case os.IsPermission(err):
		return path, err
	}
	shellPath := autocompleteUrl + shell + "_autocomplete"
	resp, err := http.Get(shellPath)
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	inFile, err := ioutil.ReadFile(path) // FIXME: There's a better way to do this
	output := bytes.Replace(inFile, []byte("$(basename ${BASH_SOURCE})"), []byte("vesh"), -1)
	err = ioutil.WriteFile(path, output, 0666)
	// fmt.Println(shellPath, resp, err, inFile, output)
	return path, err
}

package utils

import (
  "os"
  "io"
  "net/http"
)

var appName = "vesh" // TODO: Dynamically source this from app.Name
var autocompletionPath = "/etc/bash_completion.d/"
const autocompleteUrl = "https://raw.githubusercontent.com/urfave/cli/master/autocomplete/"

func GenerateShellCompletion(shell string) error {
  var err error
  switch shell {
  case "bash":
    err = bashcompletion()
  case "zsh":
    err = zshcompletion()
  }
  return err
}

func downloadCompletionFile(shell string) error {
  var err error
  path := autocompletionPath + appName
  out, err := os.Create(path)
  defer out.Close()
  switch err {
  case os.IsExist(err):
    return err
  case os.IsPermission(err):
    return err
  }
  shellPath := autocompleteUrl + shell + "_autocomlete"
  resp, err := http.Get(shellPath)
  defer resp.Body.Close()
  out, err := io.Copy(out, resp.Body)
  inFile, err := io.ioutil.ReadFile(path) // FIXME: There's a more compact way to do this
  
  return err
}

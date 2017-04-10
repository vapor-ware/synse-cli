package client

import (
  "testing"
  //"net/http"
  "strings"

  //"github.com/dghubble/sling"
)

func TestConstructUrl(t *testing.T) {
  var host = "demo.vapor.io"
  url := constructUrl(host)
  switch {
  case !strings.HasPrefix(url, "http://"):
    t.Error(url, "No http:// detected")
  case !strings.HasSuffix(url, "/"):
    t.Error(url, "URL does not end with slash")
  case !strings.Contains(url, "opendcre"):
    t.Error(url, "URL does not contain 'opendcre'")
  }
  parts := strings.Split(strings.TrimPrefix(url, "http://"), "/")
  if len(parts) <= 1 {
    t.Error(url, "URL does not appear to contain all the required parts")
  }
}
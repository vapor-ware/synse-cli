package commands

import (
	"testing"

	"github.com/vapor-ware/synse-cli/client"
)

func init() {
	client.Config("35.185.201.100")
}

func TestTestAPI(t *testing.T) {
	err := TestAPI()
	if err != nil {
		t.Error(err)
	}
}

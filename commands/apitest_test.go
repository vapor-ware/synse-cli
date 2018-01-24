package commands

import (
	"os"
	"testing"

	"github.com/vapor-ware/synse-cli/client"
)

func init() {
	client.Config(os.Getenv("SYNSE_HOST"))
}

func TestTestAPI(t *testing.T) {
	err := TestAPI()
	if err != nil {
		t.Error(err)
	}
}

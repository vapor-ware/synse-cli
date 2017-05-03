package commands

import (
	"fmt"
	"net/http"

	"github.com/vapor-ware/synse-cli/client"
)

// testpath specifies the url endpoint to test against
const testpath = "test"

// APIStatus contains the response object from the test endpoint.
type APIStatus struct {
	Status string `json:"status"`
}

// TestAPI checks the "../<testpath>" endpoint and returns the status returned.
func TestAPI() error { // This should be supressed when not directly called unless debug is set
	status := &APIStatus{}
	resp, err := client.New().Get(testpath).ReceiveSuccess(status)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return err
	}
	fmt.Println("API reported status ok")
	return nil
}

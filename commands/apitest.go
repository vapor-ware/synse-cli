package commands

import (
	"fmt"
	"net/http"

	"github.com/vapor-ware/vesh/client"
)

const testpath = "test"

type APIStatus struct {
	Status string `json:"status"`
}

// TestAPI checks the "../<testpath>" endpoint and returns the status returned.
func TestAPI(vc *client.VeshClient) error { // This should be supressed unless debug is set
	status := &APIStatus{}
	resp, err := vc.Sling.New().Get(testpath).ReceiveSuccess(status)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return err
	}
	fmt.Println("API reported status ok")
	return nil
}

package client

import (
	//"os"
	"fmt"
	"strconv"

	"github.com/dghubble/sling"
)

var VeshHostPtr = ""

func constructUrl() string {
	var vaporBase = fmt.Sprint(VeshHostPtr)
	var vaporPort = strconv.Itoa(5000)
	var defaultPath = "opendcre/1.3/" //Add a version number here
	var CompleteBase = fmt.Sprintf("http://%s:%s/%s", vaporBase, vaporPort, defaultPath)
	//fmt.Println(CompleteBase) For bug testing only
	return CompleteBase
}

type VeshClient struct {
	Sling *sling.Sling
}

func New() *VeshClient {
	cb := constructUrl()
	return &VeshClient{
		Sling: sling.New().Base(cb),
	}
}

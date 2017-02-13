package main

import (
  "os"
  "fmt"
  "strconv"
  "net/http"

  "github.com/urfave/cli"
  "github.com/dghubble/sling"
)

func constructUrl() string {
  var vaporBase = fmt.Sprint(os.Getenv("VESH_HOST"))
  var vaporPort = strconv.Itoa(5000)
  var defaultPath = "opendcre/1.3/" //Add a version number here
  var CompleteBase = fmt.Sprintf("http://%s:%i/%s", vaporBase, vaporPort, defaultPath)
  fmt.Println(CompleteBase)
  return CompleteBase
}

type VeshClient struct {
  sling *sling.Sling
}

func NewVeshClient(client *http.Client) *VeshClient {
  cb := constructUrl()
  return &VeshClient{
    sling: sling.New().Client(client).Base(cb),
  }
}

func CliScan(c *cli.Context)  {
  fmt.Println(constructUrl)
  client := new(VeshClient)
  client.cmdScan()
  fmt.Println("nope")
}

func (client *VeshClient) cmdScan()   {
  //client := new(VeshClient)
  resp, err := client.sling.New().Get("scan").Request()
  fmt.Println(resp, err)
}

func test()  {

}

func cmdListTemp(c *cli.Context)  {

}
func cmdGetTemp(c *cli.Context)  {

}
func cmdSetBootTarget(c *cli.Context)  {

}
func cmdGetBootTarget(c *cli.Context)  {

}
func cmdListLed(c *cli.Context)  {

}
func cmdGetLed(c *cli.Context)  {

}
func cmdSetLed(c *cli.Context)  {

}
func cmdBlinkled(c *cli.Context)  {

}
func cmdColorLed(c *cli.Context)  {

}
func cmdMapLocation(c *cli.Context)  {

}
func cmdFindDevice(c *cli.Context)  {

}
func cmdShowPowerLoad(c *cli.Context)  {

}
func cmdShowMemoryLoad(c *cli.Context)  {

}
func cmdShowTempratureLoad(c *cli.Context)  {

}
func cmdShowCPULoad(c *cli.Context)  {

}
func cmdShowApplicationLoad(c *cli.Context)  {

}
func cmdProvisionNew(c *cli.Context)  {

}
func cmdDeprovision(c *cli.Context)  {

}
func cmdProvisionList(c *cli.Context)  {

}
func cmdListHostname(c *cli.Context)  {

}
func cmdGetHostname(c *cli.Context)  {

}
func cmdListPower(c *cli.Context)  {

}
func cmdGetPower(c *cli.Context)  {

}
func cmdSetPower(c *cli.Context)  {

}
func cmdListFan(c *cli.Context)  {

}
func cmdGetFan(c *cli.Context)  {

}

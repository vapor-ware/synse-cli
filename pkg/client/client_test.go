package client

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/urfave/cli"

	"github.com/vapor-ware/synse-cli/internal/test"
	"github.com/vapor-ware/synse-cli/pkg/scheme"
)

func TestCheck1(t *testing.T) {
	var err error
	resp := http.Response{
		StatusCode: 200,
	}

	errScheme := new(scheme.Error)
	e := check(&resp, err, errScheme)
	test.ExpectNoError(t, e)
}

func TestCheck2(t *testing.T) {
	err := fmt.Errorf("test error")
	resp := http.Response{
		StatusCode: 200,
	}

	errScheme := new(scheme.Error)
	e := check(&resp, err, errScheme)
	test.ExpectExitCoderError(t, e)
}

func TestCheck3(t *testing.T) {
	var err error
	resp := http.Response{
		StatusCode: 500,
	}

	errScheme := new(scheme.Error)
	e := check(&resp, err, errScheme)
	test.ExpectExitCoderError(t, e)
}

func TestMakeURI(t *testing.T) {
	var testTable = []struct {
		in  []string
		out string
	}{
		{
			in:  []string{""},
			out: "",
		},
		{
			in:  []string{"foo"},
			out: "foo",
		},
		{
			in:  []string{"foo", "bar"},
			out: "foo/bar",
		},
		{
			in:  []string{"foo", "bar", "baz/"},
			out: "foo/bar/baz/",
		},
		{
			in:  []string{"foo/", "/bar"},
			out: "foo///bar",
		},
	}

	cli.OsExiter = func(int) {}

	for _, testCase := range testTable {
		r := makeURI(testCase.in...)
		if r != testCase.out {
			t.Errorf("makeURI(%v) => %s, want %s", testCase.in, r, testCase.out)
		}
	}
}

func TestGetVersioned(t *testing.T) {
	test.Setup()
	cli.OsExiter = func(int) {}

	type testData struct {
		Status string
	}

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/foobar", 200, `{"status": "ok"}`)

	test.AddServerHost(server)

	out := &testData{}
	err := getVersioned("foobar", out)

	test.ExpectNoError(t, err)

	if out.Status != "ok" {
		t.Errorf("expected status 'ok', but was %s", out.Status)
	}
}

func TestGetUnversioned(t *testing.T) {
	test.Setup()
	cli.OsExiter = func(int) {}

	type testData struct {
		Status string
	}

	mux, server := test.UnversionedServer()
	defer server.Close()

	test.Serve(t, mux, "/synse/foobar", 200, `{"status": "ok"}`)

	test.AddServerHost(server)

	out := &testData{}
	err := getUnversioned("foobar", out)

	test.ExpectNoError(t, err)

	if out.Status != "ok" {
		t.Errorf("expected status 'ok', but was %s", out.Status)
	}
}

func TestPostVersioned(t *testing.T) {
	test.Setup()
	cli.OsExiter = func(int) {}

	type testData struct {
		Status string
	}

	mux, server := test.Server()
	defer server.Close()

	test.Serve(t, mux, "/synse/v2/foobar", 200, `{"status": "ok"}`)

	test.AddServerHost(server)

	out := &testData{}
	err := postVersioned("foobar", testData{}, out)

	test.ExpectNoError(t, err)

	if out.Status != "ok" {
		t.Errorf("expected status 'ok', but was %s", out.Status)
	}
}

func TestStatus(t *testing.T) {
	test.Setup()

	client := &synseClient{}
	mux, server := test.Server()
	defer server.Close()

	in := `
{
  "status": "ok",
  "timestamp": "2018-06-28T12:59:47.625842798Z"
}`

	out := &scheme.TestStatus{
		Status:    "ok",
		Timestamp: "2018-06-28T12:59:47.625842798Z",
	}

	test.Serve(t, mux, "/synse/test", 200, in)

	test.AddServerHost(server)

	resp, err := client.Status()
	test.ExpectNoError(t, err)
	test.AssertEqual(t, out, resp)
}

func TestConfig(t *testing.T) {
	test.Setup()

	client := &synseClient{}
	mux, server := test.Server()
	defer server.Close()

	in := `
{
  "locale":"en_US",
  "pretty_json":true,
  "logging":"debug"
}`

	out := &scheme.Config{
		Locale:     "en_US",
		PrettyJSON: true,
		Logging:    "debug",
	}

	test.Serve(t, mux, "/synse/v2/config", 200, in)

	test.AddServerHost(server)

	resp, err := client.Config()
	test.ExpectNoError(t, err)
	test.AssertEqual(t, out, resp)
}

func TestCapabilities(t *testing.T) {
	test.Setup()

	client := &synseClient{}
	mux, server := test.Server()
	defer server.Close()

	in := `
[
  {
    "plugin":"vaporio\/emulator-plugin",
    "devices":[
      {
        "kind":"led",
        "outputs":[
          "led.color",
          "led.state"
        ]
      }
    ]
  }
]`

	out := []scheme.Capability{
		scheme.Capability{
			Plugin: "vaporio/emulator-plugin",
			Devices: []scheme.CapabilityDevice{
				scheme.CapabilityDevice{
					Kind: "led",
					Outputs: []string{
						"led.color",
						"led.state",
					},
				},
			},
		},
	}

	test.Serve(t, mux, "/synse/v2/capabilities", 200, in)

	test.AddServerHost(server)

	resp, err := client.Capabilities()
	test.ExpectNoError(t, err)
	test.AssertEqual(t, out, resp)
}

func TestPlugins(t *testing.T) {
	test.Setup()

	client := &synseClient{}
	mux, server := test.Server()
	defer server.Close()

	in := `
[
  {
    "tag":"vaporio\/emulator-plugin",
    "name":"sample-tcp",
    "description":"A sample emulator plugin",
    "maintainer":"vaporio",
    "vcs":"github.com\/vapor-ware\/synse-emulator-plugin",
    "version":{
      "plugin_version":"2.0.0",
      "sdk_version":"1.0.0",
      "build_date":"2018-06-25T14:39:18",
      "git_commit":"4831f12",
      "git_tag":"1.0.2-8-g4831f12",
      "arch":"amd64",
      "os":"linux"
    },
    "network":{
      "protocol":"tcp",
      "address":"emulator-plugin:5001"
    },
    "health":{
      "timestamp":"2018-06-27T18:30:46.237254715Z",
      "status":"ok",
      "message":"",
      "checks":[
        {
          "name":"read buffer health",
          "status":"ok",
          "message":"",
          "timestamp":"2018-06-27T18:30:16.531781924Z",
          "type":"periodic"
        }
      ]
    }
  }
]`

	out := []scheme.Plugin{
		scheme.Plugin{
			Tag:         "vaporio/emulator-plugin",
			Name:        "sample-tcp",
			Description: "A sample emulator plugin",
			Maintainer:  "vaporio",
			VCS:         "github.com/vapor-ware/synse-emulator-plugin",
			Version: scheme.VersionData{
				Version:    "2.0.0",
				SDKVersion: "1.0.0",
				BuildDate:  "2018-06-25T14:39:18",
				GitCommit:  "4831f12",
				GitTag:     "1.0.2-8-g4831f12",
				Arch:       "amd64",
				OS:         "linux",
			},
			Network: scheme.NetworkData{
				Protocol: "tcp",
				Address:  "emulator-plugin:5001",
			},
			Health: scheme.HealthData{
				Timestamp: "2018-06-27T18:30:46.237254715Z",
				Status:    "ok",
				Message:   "",
				Checks: []scheme.CheckData{
					scheme.CheckData{
						Name:      "read buffer health",
						Status:    "ok",
						Message:   "",
						Timestamp: "2018-06-27T18:30:16.531781924Z",
						Type:      "periodic",
					},
				},
			},
		},
	}

	test.Serve(t, mux, "/synse/v2/plugins", 200, in)

	test.AddServerHost(server)

	resp, err := client.Plugins()
	test.ExpectNoError(t, err)
	test.AssertEqual(t, out, resp)
}

// FIXME: The current definition scheme of ScanResponse is not consistent with
// other schemes. This also makes it harder to test. Need to decouple it first.
func TestScan(t *testing.T) {}

func TestRackInfo(t *testing.T) {
	test.Setup()

	client := &synseClient{}
	mux, server := test.Server()
	defer server.Close()

	in := `
{
  "rack":"rack-1",
  "boards":[
    "board-1"
  ]
}`

	out := &scheme.RackInfo{
		Rack: "rack-1",
		Boards: []string{
			"board-1",
		},
	}

	test.Serve(t, mux, "/synse/v2/info/rack-1", 200, in)

	test.AddServerHost(server)

	resp, err := client.RackInfo("rack-1")
	test.ExpectNoError(t, err)
	test.AssertEqual(t, out, resp)
}

func TestBoardInfo(t *testing.T) {
	test.Setup()

	client := &synseClient{}
	mux, server := test.Server()
	defer server.Close()

	in := `
{
  "board":"board-1",
  "location":{
    "rack":"rack-1"
  },
  "devices":[
    "device-1",
    "device-2",
    "device-3"
  ]
}`
	out := &scheme.BoardInfo{
		Board: "board-1",
		Location: map[string]string{
			"rack": "rack-1",
		},
		Devices: []string{
			"device-1",
			"device-2",
			"device-3",
		},
	}

	test.Serve(t, mux, "/synse/v2/info/rack-1/board-1", 200, in)

	test.AddServerHost(server)

	resp, err := client.BoardInfo("rack-1", "board-1")
	test.ExpectNoError(t, err)
	test.AssertEqual(t, out, resp)
}

func TestDeviceInfo(t *testing.T) {
	test.Setup()

	client := &synseClient{}
	mux, server := test.Server()
	defer server.Close()

	in := `
{
  "timestamp":"2018-06-28T12:59:47.625842798Z",
  "uid":"device-1",
  "kind":"pressure",
  "metadata":{
    "model":"emul8-pressure"
  },
  "plugin":"emulator-plugin",
  "info":"Synse Pressure Sensor 1",
  "location":{
    "rack":"rack-1",
    "board":"board-1"
  },
  "output":[
    {
      "name":"pressure",
      "type":"pressure",
      "precision":3,
      "scaling_factor":1.5,
      "unit":{
        "name":"pascals",
        "symbol":"Pa"
      }
    }
  ]
}`
	out := &scheme.DeviceInfo{
		Timestamp: "2018-06-28T12:59:47.625842798Z",
		UID:       "device-1",
		Kind:      "pressure",
		Metadata: map[string]string{
			"model": "emul8-pressure",
		},
		Plugin: "emulator-plugin",
		Info:   "Synse Pressure Sensor 1",
		Location: map[string]string{
			"rack":  "rack-1",
			"board": "board-1",
		},
		Output: []scheme.DeviceOutput{
			scheme.DeviceOutput{
				Name:          "pressure",
				Type:          "pressure",
				Precision:     int(3),
				ScalingFactor: float64(1.5),
				Unit: scheme.OutputUnit{
					Name:   "pascals",
					Symbol: "Pa",
				},
			},
		},
	}

	test.Serve(t, mux, "/synse/v2/info/rack-1/board-1/device-1", 200, in)

	test.AddServerHost(server)

	resp, err := client.DeviceInfo("rack-1", "board-1", "device-1")
	test.ExpectNoError(t, err)
	test.AssertEqual(t, out, resp)
}

func TestRead(t *testing.T) {
	test.Setup()

	client := &synseClient{}
	mux, server := test.Server()
	defer server.Close()

	in := `
{
  "kind":"temperature",
  "data":[
    {
      "value":"65",
      "timestamp":"2018-06-28T12:41:50.333443322Z",
      "unit":{
        "symbol":"C",
        "name":"celsius"
      },
      "type":"temperature",
      "info":"mock temperature response"
    }
  ]
}`

	out := &scheme.Read{
		Kind: "temperature",
		Data: []scheme.ReadData{
			scheme.ReadData{
				Value:     "65",
				Timestamp: "2018-06-28T12:41:50.333443322Z",
				Unit: scheme.OutputUnit{
					Symbol: "C",
					Name:   "celsius",
				},
				Type: "temperature",
				Info: "mock temperature response",
			},
		},
	}

	test.Serve(t, mux, "/synse/v2/read/rack-1/board-1/device-1", 200, in)

	test.AddServerHost(server)

	resp, err := client.Read("rack-1", "board-1", "device-1")
	test.ExpectNoError(t, err)
	test.AssertEqual(t, out, resp)
}

func TestTransaction(t *testing.T) {
	test.Setup()

	client := &synseClient{}
	mux, server := test.Server()
	defer server.Close()

	in := `
{
  "id":"b9u6ss6q5i6g020lau6g",
  "context":{
    "action":"color",
    "data":"000000"
  },
  "state":"ok",
  "status":"done",
  "created":"2018-06-28T12:59:47.625842798Z",
  "updated":"2018-06-28T12:59:47.625842798Z",
  "message":""
}`

	out := &scheme.Transaction{
		ID: "b9u6ss6q5i6g020lau6g",
		Context: scheme.WriteContext{
			Action: "color",
			Data:   "000000",
		},
		State:   "ok",
		Status:  "done",
		Created: "2018-06-28T12:59:47.625842798Z",
		Updated: "2018-06-28T12:59:47.625842798Z",
		Message: "",
	}

	test.Serve(t, mux, "/synse/v2/transaction/b9u6ss6q5i6g020lau6g", 200, in)

	test.AddServerHost(server)

	resp, err := client.Transaction("b9u6ss6q5i6g020lau6g")
	test.ExpectNoError(t, err)
	test.AssertEqual(t, out, resp)
}

func TestTransactionList(t *testing.T) {
	test.Setup()

	client := &synseClient{}
	mux, server := test.Server()
	defer server.Close()

	in := `
[
  "b9u6ss6q5i6g020lau6g"
]`

	out := &[]string{
		"b9u6ss6q5i6g020lau6g",
	}

	test.Serve(t, mux, "/synse/v2/transaction", 200, in)

	test.AddServerHost(server)

	resp, err := client.TransactionList()
	test.ExpectNoError(t, err)
	test.AssertEqual(t, out, resp)
}

func TestWrite(t *testing.T) {
	test.Setup()

	client := &synseClient{}
	mux, server := test.Server()
	defer server.Close()

	in := `
[
  {
    "context":{
      "action":"color",
      "data":"000000"
    },
    "transaction":"b9u6ss6q5i6g020lau6g"
  }
]`

	out := []scheme.WriteTransaction{
		scheme.WriteTransaction{
			Context: scheme.WriteContext{
				Action: "color",
				Data:   "000000",
			},
			Transaction: "b9u6ss6q5i6g020lau6g",
		},
	}

	test.Serve(t, mux, "/synse/v2/write/rack-1/board-1/device-1", 200, in)

	test.AddServerHost(server)

	resp, err := client.Write("rack-1", "board-1", "device-1", "color", "000000")
	test.ExpectNoError(t, err)
	test.AssertEqual(t, out, resp)
}

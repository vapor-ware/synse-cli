package client

import (
	"fmt"
	"net/http"
	"reflect"
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

	test.Serve(t, mux, "/synse/2.0/foobar", 200, `{"status": "ok"}`)

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

	test.Serve(t, mux, "/synse/2.0/foobar", 200, `{"status": "ok"}`)

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

	if !reflect.DeepEqual(out, resp) {
		t.Errorf("expected %+v, but was %+v", out, resp)
	}
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

	test.Serve(t, mux, "/synse/2.0/config", 200, in)

	test.AddServerHost(server)

	resp, err := client.Config()
	test.ExpectNoError(t, err)

	if !reflect.DeepEqual(out, resp) {
		t.Errorf("expected %+v, but was %+v", out, resp)
	}
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

	test.Serve(t, mux, "/synse/2.0/capabilities", 200, in)

	test.AddServerHost(server)

	resp, err := client.Capabilities()
	test.ExpectNoError(t, err)

	if !reflect.DeepEqual(out, resp) {
		t.Errorf("expected %+v, but was %+v", out, resp)
	}
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

	test.Serve(t, mux, "/synse/2.0/plugins", 200, in)

	test.AddServerHost(server)

	resp, err := client.Plugins()
	test.ExpectNoError(t, err)

	if !reflect.DeepEqual(out, resp) {
		t.Errorf("expected %+v, but was %+v", out, resp)
	}
}

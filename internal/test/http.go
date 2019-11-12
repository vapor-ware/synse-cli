// Synse CLI
// Copyright (c) 2019 Vapor IO
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package test

import (
	"github.com/vapor-ware/synse-client-go/synse"
	"github.com/vapor-ware/synse-client-go/synse/scheme"
)

// FakeHTTPClientV3 implements the synse.Client interface to allow for
// basic testing without dependencies on external services.
type FakeHTTPClientV3 struct {
	cmdError bool
}

// NewFakeHTTPClientV3 creates a new fake client whose methods will never
// return an error.
func NewFakeHTTPClientV3() synse.Client {
	return &FakeHTTPClientV3{}
}

// NewFakeHTTPClientV3Err creates a new fake client whose methods will
// always return an error.
func NewFakeHTTPClientV3Err() synse.Client {
	return &FakeHTTPClientV3{
		cmdError: true,
	}
}

func (c *FakeHTTPClientV3) Status() (*scheme.Status, error) {
	if c.cmdError {
		return nil, ErrFakeClient
	}
	return &scheme.Status{
		Status:    "ok",
		Timestamp: "2019-04-22T13:30:00Z",
	}, nil
}

func (c *FakeHTTPClientV3) Version() (*scheme.Version, error) {
	if c.cmdError {
		return nil, ErrFakeClient
	}
	return &scheme.Version{
		Version:    "3.0.0",
		APIVersion: "v3",
	}, nil
}

func (c *FakeHTTPClientV3) Config() (*scheme.Config, error) {
	if c.cmdError {
		return nil, ErrFakeClient
	}
	return &scheme.Config{
		PrettyJSON: true,
		Logging:    "debug",
		Plugin: scheme.PluginOptions{
			TCP: []string{
				"localhost:5001",
			},
		},
	}, nil
}

func (c *FakeHTTPClientV3) Plugins() ([]*scheme.PluginMeta, error) {
	if c.cmdError {
		return nil, ErrFakeClient
	}
	return []*scheme.PluginMeta{
		{
			Active:      true,
			ID:          "123-456-789",
			Name:        "fake-plugin",
			Description: "a fake plugin",
			Maintainer:  "vaporio",
			Tag:         "vaporio/fake_plugin",
			VCS:         "",
			Version: scheme.VersionOptions{
				PluginVersion: "0.0.1",
				SDKVersion:    "3.0.0",
			},
		},
		{
			Active:      true,
			ID:          "987-654-321",
			Name:        "fake-plugin2",
			Description: "a fake plugin",
			Maintainer:  "vaporio",
			Tag:         "vaporio/fake_plugin2",
			VCS:         "",
			Version: scheme.VersionOptions{
				PluginVersion: "0.0.1",
				SDKVersion:    "3.0.0",
			},
		},
	}, nil
}

func (c *FakeHTTPClientV3) Plugin(string) (*scheme.Plugin, error) {
	if c.cmdError {
		return nil, ErrFakeClient
	}
	return &scheme.Plugin{
		PluginMeta: scheme.PluginMeta{
			Active:      true,
			ID:          "123-456-678",
			Name:        "fake-plugin",
			Description: "a fake plugin",
			Maintainer:  "vaporio",
			Tag:         "vaporio/fake_plugin",
			VCS:         "",
			Version: scheme.VersionOptions{
				PluginVersion: "0.0.1",
				SDKVersion:    "3.0.0",
			},
		},
		Network: scheme.NetworkOptions{
			Protocol: "tcp",
			Address:  "localhost:5001",
		},
		Health: scheme.HealthOptions{
			Timestamp: "2019-04-22T13:30:00Z",
			Status:    "OK",
			Message:   "",
			Checks:    []scheme.CheckOptions{},
		},
	}, nil
}

func (c *FakeHTTPClientV3) PluginHealth() (*scheme.PluginHealth, error) {
	if c.cmdError {
		return nil, ErrFakeClient
	}
	return &scheme.PluginHealth{
		Status:  "OK",
		Updated: "2019-04-22T13:30:00Z",
		Healthy: []string{
			"123-456-789",
		},
		Unhealthy: []string{},
		Active:    1,
		Inactive:  0,
	}, nil
}

func (c *FakeHTTPClientV3) Scan(scheme.ScanOptions) ([]*scheme.Scan, error) {
	if c.cmdError {
		return nil, ErrFakeClient
	}
	return []*scheme.Scan{
		{
			ID:     "111-222-333",
			Alias:  "fake-device",
			Info:   "a fake device",
			Type:   "faked",
			Plugin: "123-456-789",
			Tags: []string{
				"system/id:111-222-333",
				"system/type:faked",
				"vapor/fake",
			},
		},
		{
			ID:     "444-555-666",
			Alias:  "fake-device2",
			Info:   "a fake device",
			Type:   "faked",
			Plugin: "123-456-789",
			Tags: []string{
				"system/id:444-555-666",
				"system/type:faked",
				"vapor/fake",
			},
		},
	}, nil
}

func (c *FakeHTTPClientV3) Tags(scheme.TagsOptions) ([]string, error) {
	if c.cmdError {
		return nil, ErrFakeClient
	}
	return []string{
		"system/id:111-222-333",
		"system/type:faked",
		"vapor/fake",
		"default/foo",
	}, nil
}

func (c *FakeHTTPClientV3) Info(string) (*scheme.Info, error) {
	if c.cmdError {
		return nil, ErrFakeClient
	}
	return &scheme.Info{
		Timestamp: "2019-04-22T13:30:00Z",
		ID:        "111-222-333",
		Alias:     "fake-device",
		Type:      "faked",
		Metadata:  map[string]string{},
		Plugin:    "123-456-789",
		Info:      "some device",
		Tags: []string{
			"system/id:111-222-333",
			"system/type:faked",
			"vapor/fake",
		},
		Capabilities: scheme.CapabilitiesOptions{
			Mode: "rw",
			Write: scheme.WriteOptions{
				Actions: []string{
					"foo",
				},
			},
		},
		Outputs: []scheme.OutputOptions{},
	}, nil
}

func (c *FakeHTTPClientV3) Read(scheme.ReadOptions) ([]*scheme.Read, error) {
	if c.cmdError {
		return nil, ErrFakeClient
	}
	return []*scheme.Read{
		{
			Device:     "111-222-333",
			DeviceType: "faked",
			Type:       "fake",
			Value:      7,
			Timestamp:  "2019-04-22T13:30:00Z",
			Unit: scheme.UnitOptions{
				Name:   "fake unit",
				Symbol: "fu",
			},
			Context: map[string]interface{}{
				"some": "value",
			},
		},
		{
			Device:     "444-555-666",
			DeviceType: "faked",
			Type:       "fake",
			Value:      10,
			Timestamp:  "2019-04-22T13:30:00Z",
			Unit: scheme.UnitOptions{
				Name:   "fake unit",
				Symbol: "fu",
			},
			Context: map[string]interface{}{
				"some": "value",
			},
		},
	}, nil
}

func (c *FakeHTTPClientV3) ReadDevice(string) ([]*scheme.Read, error) {
	if c.cmdError {
		return nil, ErrFakeClient
	}
	return []*scheme.Read{
		{
			Device:     "111-222-333",
			DeviceType: "faked",
			Type:       "fake",
			Value:      7,
			Timestamp:  "2019-04-22T13:30:00Z",
			Unit: scheme.UnitOptions{
				Name:   "fake unit",
				Symbol: "fu",
			},
			Context: map[string]interface{}{
				"some": "value",
			},
		},
		{
			Device:     "444-555-666",
			DeviceType: "faked",
			Type:       "fake",
			Value:      10,
			Timestamp:  "2019-04-22T13:30:00Z",
			Unit: scheme.UnitOptions{
				Name:   "fake unit",
				Symbol: "fu",
			},
			Context: map[string]interface{}{
				"some": "value",
			},
		},
	}, nil
}

func (c *FakeHTTPClientV3) ReadCache(opts scheme.ReadCacheOptions, readings chan<- *scheme.Read) error {
	if c.cmdError {
		return ErrFakeClient
	}
	defer close(readings)

	vals := []*scheme.Read{
		{
			Device:     "111-222-333",
			DeviceType: "faked",
			Type:       "fake",
			Value:      7,
			Timestamp:  "2019-04-22T13:30:00Z",
			Unit: scheme.UnitOptions{
				Name:   "fake unit",
				Symbol: "fu",
			},
			Context: map[string]interface{}{
				"some": "value",
			},
		},
		{
			Device:     "444-555-666",
			DeviceType: "faked",
			Type:       "fake",
			Value:      10,
			Timestamp:  "2019-04-22T13:30:00Z",
			Unit: scheme.UnitOptions{
				Name:   "fake unit",
				Symbol: "fu",
			},
			Context: map[string]interface{}{
				"some": "value",
			},
		},
	}

	for _, r := range vals {
		readings <- r
	}
	return nil
}

func (c *FakeHTTPClientV3) ReadStream(options scheme.ReadStreamOptions, readings chan<- *scheme.Read, quit chan struct{}) error {
	if c.cmdError {
		return ErrFakeClient
	}

	for {
		select {
		case _, open := <-quit:
			if !open {
				return nil
			}
		default:
			// do nothing
		}

		readings <- &scheme.Read{
			Device:     "111-222-333",
			DeviceType: "faked",
			Type:       "fake",
			Value:      7,
			Timestamp:  "2019-04-22T13:30:00Z",
			Unit: scheme.UnitOptions{
				Name:   "fake unit",
				Symbol: "fu",
			},
			Context: map[string]interface{}{
				"some": "value",
			},
		}
	}
}

func (c *FakeHTTPClientV3) WriteAsync(string, []scheme.WriteData) ([]*scheme.Write, error) {
	if c.cmdError {
		return nil, ErrFakeClient
	}
	return []*scheme.Write{
		{
			ID:     "abc-def",
			Device: "111-222-333",
			Context: scheme.WriteData{
				Action: "foo",
				Data:   "bar",
			},
			Timeout: "10s",
		},
		{
			ID:     "fed-cba",
			Device: "111-222-333",
			Context: scheme.WriteData{
				Action: "foo",
				Data:   "baz",
			},
			Timeout: "10s",
		},
	}, nil
}

func (c *FakeHTTPClientV3) WriteSync(string, []scheme.WriteData) ([]*scheme.Transaction, error) {
	if c.cmdError {
		return nil, ErrFakeClient
	}
	return []*scheme.Transaction{
		{
			ID:      "abc-def",
			Timeout: "10s",
			Device:  "111-222-333",
			Context: scheme.WriteData{
				Action: "foo",
				Data:   "bar",
			},
			Status:  "DONE",
			Created: "2019-04-22T13:30:00Z",
			Updated: "2019-04-22T13:30:00Z",
			Message: "",
		},
		{
			ID:      "fed-cba",
			Timeout: "10s",
			Device:  "111-222-333",
			Context: scheme.WriteData{
				Action: "foo",
				Data:   "baz",
			},
			Status:  "DONE",
			Created: "2019-04-22T13:30:00Z",
			Updated: "2019-04-22T13:30:00Z",
			Message: "",
		},
	}, nil
}

func (c *FakeHTTPClientV3) Transactions() ([]string, error) {
	if c.cmdError {
		return nil, ErrFakeClient
	}
	return []string{
		"abc-def",
		"fed-cba",
		"123-456",
	}, nil
}

func (c *FakeHTTPClientV3) Transaction(string) (*scheme.Transaction, error) {
	if c.cmdError {
		return nil, ErrFakeClient
	}
	return &scheme.Transaction{
		ID:      "abc-def",
		Timeout: "10s",
		Device:  "111-222-333",
		Context: scheme.WriteData{
			Action: "foo",
			Data:   "bar",
		},
		Status:  "DONE",
		Created: "2019-04-22T13:30:00Z",
		Updated: "2019-04-22T13:30:00Z",
		Message: "",
	}, nil
}

func (c *FakeHTTPClientV3) GetOptions() *synse.Options {
	return &synse.Options{}
}

func (c *FakeHTTPClientV3) Open() error {
	if c.cmdError {
		return ErrFakeClient
	}
	return nil
}

func (c *FakeHTTPClientV3) Close() error {
	if c.cmdError {
		return ErrFakeClient
	}
	return nil
}

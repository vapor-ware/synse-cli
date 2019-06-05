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

package context

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/vapor-ware/synse-cli/internal/golden"
	"github.com/vapor-ware/synse-cli/internal/test"
	"github.com/vapor-ware/synse-cli/pkg/config"
	"github.com/vapor-ware/synse-cli/pkg/utils"
	"os"
	"testing"
)

func TestCmdCurrent_multipleFmtFlags(t *testing.T) {
	out := bytes.Buffer{}
	fakeExiter := test.FakeExiter{Writer: &out}
	exitutil = &fakeExiter
	defer func() {
		exitutil = utils.NewDefaultExiter()
		config.Purge()
		resetFlags()
	}()

	os.Args = []string{"current", "--yaml", "--json"}
	err := cmdCurrent.Execute()
	assert.NoError(t, err)
	assert.True(t, fakeExiter.IsExited)
	golden.Check(t, out.Bytes(), "current.multiple-fmt-flags.golden")

	contexts := config.GetContexts()
	current := config.GetCurrentContext()
	assert.Len(t, contexts, 0)
	assert.Len(t, current, 0)
}

func TestCmdCurrent_extraArgs(t *testing.T) {
	out := bytes.Buffer{}
	cmdCurrent.SetOutput(&out)

	os.Args = []string{"current", "foo", "bar"}
	err := cmdCurrent.Execute()
	assert.Error(t, err)
	golden.Check(t, out.Bytes(), "current.extra-args.golden")

	contexts := config.GetContexts()
	current := config.GetCurrentContext()

	assert.Len(t, contexts, 0)
	assert.Len(t, current, 0)
}

func TestCmdCurrent_badCtxType(t *testing.T) {
	out := bytes.Buffer{}
	fakeExiter := test.FakeExiter{Writer: &out}
	exitutil = &fakeExiter
	defer func() {
		exitutil = utils.NewDefaultExiter()
		resetFlags()
	}()

	os.Args = []string{"current", "bad-ctx"}
	_ = cmdCurrent.Execute()
	golden.Check(t, out.Bytes(), "current.bad-ctx.golden")

	assert.True(t, fakeExiter.IsExited)

	contexts := config.GetContexts()
	current := config.GetCurrentContext()

	assert.Len(t, contexts, 0)
	assert.Len(t, current, 0)
}

func TestCmdCurrent_noCurrentServerContext(t *testing.T) {
	out := bytes.Buffer{}
	fakeExiter := test.FakeExiter{Writer: &out}
	exitutil = &fakeExiter
	defer func() {
		exitutil = utils.NewDefaultExiter()
		resetFlags()
	}()

	os.Args = []string{"current", "server"}
	_ = cmdCurrent.Execute()
	golden.Check(t, out.Bytes(), "current.no-server.golden")

	assert.True(t, fakeExiter.IsExited)

	contexts := config.GetContexts()
	current := config.GetCurrentContext()

	assert.Len(t, contexts, 0)
	assert.Len(t, current, 0)
}

func TestCmdCurrent_noCurrenPluginContext(t *testing.T) {
	out := bytes.Buffer{}
	fakeExiter := test.FakeExiter{Writer: &out}
	exitutil = &fakeExiter
	defer func() {
		exitutil = utils.NewDefaultExiter()
		config.Purge()
		resetFlags()
	}()

	err := config.AddContext(&config.ContextRecord{
		Name: "test-ctx",
		Type: "server",
		Context: config.Context{
			Address: "0.0.0.0",
			ClientCert: "/tmp/test/dir",
		},
	})
	assert.NoError(t, err)
	err = config.SetCurrentContext("test-ctx")
	assert.NoError(t, err)

	os.Args = []string{"current", "plugin"}
	_ = cmdCurrent.Execute()
	golden.Check(t, out.Bytes(), "current.no-plugin.golden")

	assert.True(t, fakeExiter.IsExited)

	contexts := config.GetContexts()
	current := config.GetCurrentContext()

	assert.Len(t, contexts, 1)
	assert.Len(t, current, 1)
}

func TestCmdCurrent_table(t *testing.T) {
	out := bytes.Buffer{}
	cmdCurrent.SetOutput(&out)
	defer func() {
		config.Purge()
		resetFlags()
	}()

	err := config.AddContext(&config.ContextRecord{
		Name: "test-ctx",
		Type: "server",
		Context: config.Context{
			Address: "0.0.0.0",
			ClientCert: "/tmp/test/dir",
		},
	})
	assert.NoError(t, err)
	err = config.SetCurrentContext("test-ctx")
	assert.NoError(t, err)

	os.Args = []string{"current", "server"}
	err = cmdCurrent.Execute()
	assert.NoError(t, err)
	golden.Check(t, out.Bytes(), "current.table.golden")

	contexts := config.GetContexts()
	current := config.GetCurrentContext()

	assert.Len(t, contexts, 1)
	assert.Len(t, current, 1)
}

func TestCmdCurrent_table2(t *testing.T) {
	out := bytes.Buffer{}
	cmdCurrent.SetOutput(&out)
	defer func() {
		config.Purge()
		resetFlags()
	}()

	err := config.AddContext(&config.ContextRecord{
		Name: "server-ctx",
		Type: "server",
		Context: config.Context{
			Address: "0.0.0.0",
			ClientCert: "/tmp/test/dir",
		},
	})
	assert.NoError(t, err)
	err = config.SetCurrentContext("server-ctx")
	assert.NoError(t, err)

	err = config.AddContext(&config.ContextRecord{
		Name: "plugin-ctx",
		Type: "plugin",
		Context: config.Context{
			Address: "foo/bar",
			ClientCert: "/tmp/test/dir",
		},
	})
	assert.NoError(t, err)
	err = config.SetCurrentContext("plugin-ctx")
	assert.NoError(t, err)

	os.Args = []string{"current"}
	err = cmdCurrent.Execute()
	assert.NoError(t, err)
	golden.Check(t, out.Bytes(), "current.table2.golden")

	contexts := config.GetContexts()
	current := config.GetCurrentContext()

	assert.Len(t, contexts, 2)
	assert.Len(t, current, 2)
}

func TestCmdCurrent_tableNoHeader(t *testing.T) {
	out := bytes.Buffer{}
	cmdCurrent.SetOutput(&out)
	defer func() {
		config.Purge()
		resetFlags()
	}()


	err := config.AddContext(&config.ContextRecord{
		Name: "test-ctx",
		Type: "server",
		Context: config.Context{
			Address: "0.0.0.0",
			ClientCert: "/tmp/test/dir",
		},
	})
	assert.NoError(t, err)
	err = config.SetCurrentContext("test-ctx")
	assert.NoError(t, err)

	os.Args = []string{"current", "server", "--no-header"}
	err = cmdCurrent.Execute()
	assert.NoError(t, err)
	golden.Check(t, out.Bytes(), "current.table-no-header.golden")

	contexts := config.GetContexts()
	current := config.GetCurrentContext()

	assert.Len(t, contexts, 1)
	assert.Len(t, current, 1)
}

func TestCmdCurrent_yaml(t *testing.T) {
	out := bytes.Buffer{}
	cmdCurrent.SetOutput(&out)
	defer func() {
		config.Purge()
		resetFlags()
	}()

	err := config.AddContext(&config.ContextRecord{
		Name: "test-ctx",
		Type: "server",
		Context: config.Context{
			Address: "0.0.0.0",
			ClientCert: "/tmp/test/dir",
		},
	})
	assert.NoError(t, err)
	err = config.SetCurrentContext("test-ctx")
	assert.NoError(t, err)

	os.Args = []string{"current", "server", "--yaml"}
	err = cmdCurrent.Execute()
	assert.NoError(t, err)
	golden.Check(t, out.Bytes(), "current.yaml.golden")

	contexts := config.GetContexts()
	current := config.GetCurrentContext()

	assert.Len(t, contexts, 1)
	assert.Len(t, current, 1)
}

func TestCmdCurrent_yaml2(t *testing.T) {
	out := bytes.Buffer{}
	cmdCurrent.SetOutput(&out)
	defer func() {
		config.Purge()
		resetFlags()
	}()

	err := config.AddContext(&config.ContextRecord{
		Name: "server-ctx",
		Type: "server",
		Context: config.Context{
			Address: "0.0.0.0",
			ClientCert: "/tmp/test/dir",
		},
	})
	assert.NoError(t, err)
	err = config.SetCurrentContext("server-ctx")
	assert.NoError(t, err)

	err = config.AddContext(&config.ContextRecord{
		Name: "plugin-ctx",
		Type: "plugin",
		Context: config.Context{
			Address: "foo/bar",
			ClientCert: "/tmp/test/dir",
		},
	})
	assert.NoError(t, err)
	err = config.SetCurrentContext("plugin-ctx")
	assert.NoError(t, err)

	os.Args = []string{"current", "--yaml"}
	err = cmdCurrent.Execute()
	assert.NoError(t, err)
	golden.Check(t, out.Bytes(), "current.yaml2.golden")

	contexts := config.GetContexts()
	current := config.GetCurrentContext()

	assert.Len(t, contexts, 2)
	assert.Len(t, current, 2)
}

func TestCmdCurrent_json(t *testing.T) {
	out := bytes.Buffer{}
	cmdCurrent.SetOutput(&out)
	defer func() {
		config.Purge()
		resetFlags()
	}()


	err := config.AddContext(&config.ContextRecord{
		Name: "test-ctx",
		Type: "server",
		Context: config.Context{
			Address: "0.0.0.0",
			ClientCert: "/tmp/test/dir",
		},
	})
	assert.NoError(t, err)
	err = config.SetCurrentContext("test-ctx")
	assert.NoError(t, err)

	os.Args = []string{"current", "server", "--json"}
	err = cmdCurrent.Execute()
	assert.NoError(t, err)
	golden.Check(t, out.Bytes(), "current.json.golden")

	contexts := config.GetContexts()
	current := config.GetCurrentContext()

	assert.Len(t, contexts, 1)
	assert.Len(t, current, 1)
}

func TestCmdCurrent_json2(t *testing.T) {
	out := bytes.Buffer{}
	cmdCurrent.SetOutput(&out)
	defer func() {
		config.Purge()
		resetFlags()
	}()

	err := config.AddContext(&config.ContextRecord{
		Name: "server-ctx",
		Type: "server",
		Context: config.Context{
			Address: "0.0.0.0",
			ClientCert: "/tmp/test/dir",
		},
	})
	assert.NoError(t, err)
	err = config.SetCurrentContext("server-ctx")
	assert.NoError(t, err)

	err = config.AddContext(&config.ContextRecord{
		Name: "plugin-ctx",
		Type: "plugin",
		Context: config.Context{
			Address: "foo/bar",
			ClientCert: "/tmp/test/dir",
		},
	})
	assert.NoError(t, err)
	err = config.SetCurrentContext("plugin-ctx")
	assert.NoError(t, err)

	os.Args = []string{"current", "--json"}
	err = cmdCurrent.Execute()
	assert.NoError(t, err)
	golden.Check(t, out.Bytes(), "current.json2.golden")

	contexts := config.GetContexts()
	current := config.GetCurrentContext()

	assert.Len(t, contexts, 2)
	assert.Len(t, current, 2)
}
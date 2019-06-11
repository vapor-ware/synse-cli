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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vapor-ware/synse-cli/internal/test"
	"github.com/vapor-ware/synse-cli/pkg/config"
)

func TestCmdCurrent_multipleFmtFlags(t *testing.T) {
	defer func() {
		config.Purge()
		resetFlags()
	}()

	result := test.Cmd(cmdCurrent).Args(
		"--yaml",
		"--json",
	).Run(t)
	result.AssertNoErr()
	result.AssertExited()
	result.AssertGolden("current.multiple-fmt-flags.golden")

	assert.Len(t, config.GetContexts(), 0)
	assert.Len(t, config.GetCurrentContext(), 0)
}

func TestCmdCurrent_extraArgs(t *testing.T) {
	defer func() {
		config.Purge()
		resetFlags()
	}()

	result := test.Cmd(cmdCurrent).Args(
		"foo",
		"bar",
	).Run(t)
	result.AssertErr()
	result.AssertGolden("current.extra-args.golden")

	assert.Len(t, config.GetContexts(), 0)
	assert.Len(t, config.GetCurrentContext(), 0)
}

func TestCmdCurrent_badCtxType(t *testing.T) {
	defer func() {
		config.Purge()
		resetFlags()
	}()

	result := test.Cmd(cmdCurrent).Args(
		"bad-ctx",
	).Run(t)
	result.AssertNoErr()
	result.AssertExited()
	result.AssertGolden("current.bad-ctx.golden")

	assert.Len(t, config.GetContexts(), 0)
	assert.Len(t, config.GetCurrentContext(), 0)
}

func TestCmdCurrent_noCurrentServerContext(t *testing.T) {
	defer func() {
		config.Purge()
		resetFlags()
	}()

	result := test.Cmd(cmdCurrent).Args(
		"server",
	).Run(t)
	result.AssertNoErr()
	result.AssertExited()
	result.AssertGolden("current.no-server.golden")

	assert.Len(t, config.GetContexts(), 0)
	assert.Len(t, config.GetCurrentContext(), 0)
}

func TestCmdCurrent_noCurrentPluginContext(t *testing.T) {
	defer func() {
		config.Purge()
		resetFlags()
	}()

	assert.NoError(t, config.AddContext(&config.ContextRecord{
		Name: "test-ctx",
		Type: "server",
		Context: config.Context{
			Address:    "0.0.0.0",
			ClientCert: "/tmp/test/dir",
		},
	}))
	assert.NoError(t, config.SetCurrentContext("test-ctx"))

	result := test.Cmd(cmdCurrent).Args(
		"plugin",
	).Run(t)
	result.AssertNoErr()
	result.AssertExited()
	result.AssertGolden("current.no-plugin.golden")

	assert.Len(t, config.GetContexts(), 1)
	assert.Len(t, config.GetCurrentContext(), 1)
}

func TestCmdCurrent_table(t *testing.T) {
	defer func() {
		config.Purge()
		resetFlags()
	}()

	assert.NoError(t, config.AddContext(&config.ContextRecord{
		Name: "test-ctx",
		Type: "server",
		Context: config.Context{
			Address:    "0.0.0.0",
			ClientCert: "/tmp/test/dir",
		},
	}))
	assert.NoError(t, config.SetCurrentContext("test-ctx"))

	result := test.Cmd(cmdCurrent).Args(
		"server",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("current.table.golden")

	assert.Len(t, config.GetContexts(), 1)
	assert.Len(t, config.GetCurrentContext(), 1)
}

func TestCmdCurrent_table2(t *testing.T) {
	defer func() {
		config.Purge()
		resetFlags()
	}()

	// Set a current server context
	assert.NoError(t, config.AddContext(&config.ContextRecord{
		Name: "server-ctx",
		Type: "server",
		Context: config.Context{
			Address:    "0.0.0.0",
			ClientCert: "/tmp/test/dir",
		},
	}))
	assert.NoError(t, config.SetCurrentContext("server-ctx"))

	// Set a current plugin context
	assert.NoError(t, config.AddContext(&config.ContextRecord{
		Name: "plugin-ctx",
		Type: "plugin",
		Context: config.Context{
			Address:    "foo/bar",
			ClientCert: "/tmp/test/dir",
		},
	}))
	assert.NoError(t, config.SetCurrentContext("plugin-ctx"))

	result := test.Cmd(cmdCurrent).Run(t)
	result.AssertNoErr()
	result.AssertGolden("current.table2.golden")

	assert.Len(t, config.GetContexts(), 2)
	assert.Len(t, config.GetCurrentContext(), 2)
}

func TestCmdCurrent_tableNoHeader(t *testing.T) {
	defer func() {
		config.Purge()
		resetFlags()
	}()

	assert.NoError(t, config.AddContext(&config.ContextRecord{
		Name: "test-ctx",
		Type: "server",
		Context: config.Context{
			Address:    "0.0.0.0",
			ClientCert: "/tmp/test/dir",
		},
	}))
	assert.NoError(t, config.SetCurrentContext("test-ctx"))

	result := test.Cmd(cmdCurrent).Args(
		"server",
		"--no-header",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("current.table-no-header.golden")

	assert.Len(t, config.GetContexts(), 1)
	assert.Len(t, config.GetCurrentContext(), 1)
}

func TestCmdCurrent_yaml(t *testing.T) {
	defer func() {
		config.Purge()
		resetFlags()
	}()

	assert.NoError(t, config.AddContext(&config.ContextRecord{
		Name: "test-ctx",
		Type: "server",
		Context: config.Context{
			Address:    "0.0.0.0",
			ClientCert: "/tmp/test/dir",
		},
	}))
	assert.NoError(t, config.SetCurrentContext("test-ctx"))

	result := test.Cmd(cmdCurrent).Args(
		"server",
		"--yaml",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("current.yaml.golden")

	assert.Len(t, config.GetContexts(), 1)
	assert.Len(t, config.GetCurrentContext(), 1)
}

func TestCmdCurrent_yaml2(t *testing.T) {
	defer func() {
		config.Purge()
		resetFlags()
	}()

	// Set a current server context
	assert.NoError(t, config.AddContext(&config.ContextRecord{
		Name: "server-ctx",
		Type: "server",
		Context: config.Context{
			Address:    "0.0.0.0",
			ClientCert: "/tmp/test/dir",
		},
	}))
	assert.NoError(t, config.SetCurrentContext("server-ctx"))

	// Set a current plugin context
	assert.NoError(t, config.AddContext(&config.ContextRecord{
		Name: "plugin-ctx",
		Type: "plugin",
		Context: config.Context{
			Address:    "foo/bar",
			ClientCert: "/tmp/test/dir",
		},
	}))
	assert.NoError(t, config.SetCurrentContext("plugin-ctx"))

	result := test.Cmd(cmdCurrent).Args(
		"--yaml",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("current.yaml2.golden")

	assert.Len(t, config.GetContexts(), 2)
	assert.Len(t, config.GetCurrentContext(), 2)
}

func TestCmdCurrent_json(t *testing.T) {
	defer func() {
		config.Purge()
		resetFlags()
	}()

	assert.NoError(t, config.AddContext(&config.ContextRecord{
		Name: "test-ctx",
		Type: "server",
		Context: config.Context{
			Address:    "0.0.0.0",
			ClientCert: "/tmp/test/dir",
		},
	}))
	assert.NoError(t, config.SetCurrentContext("test-ctx"))

	result := test.Cmd(cmdCurrent).Args(
		"server",
		"--json",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("current.json.golden")

	assert.Len(t, config.GetContexts(), 1)
	assert.Len(t, config.GetCurrentContext(), 1)
}

func TestCmdCurrent_json2(t *testing.T) {
	defer func() {
		config.Purge()
		resetFlags()
	}()

	// Set a current server context
	assert.NoError(t, config.AddContext(&config.ContextRecord{
		Name: "server-ctx",
		Type: "server",
		Context: config.Context{
			Address:    "0.0.0.0",
			ClientCert: "/tmp/test/dir",
		},
	}))
	assert.NoError(t, config.SetCurrentContext("server-ctx"))

	// Set a current plugin context
	assert.NoError(t, config.AddContext(&config.ContextRecord{
		Name: "plugin-ctx",
		Type: "plugin",
		Context: config.Context{
			Address:    "foo/bar",
			ClientCert: "/tmp/test/dir",
		},
	}))
	assert.NoError(t, config.SetCurrentContext("plugin-ctx"))

	result := test.Cmd(cmdCurrent).Args(
		"--json",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("current.json2.golden")

	assert.Len(t, config.GetContexts(), 2)
	assert.Len(t, config.GetCurrentContext(), 2)
}

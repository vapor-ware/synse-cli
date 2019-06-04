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
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vapor-ware/synse-cli/internal/golden"
	"github.com/vapor-ware/synse-cli/internal/test"
	"github.com/vapor-ware/synse-cli/pkg/config"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

func TestCmdAdd_notEnoughArgs(t *testing.T) {
	out := bytes.Buffer{}
	cmdAdd.SetOutput(&out)

	os.Args = []string{"add"}
	err := cmdAdd.Execute()
	assert.Error(t, err)
	golden.Check(t, out.Bytes(), "add.no-args.golden")

	contexts := config.GetContexts()
	current := config.GetCurrentContext()

	assert.Len(t, contexts, 0)
	assert.Len(t, current, 0)
}

func TestCmdAdd_badCtxType(t *testing.T) {
	out := bytes.Buffer{}
	fakeExiter := test.FakeExiter{Writer: &out}
	exitutil = &fakeExiter
	defer func() {
		exitutil = utils.NewDefaultExiter()
	}()

	os.Args = []string{"add", "bad-type", "test-name", "test-address"}
	_ = cmdAdd.Execute()
	golden.Check(t, out.Bytes(), "add.bad-type.golden")

	assert.True(t, fakeExiter.IsExited)

	contexts := config.GetContexts()
	current := config.GetCurrentContext()

	assert.Len(t, contexts, 0)
	assert.Len(t, current, 0)
}

func TestCmdAdd_duplicate(t *testing.T) {
	out := bytes.Buffer{}
	fakeExiter := test.FakeExiter{Writer: &out}
	exitutil = &fakeExiter
	defer func() {
		config.Purge()
		exitutil = utils.NewDefaultExiter()
	}()

	err := config.AddContext(&config.ContextRecord{
		Name: "test-ctx",
		Type: "server",
	})
	assert.NoError(t, err)

	os.Args = []string{"add", "server", "test-ctx", "test-address"}
	_ = cmdAdd.Execute()
	golden.Check(t, out.Bytes(), "add.duplicate-add.golden")

	assert.True(t, fakeExiter.IsExited)

	contexts := config.GetContexts()
	current := config.GetCurrentContext()

	assert.Len(t, contexts, 1)
	assert.Len(t, current, 0)
}

func TestCmdAdd_addContext(t *testing.T) {
	defer config.Purge()

	out := bytes.Buffer{}
	cmdAdd.SetOutput(&out)

	os.Args = []string{"add", "server", "test_name", "test_address"}
	err := cmdAdd.Execute()
	assert.NoError(t, err)
	golden.Check(t, out.Bytes(), "empty.golden")

	contexts := config.GetContexts()
	current := config.GetCurrentContext()

	assert.Len(t, contexts, 1)
	assert.Len(t, current, 0)

	assert.Equal(t, contexts[0].Name, "test_name")
	assert.Equal(t, contexts[0].Type, "server")
	assert.Equal(t, contexts[0].Context.Address, "test_address")
	assert.Equal(t, contexts[0].Context.ClientCert, "")
}

func TestCmdAdd_addAndSetContext(t *testing.T) {
	defer config.Purge()

	out := bytes.Buffer{}
	cmdAdd.SetOutput(&out)

	os.Args = []string{"add", "server", "test_name", "test_address", "--set"}
	err := cmdAdd.Execute()
	assert.NoError(t, err)
	golden.Check(t, out.Bytes(), "empty.golden")

	contexts := config.GetContexts()
	current := config.GetCurrentContext()

	assert.Len(t, contexts, 1)
	assert.Len(t, current, 1)

	assert.Equal(t, contexts[0].Name, "test_name")
	assert.Equal(t, contexts[0].Type, "server")
	assert.Equal(t, contexts[0].Context.Address, "test_address")
	assert.Equal(t, contexts[0].Context.ClientCert, "")

	serverCtx, ok := current["server"]
	assert.True(t, ok)
	assert.Equal(t, serverCtx.Name, "test_name")
	assert.Equal(t, serverCtx.Type, "server")
	assert.Equal(t, serverCtx.Context.Address, "test_address")
	assert.Equal(t, serverCtx.Context.ClientCert, "")

	pluginCtx, ok := current["plugin"]
	assert.False(t, ok)
	assert.Nil(t, pluginCtx)
}

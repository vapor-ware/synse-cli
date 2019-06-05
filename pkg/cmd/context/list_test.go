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

func TestCmdList_noContexts(t *testing.T) {
	out := bytes.Buffer{}
	cmdList.SetOutput(&out)

	os.Args = []string{"list"}
	err := cmdList.Execute()
	assert.NoError(t, err)
	golden.Check(t, out.Bytes(), "empty.golden")

	contexts := config.GetContexts()
	current := config.GetCurrentContext()
	assert.Len(t, contexts, 0)
	assert.Len(t, current, 0)
}

func TestCmdList_multipleFmtFlags(t *testing.T) {
	out := bytes.Buffer{}
	fakeExiter := test.FakeExiter{Writer: &out}
	exitutil = &fakeExiter
	defer func() {
		exitutil = utils.NewDefaultExiter()
		config.Purge()
		resetFlags()
	}()

	os.Args = []string{"list", "--yaml", "--json"}
	err := cmdList.Execute()
	assert.NoError(t, err)
	assert.True(t, fakeExiter.IsExited)
	golden.Check(t, out.Bytes(), "list.multiple-fmt-flags.golden")

	contexts := config.GetContexts()
	current := config.GetCurrentContext()
	assert.Len(t, contexts, 0)
	assert.Len(t, current, 0)
}

func TestCmdList(t *testing.T) {
	out := bytes.Buffer{}
	cmdList.SetOutput(&out)
	defer func() {
		config.Purge()
		resetFlags()
	}()

	err := config.AddContext(&config.ContextRecord{
		Name: "test-ctx-1",
		Type: "server",
		Context: config.Context{
			Address: "0.0.0.0",
			ClientCert: "/tmp/test/dir",
		},
	})
	assert.NoError(t, err)
	err = config.SetCurrentContext("test-ctx-1")
	assert.NoError(t, err)

	err = config.AddContext(&config.ContextRecord{
		Name: "test-ctx-2",
		Type: "plugin",
		Context: config.Context{
			Address: "foo/bar",
			ClientCert: "/tmp/test/dir",
		},
	})
	assert.NoError(t, err)

	os.Args = []string{"list"}
	err = cmdList.Execute()
	assert.NoError(t, err)
	golden.Check(t, out.Bytes(), "list.golden")

	contexts := config.GetContexts()
	current := config.GetCurrentContext()
	assert.Len(t, contexts, 2)
	assert.Len(t, current, 1)
}
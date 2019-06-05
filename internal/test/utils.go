package test

import (
	"bytes"
	"os"
	"testing"

	"bou.ke/monkey"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/vapor-ware/synse-cli/internal/golden"
)

type Result struct {
	t      *testing.T
	err    error
	exited bool
	out    []byte
}

func (r *Result) AssertNoErr() {
	assert.NoError(r.t, r.err)
}

func (r *Result) AssertErr() {
	assert.Error(r.t, r.err)
}

func (r *Result) AssertGolden(filename string) {
	golden.Check(r.t, r.out, filename)
}

func (r *Result) AssertExited() {
	assert.True(r.t, r.exited)
}

type Builder struct {
	cmd  *cobra.Command
	name string
	args []string
	t    *testing.T
}

func (b *Builder) Args(args ...string) *Builder {
	b.args = append(b.args, args...)
	return b
}

func (b *Builder) Run(t *testing.T) (result *Result) {
	b.t = t

	cmdOut := bytes.Buffer{}
	b.cmd.SetOutput(&cmdOut)

	var exitCalled bool
	patch := monkey.Patch(os.Exit, func(code int) {
		exitCalled = true
		// the exiter is expected to terminate the program, so if we
		// do not stop command execution here, it may continue on to
		// run other commands, which would generate test results which
		// do not reflect reality. to remedy this, we panic here and
		// catch it on defer.
		panic("exitpanic")
	})
	defer patch.Unpatch()

	defer func() {
		if r := recover(); r != nil {
			msg, ok := r.(string)
			if !ok {
				b.t.Fatal(r)
			}

			if msg == "exitpanic" {
				result = &Result{
					t:      b.t,
					err:    nil,
					exited: exitCalled,
					out:    cmdOut.Bytes(),
				}
			}
		}
	}()

	os.Args = b.args
	err := b.cmd.Execute()
	result = &Result{
		t:      b.t,
		err:    err,
		exited: exitCalled,
		out:    cmdOut.Bytes(),
	}

	return
}

func Cmd(cmd *cobra.Command) *Builder {
	return &Builder{
		cmd:  cmd,
		name: cmd.Name(),
		args: []string{cmd.Name()},
	}
}

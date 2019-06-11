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
	root string
	name string
	args []string
	t    *testing.T
}

func (b *Builder) Args(args ...string) *Builder {
	b.args = append(b.args, args...)
	return b
}

func (b *Builder) WithRoot(root string) *Builder {
	b.root = root
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
			} else {
				panic(r)
			}
		}
	}()

	var args []string
	if b.root != "" {
		args = append(args, b.root)
	}
	args = append(args, b.args...)

	os.Args = args
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

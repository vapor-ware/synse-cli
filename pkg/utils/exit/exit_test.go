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

package exit

import (
	"bytes"
	"os"
	"testing"

	"bou.ke/monkey"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestFromCmd(t *testing.T) {
	cmd := &cobra.Command{
		Use: "foo",
	}

	exiter := FromCmd(cmd)
	assert.NotNil(t, exiter)
}

func TestCommandExiter_Err(t *testing.T) {
	var exitCalled bool
	patch := monkey.Patch(os.Exit, func(code int) {
		exitCalled = true
		assert.Equal(t, 1, code)
	})
	defer patch.Unpatch()

	out := bytes.Buffer{}
	exiter := commandExiter{out: &out}
	exiter.Err("test error")

	assert.True(t, exitCalled)
	assert.Equal(t, out.String(), "Error: test error\n")
}

func TestCommandExiter_Exit(t *testing.T) {
	var exitCalled bool
	patch := monkey.Patch(os.Exit, func(code int) {
		exitCalled = true
		assert.Equal(t, 2, code)
	})
	defer patch.Unpatch()

	out := bytes.Buffer{}
	exiter := commandExiter{out: &out}
	exiter.Exit(2)

	assert.True(t, exitCalled)
	assert.Equal(t, out.String(), "")
}

func TestCommandExiter_Exitf(t *testing.T) {
	var exitCalled bool
	patch := monkey.Patch(os.Exit, func(code int) {
		exitCalled = true
		assert.Equal(t, 3, code)
	})
	defer patch.Unpatch()

	out := bytes.Buffer{}
	exiter := commandExiter{out: &out}
	exiter.Exitf(3, "test error")

	assert.True(t, exitCalled)
	assert.Equal(t, out.String(), "test error")
}

func TestCommandExiter_Fatal(t *testing.T) {
	var exitCalled bool
	patch := monkey.Patch(os.Exit, func(code int) {
		exitCalled = true
		assert.Equal(t, 1, code)
	})
	defer patch.Unpatch()

	out := bytes.Buffer{}
	exiter := commandExiter{out: &out}
	exiter.Fatal("test error")

	assert.True(t, exitCalled)
	assert.Equal(t, out.String(), "Error: test error\n")
}

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

package utils

import (
	"bytes"
	"os"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

func TestNewDefaultExiter(t *testing.T) {
	exiter := NewDefaultExiter()
	assert.NotNil(t, exiter)
	assert.Equal(t, exiter.writer, os.Stderr)
}

func TestDefaultExiter_SetWriter(t *testing.T) {
	exiter := DefaultExiter{}
	assert.Nil(t, exiter.writer)

	exiter.SetWriter(os.Stdout)
	assert.Equal(t, exiter.writer, os.Stdout)
}

func TestDefaultExiter_Err(t *testing.T) {
	var exitCalled bool
	patch := monkey.Patch(os.Exit, func(code int) {
		exitCalled = true
		assert.Equal(t, 1, code)
	})
	defer patch.Unpatch()

	out := bytes.Buffer{}
	exiter := DefaultExiter{writer: &out}
	exiter.Err("test error")

	assert.True(t, exitCalled)
	assert.Equal(t, out.String(), "Error: test error\n")
}

func TestDefaultExiter_Exit(t *testing.T) {
	var exitCalled bool
	patch := monkey.Patch(os.Exit, func(code int) {
		exitCalled = true
		assert.Equal(t, 2, code)
	})
	defer patch.Unpatch()

	out := bytes.Buffer{}
	exiter := DefaultExiter{writer: &out}
	exiter.Exit(2)

	assert.True(t, exitCalled)
	assert.Equal(t, out.String(), "")
}

func TestDefaultExiter_Exitf(t *testing.T) {
	var exitCalled bool
	patch := monkey.Patch(os.Exit, func(code int) {
		exitCalled = true
		assert.Equal(t, 3, code)
	})
	defer patch.Unpatch()

	out := bytes.Buffer{}
	exiter := DefaultExiter{writer: &out}
	exiter.Exitf(3, "test error")

	assert.True(t, exitCalled)
	assert.Equal(t, out.String(), "test error")
}

func TestDefaultExiter_Fatal(t *testing.T) {
	var exitCalled bool
	patch := monkey.Patch(os.Exit, func(code int) {
		exitCalled = true
		assert.Equal(t, 1, code)
	})
	defer patch.Unpatch()

	out := bytes.Buffer{}
	exiter := DefaultExiter{writer: &out}
	exiter.Fatal("test error")

	assert.True(t, exitCalled)
	assert.Equal(t, out.String(), "Error: test error\n")
}

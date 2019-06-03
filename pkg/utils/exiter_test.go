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
	"os"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

func TestDefaultExiter_Err(t *testing.T) {
	var exitCalled bool
	patch := monkey.Patch(os.Exit, func(code int) {
		exitCalled = true
		assert.Equal(t, 1, code)
	})
	defer patch.Unpatch()

	exiter := DefaultExiter{}
	exiter.Err("test error")

	assert.True(t, exitCalled)
}

func TestDefaultExiter_Exit(t *testing.T) {
	var exitCalled bool
	patch := monkey.Patch(os.Exit, func(code int) {
		exitCalled = true
		assert.Equal(t, 2, code)
	})
	defer patch.Unpatch()

	exiter := DefaultExiter{}
	exiter.Exit(2)

	assert.True(t, exitCalled)
}

func TestDefaultExiter_Exitf(t *testing.T) {
	var exitCalled bool
	patch := monkey.Patch(os.Exit, func(code int) {
		exitCalled = true
		assert.Equal(t, 3, code)
	})
	defer patch.Unpatch()

	exiter := DefaultExiter{}
	exiter.Exitf(3, "test error")

	assert.True(t, exitCalled)
}

func TestDefaultExiter_Fatal(t *testing.T) {
	var exitCalled bool
	patch := monkey.Patch(os.Exit, func(code int) {
		exitCalled = true
		assert.Equal(t, 1, code)
	})
	defer patch.Unpatch()

	exiter := DefaultExiter{}
	exiter.Fatal("test error")

	assert.True(t, exitCalled)
}

func TestErr(t *testing.T) {
	var exitCalled bool
	patch := monkey.Patch(os.Exit, func(code int) {
		exitCalled = true
		assert.Equal(t, 1, code)
	})
	defer patch.Unpatch()

	Err("test error")

	assert.True(t, exitCalled)
}

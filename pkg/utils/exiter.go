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
	"fmt"
	"os"
)

var defaultExiter DefaultExiter

// Exiter is an interface for exiting the CLI.
//
// It is a useful way to test command exiting without terminating.
type Exiter interface {
	Exit(code int)
	Exitf(code int, format string, a ...interface{})
	Err(err interface{})
	Fatal(msg interface{})
}

// TODO: could specify an IO Writer for the exiter.. makes it easier for testing.

// DefaultExiter is the default Exiter implementation that the CLI uses.
type DefaultExiter struct{}

// Exit terminates the application.
func (exiter *DefaultExiter) Exit(code int) {
	os.Exit(code)
}

// Exit prints a message and terminates the application.
func (exiter *DefaultExiter) Exitf(code int, format string, a ...interface{}) {
	fmt.Printf(format, a...)
	os.Exit(code)
}

// Err checks if the input is nil; if not it will exit via Fatal.
func (exiter *DefaultExiter) Err(err interface{}) {
	if err != nil {
		exiter.Fatal(err)
	}
}

// Fatal prints a message to console and terminates the application.
func (exiter *DefaultExiter) Fatal(msg interface{}) {
	exiter.Exitf(1, "Error: %s\n", msg)
}

// Err is a utility function which prints an error message and terminates
// the application if it is passed an error.
func Err(err interface{}) {
	defaultExiter.Err(err)
}

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
	"fmt"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Exiter is an interface for exiting the CLI.
//
// It is a useful way to test command exiting without terminating.
type Exiter interface {
	Exit(code int)
	Exitf(code int, format string, a ...interface{})
	Err(err interface{})
	Fatal(msg interface{})
}

type commandExiter struct {
	out io.Writer
}

// Exit terminates the application.
func (exiter *commandExiter) Exit(code int) {
	os.Exit(code)
}

// Exitf prints a message and terminates the application.
func (exiter *commandExiter) Exitf(code int, format string, a ...interface{}) {
	if _, err := fmt.Fprintf(exiter.out, format, a...); err != nil {
		log.Fatal(err)
	}
	exiter.Exit(code)
}

// Err checks if the input is nil; if not it will exit via Fatal.
func (exiter *commandExiter) Err(err interface{}) {
	if err != nil {
		exiter.Fatal(err)
	}
}

// Fatal prints a message to console and terminates the application.
func (exiter *commandExiter) Fatal(msg interface{}) {
	exiter.Exitf(1, "Error: %s\n", msg)
}

// FromCmd creates a command exiter from the specified command, using the
// command's configured output as the exiter output writer.
func FromCmd(cmd *cobra.Command) Exiter {
	return &commandExiter{
		cmd.OutOrStderr(),
	}
}

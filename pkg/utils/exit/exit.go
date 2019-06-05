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

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
	"fmt"
	"io"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/vapor-ware/synse-cli/pkg/config"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

func init() {
	cmdCurrent.Flags().BoolVarP(&flagFull, "full", "f", false, "display the full context record")
	cmdCurrent.Flags().BoolVarP(&flagNoHeader, "no-header", "n", false, "do not print out column headers")
}

var cmdCurrent = &cobra.Command{
	Use:   "current [TYPE]",
	Short: "Display the current context",
	Long: utils.Doc(`
		Display the name of the current context(s).

		If no context is active, this command will result in an error.

		To get the current context for a specific Synse component, the
		TYPE may be specified. Valid types include:
		- <bold>server</>
		- <bold>plugin</>
	`),
	SuggestFor: []string{
		"active",
	},
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var ctxType string
		if len(args) != 0 {
			ctxType = args[0]
		}

		utils.Err(getCurrentContext(cmd.OutOrStdout(), ctxType))
	},
}

func getCurrentContext(out io.Writer, ctxType string) error {
	log.WithFields(log.Fields{
		"type": ctxType,
	}).Debug("getting current context")

	// Verify that the provided context type is supported.
	if ctxType != "" && ctxType != "plugin" && ctxType != "server" {
		return fmt.Errorf("unsupported context type: %s", ctxType)
	}

	currentContexts := config.GetCurrentContext()
	if len(currentContexts) == 0 {
		return fmt.Errorf("no current context is set (see 'synse context set')")
	}

	if ctxType != "" && currentContexts[ctxType] == nil {
		return fmt.Errorf("no current context is set for type '%s' (see 'synse context set')", ctxType)
	}

	w := utils.NewTabWriter(out)
	defer w.Flush()

	if !flagNoHeader {
		if err := printContextHeader(w, flagFull); err != nil {
			return err
		}
	}

	for _, ctx := range config.GetCurrentContext() {
		if ctxType != "" && ctx.Type != ctxType {
			continue
		}
		if err := printContext(w, ctx, flagFull); err != nil {
			return err
		}
	}
	return nil
}

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

package cmd

import (
	"fmt"
	"os"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/vapor-ware/synse-cli/pkg"
	"github.com/vapor-ware/synse-cli/pkg/templates"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

func init() {
	cmdVersion.Flags().BoolVarP(&flagSimple, "simple", "s", false, "display only the version number")
}

var flagSimple bool

var cmdVersion = &cobra.Command{
	Use:   "version",
	Short: "Display version information for synse",
	Long: utils.Doc(`
		Display the version and build-time information for the synse CLI binary.

		Version information for Synse components (e.g. server, plugin) can
		be printed from their corresponding sub-command.

		This command will print information in the format:

		  synse:
		   version     : canary
		   build date  : 2019-05-08T14:37:43
		   git commit  : 499e641
		   git tag     : 3.0.0-5-g499e641
		   go version  : go1.11.4
		   go compiler : gc
		   platform    : darwin/amd64

		- <bold>version</>: The semantic version of the CLI binary.
		- <bold>build date</>: The date and time at which the binary was built.
		- <bold>git commit</>: The commit at which the binary was built.
		- <bold>git tag</>: The tag at which the binary was built.
		- <bold>go version</>: The version of Go used to build the binary.
		- <bold>go compiler</>: The compiler used to build the binary.
		- <bold>platform</>: The operating system and architecture of the system used to
		    build the binary.

		More information about a specific release version can be found on the
		project's GitHub page: <underscore>https://github.com/vapor-ware/synse-cli/releases</>
	`),

	Run: func(cmd *cobra.Command, args []string) {
		v := pkg.GetVersion()

		if flagSimple {
			fmt.Println(v.Version)
			return
		}

		tmpl, err := template.New("version").Parse(templates.CmdVersionTemplate)
		utils.Err(err)

		err = tmpl.ExecuteTemplate(os.Stdout, "version", v)
		utils.Err(err)
	},
}

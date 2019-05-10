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
	"os"

	"github.com/spf13/cobra"
	"github.com/vapor-ware/synse-cli/pkg/utils"
	"github.com/vapor-ware/synse-cli/pkg/utils/exit"
)

var cmdCompletion = &cobra.Command{
	Use:   "completion",
	Short: "Generate bash completion scripts",
	Long: utils.Doc(`
		Generate bash completion scripts.

		To load bash completion for the current session, run:

		  <bold>. <(synse completion)</>

		To configure your bash shell to load synse completion for all
		new sessions, add the above to your bashrc, e.g.

		  <bold>echo ". <(synse completion)" >> ~/.bashrc</>
	`),

	Run: func(cmd *cobra.Command, args []string) {
		exit.FromCmd(cmd).Err(
			rootCmd.GenBashCompletion(os.Stdout),
		)
	},
}


const bash_completion_func = `

__synse_server_parse_scan() {
	local synse_server_output out
	if synse_server_output=$(synse server scan -n "$1" 2>/dev/null); then
		out=($(echo "${synse_server_output}" | awk '{print $1}'))
		COMPREPLY=( $( compgen -W "${out[*]}" -- "$cur" ) )
	fi
}

__synse_server_read_device() {
	if [[ ${#nouns[@]} -eq 0 ]]; then
		return 1
	fi
	__synse_server_parse_scan ${nouns[${#nouns[@]} -1]}
    if [[ $? -eq 0 ]]; then
        return 0
    fi
}

__synse_custom_func() {
	echo "calling custom func"
	case ${last_command} in 
		synse_server_read)
			__synse_server_read_device
			return
			;;
		*)
			;;
	esac
}
`
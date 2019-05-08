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

package plugin

import (
	"context"
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	"github.com/vapor-ware/synse-cli/pkg/utils"
	synse "github.com/vapor-ware/synse-server-grpc/go"
)

var cmdTest = &cobra.Command{
	Use:   "test",
	Short: "",
	Long:  heredoc.Doc(``),
	Run: func(cmd *cobra.Command, args []string) {
		utils.Err(testPlugin())
	},
}

func testPlugin() error {
	conn, client, err := utils.NewSynseGrpcClient()
	defer conn.Close()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	response, err := client.Test(ctx, &synse.Empty{})
	if err != nil {
		return err
	}

	// TODO: figure out response formatting
	fmt.Println(response)
	return nil
}

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

package server

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/gosuri/uilive"
	"github.com/spf13/cobra"
	"github.com/vapor-ware/synse-cli/pkg/utils"
	"github.com/vapor-ware/synse-cli/pkg/utils/exit"
	"github.com/vapor-ware/synse-client-go/synse/scheme"
	"golang.org/x/sync/errgroup"
)

func init() {
	cmdStream.Flags().StringSliceVarP(&flagDeviceIds, "id", "i", []string{}, "specify device IDs to use as selectors")
	cmdStream.Flags().StringSliceVarP(&flagTags, "tag", "t", []string{}, "specify tags to use as device selectors")
}

var cmdStream = &cobra.Command{
	Use:   "stream",
	Short: "Stream current reading data from Synse",
	Long: utils.Doc(`
		Get a live stream of reading data as it is read into the system.

		If no flags are specified, this will stream the readings for all devices.
		The '--id' and '--tag' flags can be used to narrow down the devices for which
		reading data is streamed back.

		Tags are strings with three components: a namespace (optional), an
		annotation (optional), and a label (required). They follow the format
		"namespace/annotation:label". Multiple tags can be specified either
		by calling the '--tag' flag multiple times, or by providing a comma
		separated list of tags. For example, the two lines below are equivalent:

		   --tag default/foo --tag default/type:bar
		   --tag default/foo,default/type:bar

		IDs are the device ID strings. Multiple IDs may be specified either by
		calling the '--id' flag multiple times, or by providing a comma separated
		list of IDs. For example, the following two lines are equivalent:

		  --id 285d23c1299ab8e5 --id 33aa9822b188d092fe
		  --id 285d23c1299ab8e5,33aa9822b188d092fe

		You cannot specify devices both by ID and tag. Doing so will result in
		an error.
	`),
	Run: func(cmd *cobra.Command, args []string) {
		exiter := exit.FromCmd(cmd)

		// Error out if device IDs and tag selectors are both specified.
		if len(flagDeviceIds) != 0 && len(flagTags) != 0 {
			exiter.Err("cannot specify device IDs and device tags together")
		}

		exiter.Err(serverStream(cmd.OutOrStdout()))
	},
}

func serverStream(out io.Writer) error {
	client, err := utils.NewSynseWebsocketClient(flagContext, flagTLSCert)
	if err != nil {
		return err
	}

	defer client.Close()
	if err := client.Open(); err != nil {
		return err
	}

	// Create a channel which will be used to stop the stream.
	stop := make(chan struct{}, 1)

	c := make(chan os.Signal, 1)

	go func() {
		<-c
		close(stop)
	}()

	// Create a channel which will be used to collect the readings as they come in.
	readings := make(chan *scheme.Read)
	defer close(readings)

	var g errgroup.Group
	g.Go(func() error {
		return client.ReadStream(
			scheme.ReadStreamOptions{
				Ids:  flagDeviceIds,
				Tags: flagTags,
			},
			readings,
			stop,
		)
	})

	writer := uilive.New()
	writer.Start()

	err = writer.Flush()
	if err != nil {
		return err
	}

	tw := utils.NewTabWriter(writer)
	defer tw.Flush()
	defer writer.Stop()

	header := []string{"ID", "TYPE", "VALUE", "UNIT", "TIMESTAMP"}
	rows := map[string]string{}

	go func() {
		for {
			select {
			case _, open := <-stop:
				if !open {
					return
				}
			default:
				// do nothing
			}

			str := fmt.Sprintf("%s\n", strings.Join(header, "\t"))
			var data []string
			for _, r := range rows {
				data = append(data, r)
			}
			sort.Strings(data)
			str += strings.Join(data, "\n")

			err := writer.Flush()
			if err != nil {
				fmt.Printf("error: %s", err)
				close(stop)
				return
			}

			err = tw.Flush()
			if err != nil {
				fmt.Printf("error: %s", err)
				close(stop)
				return
			}

			fmt.Fprint(out, "\033[2J")
			fmt.Fprint(out, "\033[H")

			_, err = fmt.Fprint(tw, str)
			if err != nil {
				fmt.Printf("error: %s", err)
				close(stop)
				return
			}

			time.Sleep(100 * time.Millisecond)
		}
	}()

readLoop:
	for reading := range readings {
		select {
		case _, open := <-stop:
			if !open {
				break readLoop
			}
		default:
			// do nothing and continue on
		}

		// Special casing for unit symbol.
		symbol := reading.Unit.Symbol
		if symbol == "%" {
			symbol = "%%"
		}

		row := strings.Join([]string{
			reading.Device,
			reading.Type,
			fmt.Sprintf("%v", reading.Value),
			symbol,
			reading.Timestamp,
		}, "\t")
		rows[fmt.Sprintf("%s-%s", reading.Device, reading.Type)] = row
	}

	return g.Wait()
}

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
	"testing"

	"bou.ke/monkey"
	"github.com/vapor-ware/synse-cli/internal/test"
	"github.com/vapor-ware/synse-cli/pkg/utils"
	"github.com/vapor-ware/synse-client-go/synse"
)

func TestCmdConfig_badClient(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseHTTPClient, func(ctx string, cert string) (synse.Client, error) {
		return nil, fmt.Errorf("test error message")
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdConfig).Run(t)
	result.AssertNoErr()
	result.AssertExited()
	result.AssertGolden("bad-client.golden")
}

func TestCmdConfig_requestError(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseHTTPClient, func(ctx string, cert string) (synse.Client, error) {
		return test.NewFakeHTTPClientV3Err(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdConfig).Run(t)
	result.AssertNoErr()
	result.AssertExited()
	result.AssertGolden("request-err.golden")
}

func TestCmdConfig_json(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseHTTPClient, func(ctx string, cert string) (synse.Client, error) {
		return test.NewFakeHTTPClientV3(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdConfig).Run(t)
	result.AssertNoErr()
	result.AssertGolden("config.json.golden")
}

func TestCmdConfig_yaml(t *testing.T) {
	patch := monkey.Patch(utils.NewSynseHTTPClient, func(ctx string, cert string) (synse.Client, error) {
		return test.NewFakeHTTPClientV3(), nil
	})
	defer patch.Unpatch()
	defer resetFlags()

	result := test.Cmd(cmdConfig).Args(
		"--yaml",
	).Run(t)
	result.AssertNoErr()
	result.AssertGolden("config.yaml.golden")
}

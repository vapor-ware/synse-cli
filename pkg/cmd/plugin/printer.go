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
	"fmt"

	synse "github.com/vapor-ware/synse-server-grpc/go"
)

func pluginTestRowFunc(data interface{}) ([]interface{}, error) {
	i, ok := data.(*synse.V3TestStatus)
	if !ok {
		return nil, fmt.Errorf("invalid row data")
	}

	var s = "OK"
	if !i.Ok {
		s = "ERROR"
	}

	return []interface{}{
		s,
	}, nil
}

func pluginVersionRowFunc(data interface{}) ([]interface{}, error) {
	i, ok := data.(*synse.V3Version)
	if !ok {
		return nil, fmt.Errorf("invalid row data")
	}

	return []interface{}{
		i.PluginVersion,
		i.SdkVersion,
		i.BuildDate,
		i.Os,
		i.Arch,
	}, nil
}

func pluginReadingRowFunc(data interface{}) ([]interface{}, error) {
	i, ok := data.(*synse.V3Reading)
	if !ok {
		return nil, fmt.Errorf("invalid row data")
	}

	// Special casing for reading unit symbol. % is a formatting
	// directive, so it needs to be escaped as a double percent.
	symbol := i.Unit.Symbol
	if symbol == "%" {
		symbol = "%%"
	}

	return []interface{}{
		i.Id,
		i.Value,
		symbol,
		i.Type,
		i.Timestamp,
	}, nil
}

func pluginTransactionStatusRowFunc(data interface{}) ([]interface{}, error) {
	i, ok := data.(*synse.V3TransactionStatus)
	if !ok {
		return nil, fmt.Errorf("invalid row data")
	}

	return []interface{}{
		i.Id,
		i.Status,
		i.Message,
		i.Created,
		i.Updated,
	}, nil
}

func pluginTransactionInfoRowFunc(data interface{}) ([]interface{}, error) {
	i, ok := data.(*synse.V3WriteTransaction)
	if !ok {
		return nil, fmt.Errorf("invalid row data")
	}

	return []interface{}{
		i.Id,
		i.Context.Action,
		i.Context.Data,
		i.Device,
	}, nil
}

func pluginMetadataRowFunc(data interface{}) ([]interface{}, error) {
	i, ok := data.(*synse.V3Metadata)
	if !ok {
		return nil, fmt.Errorf("invalid row data")
	}

	return []interface{}{
		i.Id,
		i.Tag,
		i.Description,
	}, nil
}

func pluginDeviceRowFunc(data interface{}) ([]interface{}, error) {
	i, ok := data.(*synse.V3Device)
	if !ok {
		return nil, fmt.Errorf("invalid row data")
	}

	return []interface{}{
		i.Id,
		i.Alias,
		i.Type,
		i.Info,
		i.Plugin,
	}, nil
}

func pluginHealthRowFunc(data interface{}) ([]interface{}, error) {
	i, ok := data.(*synse.V3Health)
	if !ok {
		return nil, fmt.Errorf("invalid row data")
	}

	return []interface{}{
		i.Status,
		i.Timestamp,
		len(i.Checks),
	}, nil
}

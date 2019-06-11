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
	"github.com/pkg/errors"
	"github.com/vapor-ware/synse-client-go/synse/scheme"
)

var (
	// ErrInvalidRowData is a printer function error which indicates that the
	// data type given to the printer is unexpected. This error should never be
	// induced by a user error, but may occur if there are changes in modeling.
	ErrInvalidRowData = errors.New("invalid row data")

	// ErrNilData is a printer function error which indicates that the value
	// passed to the printer is nil and can not be printed.
	ErrNilData = errors.New("row handler got nil data")
)

func serverReadRowFunc(data interface{}) ([]interface{}, error) {
	i, ok := data.(*scheme.Read)
	if !ok {
		return nil, ErrInvalidRowData
	}
	if i == nil {
		return nil, ErrNilData
	}

	// Special casing for reading unit symbol. % is a formatting
	// directive, so it needs to be escaped as a double percent.
	symbol := i.Unit.Symbol
	if symbol == "%" {
		symbol = "%%"
	}

	return []interface{}{
		i.Device,
		i.Value,
		symbol,
		i.Type,
		i.Timestamp,
	}, nil
}

func serverScanRowFunc(data interface{}) ([]interface{}, error) {
	i, ok := data.(*scheme.Scan)
	if !ok {
		return nil, ErrInvalidRowData
	}
	if i == nil {
		return nil, ErrNilData
	}

	return []interface{}{
		i.ID,
		i.Type,
		i.Info,
	}, nil
}

func serverStatusRowFunc(data interface{}) ([]interface{}, error) {
	i, ok := data.(*scheme.Status)
	if !ok {
		return nil, ErrInvalidRowData
	}
	if i == nil {
		return nil, ErrNilData
	}

	return []interface{}{
		i.Status,
		i.Timestamp,
	}, nil
}

func serverTagsRowFunc(data interface{}) ([]interface{}, error) {
	i, ok := data.(string)
	if !ok {
		return nil, ErrInvalidRowData
	}

	return []interface{}{
		i,
	}, nil
}

func serverTransactionRowFunc(data interface{}) ([]interface{}, error) {
	i, ok := data.(*scheme.Transaction)
	if !ok {
		return nil, ErrInvalidRowData
	}
	if i == nil {
		return nil, ErrNilData
	}

	return []interface{}{
		i.ID,
		i.Status,
		i.Message,
		i.Created,
		i.Updated,
	}, nil
}

func serverTransactionsRowFunc(data interface{}) ([]interface{}, error) {
	i, ok := data.(string)
	if !ok {
		return nil, ErrInvalidRowData
	}
	return []interface{}{
		i,
	}, nil
}

func serverTransactionSummaryRowFunc(data interface{}) ([]interface{}, error) {
	i, ok := data.(*scheme.Write)
	if !ok {
		return nil, ErrInvalidRowData
	}
	if i == nil {
		return nil, ErrNilData
	}

	return []interface{}{
		i.ID,
		i.Context.Action,
		i.Context.Data,
		i.Device,
	}, nil
}

func serverVersionRowFunc(data interface{}) ([]interface{}, error) {
	i, ok := data.(*scheme.Version)
	if !ok {
		return nil, ErrInvalidRowData
	}
	if i == nil {
		return nil, ErrNilData
	}

	return []interface{}{
		i.Version,
		i.APIVersion,
	}, nil
}

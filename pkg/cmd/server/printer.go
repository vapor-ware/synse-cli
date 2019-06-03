package server

import (
	"fmt"

	"github.com/vapor-ware/synse-client-go/synse/scheme"
)

func serverReadRowFunc(data interface{}) ([]interface{}, error) {
	i, ok := data.(*scheme.Read)
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
		return nil, fmt.Errorf("invalid row data")
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
		return nil, fmt.Errorf("invalid row data")
	}
	return []interface{}{
		i.Status,
		i.Timestamp,
	}, nil
}

func serverTagsRowFunc(data interface{}) ([]interface{}, error) {
	i, ok := data.(string)
	if !ok {
		return nil, fmt.Errorf("invalid row data")
	}
	return []interface{}{
		i,
	}, nil
}

func serverTransactionRowFunc(data interface{}) ([]interface{}, error) {
	i, ok := data.(*scheme.Transaction)
	if !ok {
		return nil, fmt.Errorf("invalid row data")
	}
	return []interface{}{
		i.ID,
		i.Status,
		i.Message,
		i.Created,
		i.Updated,
	}, nil
}

func serverTransactionSummaryRowFunc(data interface{}) ([]interface{}, error) {
	i, ok := data.(*scheme.Write)
	if !ok {
		return nil, fmt.Errorf("invalid row data")
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
		return nil, fmt.Errorf("invalid row data")
	}
	return []interface{}{
		i.Version,
		i.APIVersion,
	}, nil
}

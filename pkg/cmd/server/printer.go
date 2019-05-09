package server

import (
	"fmt"
	"io"
	"strings"

	"github.com/vapor-ware/synse-client-go/synse/scheme"
)

func printReadingHeader(out io.Writer) error {
	columns := []string{"ID", "VALUE", "UNIT", "TYPE", "TIMESTAMP"}

	_, err := fmt.Fprintf(out, "%s\n", strings.Join(columns, "\t"))
	return err
}

func printReadingRow(out io.Writer, read scheme.Read) error {
	// Special casing for reading unit symbol. % is a formatting
	// directive, so it needs to be escaped as a double percent.
	symbol := read.Unit.Symbol
	if symbol == "%" {
		symbol = "%%"
	}

	row := fmt.Sprintf("%s\t%v\t%s\t%s\t%s\n", read.Device, read.Value, symbol, read.Type, read.Timestamp)
	_, err := fmt.Fprintf(out, row)
	return err
}

func printVersionHeader(out io.Writer) error {
	columns := []string{"VERSION", "API_VERSION"}

	_, err := fmt.Fprintf(out, "%s\n", strings.Join(columns, "\t"))
	return err
}

func printVersionRow(out io.Writer, version *scheme.Version) error {
	row := fmt.Sprintf("%s\t%s\n", version.Version, version.APIVersion)
	_, err := fmt.Fprintf(out, row)
	return err
}

func printStatusHeader(out io.Writer) error {
	columns := []string{"STATUS", "TIMESTAMP"}

	_, err := fmt.Fprintf(out, "%s\n", strings.Join(columns, "\t"))
	return err
}

func printStatusRow(out io.Writer, status *scheme.Status) error {
	row := fmt.Sprintf("%s\t%s\n", status.Status, status.Timestamp)
	_, err := fmt.Fprintf(out, row)
	return err
}

func printScanHeader(out io.Writer, full bool) error {
	columns := []string{"ID", "TYPE", "INFO", "PLUGIN"}
	if !full {
		columns = columns[:3]
	}

	_, err := fmt.Fprintf(out, "%s\n", strings.Join(columns, "\t"))
	return err
}

func printScanRow(out io.Writer, scan *scheme.Scan, full bool) error {
	var row string
	if full {
		row = fmt.Sprintf("%s\t%s\t%s\t%s\n", scan.ID, scan.Type, scan.Info, scan.Plugin)
	} else {
		row = fmt.Sprintf("%s\t%s\t%s\n", scan.ID, scan.Type, scan.Info)
	}

	_, err := fmt.Fprintf(out, row)
	return err
}

func printTagsHeader(out io.Writer) error {
	columns := []string{"TAG"}

	_, err := fmt.Fprintf(out, "%s\n", strings.Join(columns, "\t"))
	return err
}

func printTagsRow(out io.Writer, tag string) error {
	row := fmt.Sprintf("%s\n", tag)
	_, err := fmt.Fprintf(out, row)
	return err
}

func printTransactionHeader(out io.Writer) error {
	columns := []string{"ID", "STATUS", "MESSAGE", "CREATED", "UPDATED"}

	_, err := fmt.Fprintf(out, "%s\n", strings.Join(columns, "\t"))
	return err
}

func printTransactionRow(out io.Writer, txn *scheme.Transaction) error {
	row := fmt.Sprintf("%s\t%s\t%s\t%s\t%s\n", txn.ID, txn.Status, txn.Message, txn.Created, txn.Updated)
	_, err := fmt.Fprintf(out, row)
	return err
}

func printTransactionSummaryHeader(out io.Writer) error {
	columns := []string{"ID", "ACTION", "DATA", "DEVICE"}

	_, err := fmt.Fprintf(out, "%s\n", strings.Join(columns, "\t"))
	return err
}

func printTransactionSummaryRow(out io.Writer, w *scheme.Write) error {
	row := fmt.Sprintf("%s\t%s\t%s\t%s\n", w.Transaction, w.Context.Action, w.Context.Data, w.Device)
	_, err := fmt.Fprintf(out, row)
	return err
}

package plugin

import (
	"fmt"
	"io"
	"strings"

	synse "github.com/vapor-ware/synse-server-grpc/go"
)

func printTestHeader(out io.Writer) error {
	columns := []string{"STATUS"}

	_, err := fmt.Fprintf(out, "%s\n", strings.Join(columns, "\t"))
	return err
}

func printTestRow(out io.Writer, status *synse.V3TestStatus) error {
	var s = "OK"
	if !status.Ok {
		s = "ERROR"
	}

	row := fmt.Sprintf("%s\n", s)
	_, err := fmt.Fprintf(out, row)
	return err
}

func printVersionHeader(out io.Writer) error {
	columns := []string{"VERSION", "SDK", "BUILD DATE", "OS", "ARCH"}

	_, err := fmt.Fprintf(out, "%s\n", strings.Join(columns, "\t"))
	return err
}

func printVersionRow(out io.Writer, v *synse.V3Version) error {
	row := fmt.Sprintf("%s\t%s\t%s\t%s\t%s\n", v.PluginVersion, v.SdkVersion, v.BuildDate, v.Os, v.Arch)
	_, err := fmt.Fprintf(out, row)
	return err
}

func printReadingHeader(out io.Writer) error {
	columns := []string{"ID", "VALUE", "UNIT", "TYPE", "TIMESTAMP"}

	_, err := fmt.Fprintf(out, "%s\n", strings.Join(columns, "\t"))
	return err
}

func printReadingRow(out io.Writer, r *synse.V3Reading) error {
	// Special casing for reading unit symbol. % is a formatting
	// directive, so it needs to be escaped as a double percent.
	symbol := r.Unit.Symbol
	if symbol == "%" {
		symbol = "%%"
	}

	row := fmt.Sprintf("%s\t%v\t%s\t%s\t%s\n", r.Id, r.Value, r.Unit.Symbol, r.Type, r.Timestamp)
	_, err := fmt.Fprintf(out, row)
	return err
}

func printTransactionStatusHeader(out io.Writer) error {
	columns := []string{"ID", "STATUS", "MESSAGE", "CREATED", "UPDATED"}

	_, err := fmt.Fprintf(out, "%s\n", strings.Join(columns, "\t"))
	return err
}

func printTransactionStatusRow(out io.Writer, s *synse.V3TransactionStatus) error {
	row := fmt.Sprintf("%s\t%s\t%s\t%s\t%s\n", s.Id, s.Status, s.Message, s.Created, s.Updated)
	_, err := fmt.Fprintf(out, row)
	return err
}

func printTransactionInfoHeader(out io.Writer) error {
	columns := []string{"TRANSACTION", "ACTION", "DATA", "DEVICE"}

	_, err := fmt.Fprintf(out, "%s\n", strings.Join(columns, "\t"))
	return err
}

func printTransactionInfoRow(out io.Writer, t *synse.V3WriteTransaction) error {
	row := fmt.Sprintf("%s\t%s\t%s\t%s\n", t.Id, t.Context.Action, t.Context.Data, t.Device)
	_, err := fmt.Fprintf(out, row)
	return err
}

func printMetadataHeader(out io.Writer) error {
	columns := []string{"ID", "TAG", "DESCRIPTION"}

	_, err := fmt.Fprintf(out, "%s\n", strings.Join(columns, "\t"))
	return err
}

func printMetadataRow(out io.Writer, m *synse.V3Metadata) error {
	row := fmt.Sprintf("%s\t%s\t%s\n", m.Id, m.Tag, m.Description)
	_, err := fmt.Fprintf(out, row)
	return err
}

func printDeviceHeader(out io.Writer) error {
	columns := []string{"ID", "ALIAS", "TYPE", "INFO", "PLUGIN"}

	_, err := fmt.Fprintf(out, "%s\n", strings.Join(columns, "\t"))
	return err
}

func printDeviceRow(out io.Writer, d *synse.V3Device) error {
	row := fmt.Sprintf("%s\t%s\t%s\t%s\t%s\n", d.Id, d.Alias, d.Type, d.Info, d.Plugin)
	_, err := fmt.Fprintf(out, row)
	return err
}

func printHealthHeader(out io.Writer) error {
	columns := []string{"STATUS", "TIMESTAMP", "CHECKS"}

	_, err := fmt.Fprintf(out, "%s\n", strings.Join(columns, "\t"))
	return err
}

func printHealthRow(out io.Writer, h *synse.V3Health) error {
	row := fmt.Sprintf("%s\t%s\t%d\n", h.Status, h.Timestamp, len(h.Checks))
	_, err := fmt.Fprintf(out, row)
	return err
}

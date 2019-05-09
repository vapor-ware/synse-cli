package plugins

import (
	"fmt"
	"io"
	"strings"

	"github.com/vapor-ware/synse-client-go/synse/scheme"
)

func printPluginSummaryHeader(out io.Writer) error {
	columns := []string{"ACTIVE", "ID", "VERSION", "TAG", "DESCRIPTION"}

	_, err := fmt.Fprintf(out, "%s\n", strings.Join(columns, "\t"))
	return err
}

func printPluginSummaryRow(out io.Writer, p *scheme.PluginMeta) error {
	var isActive = "✓"
	if !p.Active {
		isActive = "✗"
	}

	row := fmt.Sprintf("%s\t%s\t%s\t%s\t%s\n", isActive, p.ID, p.Version.PluginVersion, p.Tag, p.Description)
	_, err := fmt.Fprintf(out, row)
	return err
}

func printPluginHeader(out io.Writer) error {
	columns := []string{"ACTIVE", "ID", "TAG", "ADDRESS", "STATUS", "LAST_CHECK"}

	_, err := fmt.Fprintf(out, "%s\n", strings.Join(columns, "\t"))
	return err
}

func printPluginRow(out io.Writer, p *scheme.Plugin) error {
	var isActive = "✓"
	if !p.Active {
		isActive = "✗"
	}

	addr := fmt.Sprintf("%s://%s", p.Network.Protocol, p.Network.Address)

	row := fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%s\n", isActive, p.ID, p.Tag, addr, p.Health.Status, p.Health.Timestamp)
	_, err := fmt.Fprintf(out, row)
	return err
}

func printPluginHealthHeader(out io.Writer) error {
	columns := []string{"STATUS", "HEALTHY", "UNHEALTHY", "ACTIVE", "INACTIVE"}

	_, err := fmt.Fprintf(out, "%s\n", strings.Join(columns, "\t"))
	return err
}

func printPluginHealthRow(out io.Writer, p *scheme.PluginHealth) error {
	row := fmt.Sprintf("%s\t%d\t%d\t%d\t%d\n", p.Status, len(p.Healthy), len(p.Unhealthy), p.Active, p.Inactive)
	_, err := fmt.Fprintf(out, row)
	return err
}

package utils

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"
)

// TableOutput renders table output.
func TableOutput(header []string, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.SetBorder(false)
	table.SetColumnSeparator("")
	table.SetCenterSeparator("")
	table.SetAutoMergeCells(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.AppendBulk(data)
	table.Render()
}

// MinimalTableOutput renders a table with no borders, separators, or headers.
func MinimalTableOutput(data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(false)
	table.SetColumnSeparator("")
	table.AppendBulk(data)
	table.Render()
}

// ProgressBar renders an instance of the progress bar with the default configuration
// values.
func ProgressBar(length int, title string) (*mpb.Bar, *mpb.Progress) {
	length64 := int64(length)
	bar := mpb.New()
	progressBar := bar.AddBar((length64),
		mpb.AppendDecorators(
			decor.ETA(4, decor.DSyncSpace),
			decor.Percentage(5, 0),
		),
		mpb.PrependDecorators(
			decor.Name(title+":", len(title), decor.DwidthSync),
			decor.Elapsed(5, decor.DextraSpace),
		),
	)

	return progressBar, bar
}

// ProgressBarStop takes the progress bar object and terminates it wether
// rendering is complete or not.
func ProgressBarStop(pb *mpb.Progress) {
	pb.Stop()
}

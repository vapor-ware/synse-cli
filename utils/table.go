package utils

import (
	//"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/vbauerster/mpb"
)

func TableOutput(header []string, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.SetBorder(false)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.SetAutoMergeCells(false)
	table.AppendBulk(data)
	table.Render()
}

func ProgressBar(length int, title string) (*mpb.Bar, *mpb.Progress) {
	length64 := int64(length)
	bar := mpb.New()
	progressBar := bar.AddBar(length64)
	progressBar.AppendETA(4, mpb.DwidthSync|mpb.DextraSpace)
	progressBar.AppendPercentage(5, 0)
	progressBar.PrependName(title+":", len(title), mpb.DwidthSync)
	progressBar.PrependElapsed(5, mpb.DextraSpace)

	return progressBar, bar
}

func ProgressBarStop(pb *mpb.Progress) {
	pb.Stop()
}

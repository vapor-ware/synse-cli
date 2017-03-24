package utils

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/gosuri/uiprogress"
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

func ProgressBar(length int) *uiprogress.Bar {
	uiprogress.Start()
	progressBar:= uiprogress.AddBar(length)
	progressBar.AppendCompleted()
	progressBar.PrependElapsed()

	return progressBar
}

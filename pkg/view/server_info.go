package view

import (
	"fmt"
	"github.com/rivo/tview"
	"github.com/vapor-ware/synse-cli/pkg"
)

type ServerInfo struct {
	*tview.Table

	instance *Instance
	data     map[string]string
}

func NewServerInfo(i *Instance) *ServerInfo {
	return &ServerInfo{
		Table:    tview.NewTable(),
		instance: i,
		data:     map[string]string{},
	}
}

func (s *ServerInfo) Init() {
	s.SetBorderPadding(0, 0, 1, 0)
	s.loadData()
	s.layout()
}

func (s *ServerInfo) layout() {
	for row, name := range []string{"Address", "Plugin", "Server Version", "Client Version"} {
		s.SetCell(row, 0, tview.NewTableCell(fmt.Sprintf("%s:", name)))

		val, ok := s.data[name]
		if !ok {
			val = ""
		}
		s.SetCell(row, 1, tview.NewTableCell(val))
	}
}

func (s *ServerInfo) loadData() {
	s.data["Address"] = "localhost:5000"
	s.data["Plugin"] = s.instance.APIClient.PluginVersion()
	s.data["Server Version"] = s.instance.APIClient.ServerVersion()
	s.data["Client Version"] = pkg.Version
}

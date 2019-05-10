package plugins

import (
	"fmt"

	"github.com/vapor-ware/synse-client-go/synse/scheme"
)

func serverPluginSummaryRowFunc(data interface{}) ([]interface{}, error) {
	i, ok := data.(*scheme.PluginMeta)
	if !ok {
		return nil, fmt.Errorf("invalid row data")
	}

	var isActive = "✓"
	if !i.Active {
		isActive = "✗"
	}

	return []interface{}{
		isActive,
		i.ID,
		i.Version.PluginVersion,
		i.Tag,
		i.Description,
	}, nil
}

func serverPluginRowFunc(data interface{}) ([]interface{}, error) {
	i, ok := data.(*scheme.Plugin)
	if !ok {
		return nil, fmt.Errorf("invalid row data")
	}

	var isActive = "✓"
	if !i.Active {
		isActive = "✗"
	}

	addr := fmt.Sprintf("%s://%s", i.Network.Protocol, i.Network.Address)

	return []interface{}{
		isActive,
		i.ID,
		i.Tag,
		addr,
		i.Health.Status,
		i.Health.Timestamp,
	}, nil
}

func serverPluginHealthRowFunc(data interface{}) ([]interface{}, error) {
	i, ok := data.(*scheme.PluginHealth)
	if !ok {
		return nil, fmt.Errorf("invalid row data")
	}

	return []interface{}{
		i.Status,
		len(i.Healthy),
		len(i.Unhealthy),
		i.Active,
		i.Inactive,
	}, nil
}

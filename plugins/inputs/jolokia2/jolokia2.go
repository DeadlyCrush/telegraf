package jolokia2

import (
	"github.com/DeadlyCrush/telegraf"
	"github.com/DeadlyCrush/telegraf/plugins/inputs"
	"github.com/DeadlyCrush/telegraf/plugins/inputs/jolokia2/common"
	"github.com/DeadlyCrush/telegraf/plugins/inputs/jolokia2/jolokia2_agent"
	"github.com/DeadlyCrush/telegraf/plugins/inputs/jolokia2/jolokia2_proxy"
)

func init() {
	inputs.Add("jolokia2_agent", func() telegraf.Input {
		return &jolokia2_agent.JolokiaAgent{
			Metrics:               []common.MetricConfig{},
			DefaultFieldSeparator: ".",
		}
	})
	inputs.Add("jolokia2_proxy", func() telegraf.Input {
		return &jolokia2_proxy.JolokiaProxy{
			Metrics:               []common.MetricConfig{},
			DefaultFieldSeparator: ".",
		}
	})
}

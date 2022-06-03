//go:build !linux
// +build !linux

package synproxy

import (
	"github.com/DeadlyCrush/telegraf"
	"github.com/DeadlyCrush/telegraf/plugins/inputs"
)

func (k *Synproxy) Init() error {
	k.Log.Warn("Current platform is not supported")
	return nil
}

func (k *Synproxy) Gather(acc telegraf.Accumulator) error {
	return nil
}

func init() {
	inputs.Add("synproxy", func() telegraf.Input {
		return &Synproxy{}
	})
}

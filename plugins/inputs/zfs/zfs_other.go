//go:build !linux && !freebsd
// +build !linux,!freebsd

package zfs

import (
	"github.com/DeadlyCrush/telegraf"
	"github.com/DeadlyCrush/telegraf/plugins/inputs"
)

func (z *Zfs) Gather(acc telegraf.Accumulator) error {
	return nil
}

func init() {
	inputs.Add("zfs", func() telegraf.Input {
		return &Zfs{}
	})
}

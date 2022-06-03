//go:build !linux
// +build !linux

package cgroup

import (
	"github.com/DeadlyCrush/telegraf"
)

func (g *CGroup) Gather(acc telegraf.Accumulator) error {
	return nil
}

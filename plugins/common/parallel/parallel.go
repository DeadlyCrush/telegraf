package parallel

import "github.com/DeadlyCrush/telegraf"

type Parallel interface {
	Enqueue(telegraf.Metric)
	Stop()
}

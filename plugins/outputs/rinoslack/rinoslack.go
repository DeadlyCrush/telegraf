//go:generate ../../../tools/readme_config_includer/generator
package rinoslack

import (
	_ "embed"
	"os"

	"github.com/DeadlyCrush/telegraf"
	"github.com/DeadlyCrush/telegraf/plugins/outputs"
)

// DO NOT REMOVE THE NEXT TWO LINES! This is required to embed the sampleConfig data.
//go:embed sample.conf
var sampleConfig string

type Simple struct {
	Ok  bool            `toml:"ok"`
	Log telegraf.Logger `toml:"-"`
}

func (*Simple) SampleConfig() string {
	return sampleConfig
}

// Init is for setup, and validating config.
func (s *Simple) Init() error {
	return nil
}

func (s *Simple) Connect() error {
	// Make any connection required here
	return nil
}

func (s *Simple) Close() error {
	// Close any connections here.
	// Write will not be called once Close is called, so there is no need to synchronize.
	return nil
}

// Write should write immediately to the output, and not buffer writes
// (Telegraf manages the buffer for you). Returning an error will fail this
// batch of writes and the entire batch will be retried automatically.
func (s *Simple) Write(metrics []telegraf.Metric) error {
	for _, metric := range metrics {
		// write `metric` to the output sink here
		matic_string := metric.Time().Format("2006-01-02 15:04:05") + ": " + metric.Name() + "\n"
		filename := "C:\\rinopark\\repos\\telegraf\\test\\test.txt"
		f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			return err //panic(err)
		}

		defer f.Close()

		if _, err = f.WriteString(matic_string); err != nil {
			return err //panic(err)
		}
	}
	return nil
}

func init() {
	outputs.Add("rinoslack", func() telegraf.Output { return &Simple{} })
}

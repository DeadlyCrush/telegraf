package all

import (
	//Blank imports for plugins to register themselves
	_ "github.com/DeadlyCrush/telegraf/plugins/aggregators/basicstats"
	_ "github.com/DeadlyCrush/telegraf/plugins/aggregators/derivative"
	_ "github.com/DeadlyCrush/telegraf/plugins/aggregators/final"
	_ "github.com/DeadlyCrush/telegraf/plugins/aggregators/histogram"
	_ "github.com/DeadlyCrush/telegraf/plugins/aggregators/merge"
	_ "github.com/DeadlyCrush/telegraf/plugins/aggregators/minmax"
	_ "github.com/DeadlyCrush/telegraf/plugins/aggregators/quantile"
	_ "github.com/DeadlyCrush/telegraf/plugins/aggregators/starlark"
	_ "github.com/DeadlyCrush/telegraf/plugins/aggregators/valuecounter"
)

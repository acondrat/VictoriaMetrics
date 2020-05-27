package main

import (
	"context"

	"github.com/VictoriaMetrics/VictoriaMetrics/app/vmalert/datasource"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/prompbmarshal"
)

type Rule interface {
	ID() uint64
	Exec(context.Context, datasource.Querier, bool) ([]prompbmarshal.TimeSeries, error)
	UpdateWith(Rule) error
}

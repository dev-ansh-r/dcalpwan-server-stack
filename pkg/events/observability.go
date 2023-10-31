package events

import (
	"github.com/prometheus/client_golang/prometheus"
	"go.thethings.network/lorawan-stack/v3/pkg/metrics"
)

const subsystem = "events"

var publishes = metrics.NewCounterVec(
	prometheus.CounterOpts{
		Subsystem: subsystem,
		Name:      "publishes_total",
		Help:      "Number of Publishes to the default events PubSub",
	},
	[]string{"name"},
)

var channelDropped = metrics.NewCounterVec(
	prometheus.CounterOpts{
		Subsystem: subsystem,
		Name:      "channel_dropped_total",
		Help:      "Number of events dropped from event channels",
	},
	[]string{"name"},
)

func initMetrics(name string) {
	publishes.WithLabelValues(name).Add(0)
	channelDropped.WithLabelValues(name).Add(0)
}

func init() {
	metrics.MustRegister(publishes, channelDropped)
}

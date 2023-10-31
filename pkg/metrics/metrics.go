// Package metrics creates the metrics registry and exposes some common metrics.
package metrics

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"go.thethings.network/lorawan-stack/v3/pkg/version"
)

// Namespace for metrics.
const Namespace = "ttn_lw"

// ContextLabelNames are the label names that can be retrieved from a context for XXXVec metrics.
var ContextLabelNames []string

// LabelsFromContext returns the values for ContextLabelNames.
var LabelsFromContext func(ctx context.Context) prometheus.Labels

var ttnInfo = prometheus.NewGauge(prometheus.GaugeOpts{
	Namespace: Namespace,
	Name:      "info",
	Help:      "Information about The Things Stack for LoRaWAN",
	ConstLabels: prometheus.Labels{
		"version":    version.TTN,
		"build_date": version.BuildDate,
		"git_commit": version.GitCommit,
	},
})

func init() {
	ttnInfo.Set(1)
}

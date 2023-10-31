package metrics

import (
	"net/http"

	ocprom "contrib.go.opencensus.io/exporter/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"go.opencensus.io/stats/view"
)

var registry = prometheus.NewRegistry()

var exporter, _ = ocprom.NewExporter(ocprom.Options{
	Registry: registry,
})

// Exporter for the metrics registry.
var Exporter http.Handler = exporter

// Registry for metrics.
var Registry prometheus.Registerer = registry

func init() {
	registry.MustRegister(prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}))
	registry.MustRegister(prometheus.NewGoCollector())
	registry.MustRegister(ttnInfo)
	view.RegisterExporter(exporter)
}

// Register registers the given Collector in the registry.
func Register(c prometheus.Collector) error {
	return registry.Register(c)
}

// MustRegister registers the given Collectors in the registry and panics on errors.
func MustRegister(cs ...prometheus.Collector) {
	registry.MustRegister(cs...)
}

// Unregister the given Collector from the Prometheus registry.
func Unregister(c prometheus.Collector) bool {
	return registry.Unregister(c)
}

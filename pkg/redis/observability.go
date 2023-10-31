package redis

import (
	"net"

	"github.com/prometheus/client_golang/prometheus"
	"go.thethings.network/lorawan-stack/v3/pkg/metrics"
)

var (
	bytesRx = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: "redis_client",
			Name:      "receive_bytes_total",
			Help:      "Total number of bytes received by the Redis client",
		},
		[]string{"remote_address"},
	)
	bytesTx = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: "redis_client",
			Name:      "transmit_bytes_total",
			Help:      "Total number of bytes transmitted by the Redis client",
		},
		[]string{"remote_address"},
	)
	protosUnmarshaled = prometheus.NewCounter(
		prometheus.CounterOpts{
			Subsystem: "redis",
			Name:      "protos_unmarshaled_total",
			Help:      "Total number of protos unmarshaled by the Redis client",
		},
	)
	protosMarshaled = prometheus.NewCounter(
		prometheus.CounterOpts{
			Subsystem: "redis",
			Name:      "protos_marshaled_total",
			Help:      "Total number of protos marshaled by the Redis client",
		},
	)
)

func init() {
	metrics.MustRegister(
		bytesRx,
		bytesTx,
		protosUnmarshaled,
		protosMarshaled,
	)
}

type observableConn struct {
	addr string
	net.Conn
}

func (c *observableConn) Read(b []byte) (n int, err error) {
	n, err = c.Conn.Read(b)
	bytesRx.WithLabelValues(c.addr).Add(float64(n))
	return n, err
}

func (c *observableConn) Write(b []byte) (n int, err error) {
	n, err = c.Conn.Write(b)
	bytesTx.WithLabelValues(c.addr).Add(float64(n))
	return n, err
}

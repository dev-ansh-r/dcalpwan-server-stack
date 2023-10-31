package component

import (
	"crypto/tls"
	"io"
	stdlog "log"
	"net"
	"net/http"
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/config/tlsconfig"
	"go.thethings.network/lorawan-stack/v3/pkg/interop"
)

// RegisterInterop registers an interop subsystem to the component.
func (c *Component) RegisterInterop(s interop.Registerer) {
	c.interopSubsystems = append(c.interopSubsystems, s)
}

func (c *Component) serveInterop(lis net.Listener) error {
	srv := http.Server{
		Handler:           c.interop,
		ReadTimeout:       120 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		ErrorLog:          stdlog.New(io.Discard, "", 0),
	}
	go func() {
		<-c.Context().Done()
		srv.Close()
	}()
	return srv.Serve(lis)
}

func (c *Component) interopEndpoints() []Endpoint {
	return []Endpoint{
		NewTCPEndpoint(c.config.Interop.Listen, "Interop"),
		NewTLSEndpoint(c.config.Interop.ListenTLS, "Interop",
			tlsconfig.WithTLSClientAuth(tls.VerifyClientCertIfGiven, c.interop.ClientCAPool(), nil),
			tlsconfig.WithNextProtos("h2", "http/1.1"),
		),
	}
}

func (c *Component) listenInterop() error {
	return c.serveOnEndpoints(c.interopEndpoints(), (*Component).serveInterop, "interop")
}

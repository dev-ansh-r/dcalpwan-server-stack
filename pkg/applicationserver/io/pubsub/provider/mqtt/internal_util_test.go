package mqtt

import (
	"context"
	"crypto/tls"

	mqttnet "github.com/TheThingsIndustries/mystique/pkg/net"
	"github.com/TheThingsIndustries/mystique/pkg/server"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
)

var timeout = (1 << 8) * test.Delay

func startMQTTServer(ctx context.Context, tlsConfig *tls.Config) (mqttnet.Listener, mqttnet.Listener, error) {
	logger := log.FromContext(ctx)
	s := server.New(ctx)

	lis, err := mqttnet.Listen("tcp", ":0")
	if err != nil {
		return nil, nil, err
	}
	logger.Infof("Listening on %v", lis.Addr())
	go func() {
		for {
			conn, err := lis.Accept()
			if err != nil {
				logger.WithError(err).Error("Could not accept connection")
				return
			}
			go s.Handle(conn)
		}
	}()

	if tlsConfig != nil {
		tlsTCPLis, err := tls.Listen("tcp", ":0", tlsConfig)
		if err != nil {
			lis.Close()
			return nil, nil, err
		}
		tlsLis := mqttnet.NewListener(tlsTCPLis, "tls")
		logger.Infof("Listening on TLS %v", tlsLis.Addr())
		go func() {
			for {
				conn, err := tlsLis.Accept()
				if err != nil {
					logger.WithError(err).Error("Could not accept connection")
					return
				}
				go s.Handle(conn)
			}
		}()
		return lis, tlsLis, nil
	}

	return lis, nil, nil
}

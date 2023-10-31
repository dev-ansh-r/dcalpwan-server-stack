package mqtt

import (
	"crypto/tls"
	"crypto/x509"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
)

var errInvalidCAPEMData = errors.DefineInvalidArgument("ca_pem_data", "CA PEM data is invalid")

func createTLSConfig(caPEM []byte, certPEM []byte, keyPEM []byte) (*tls.Config, error) {
	// Change the CA certificate pool only if a CA has been provided.
	// This allows the system-wide CA pool to be used.
	var certPool *x509.CertPool
	if len(caPEM) != 0 {
		certPool = x509.NewCertPool()
		if !certPool.AppendCertsFromPEM(caPEM) {
			return nil, errInvalidCAPEMData.New()
		}
	}
	cert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		return nil, err
	}
	return &tls.Config{
		RootCAs:      certPool,
		Certificates: []tls.Certificate{cert},
	}, nil
}

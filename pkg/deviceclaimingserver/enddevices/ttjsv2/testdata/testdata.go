// NOTE: Set CAROOT=. when running go generate, i.e.:
// $ CAROOT=. go generate .

//go:generate mkcert -cert-file servercert.pem -key-file serverkey.pem localhost 127.0.0.1 ::1
//go:generate mkcert -cert-file clientcert-1.pem -key-file clientkey-1.pem -client client1.local
//go:generate mkcert -cert-file clientcert-2.pem -key-file clientkey-2.pem -client client2.local

// Package testdata provides test data.
package testdata

import (
	"crypto/x509"
	_ "embed"
	"encoding/pem"
)

var (
	//go:embed clientcert-1.pem
	client1CertData []byte
	//go:embed clientcert-2.pem
	client2CertData []byte
)

func x509Certificate(pemData []byte) *x509.Certificate {
	b, _ := pem.Decode(pemData)
	if b == nil || b.Type != "CERTIFICATE" {
		panic("invalid PEM data")
	}
	cert, err := x509.ParseCertificate(b.Bytes)
	if err != nil {
		panic(err)
	}
	return cert
}

// Client certificates.
var (
	Client1Cert = x509Certificate(client1CertData)
	Client2Cert = x509Certificate(client2CertData)
)

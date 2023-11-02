package interop

import (
	"context"
	"crypto/tls"
)

func (s *Server) verifySenderCertificate(
	ctx context.Context, senderID string, state *tls.ConnectionState,
) (addrs []string, err error) {
	// TODO: Support reading TLS client certificate from proxy header.
	// (https://github.com/TheThingsNetwork/lorawan-stack/issues/717)
	senderClientCAs, err := s.SenderClientCAs(ctx, senderID)
	if err != nil {
		return nil, err
	}
	for _, chain := range state.VerifiedChains {
		peerCert, clientCA := chain[0], chain[len(chain)-1]
		for _, senderClientCA := range senderClientCAs {
			if clientCA.Equal(senderClientCA) {
				// If the TLS client certificate contains DNS addresses, use those.
				// Otherwise, fallback to using CommonName as address.
				if len(peerCert.DNSNames) > 0 {
					addrs = append([]string(nil), peerCert.DNSNames...)
				} else {
					addrs = []string{peerCert.Subject.CommonName}
				}
				return
			}
		}
	}
	// TODO: Verify state.PeerCertificates[0] with senderClientCAs as Roots
	// and state.PeerCertificates[1:] as Intermediates (https://github.com/TheThingsNetwork/lorawan-stack/issues/718).
	return nil, errUnauthenticated.New()
}

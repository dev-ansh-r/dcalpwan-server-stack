package cluster

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"google.golang.org/grpc"
)

// Peer interface
type Peer interface {
	// Name of the peer
	Name() string
	// gRPC ClientConn to the peer (if available)
	Conn() (*grpc.ClientConn, error)
	// Roles announced by the peer
	Roles() []ttnpb.ClusterRole
	// HasRole returns true iff the peer has the given role
	HasRole(ttnpb.ClusterRole) bool
	// Tags announced by the peer
	Tags() map[string]string
}

type peer struct {
	name  string
	roles []ttnpb.ClusterRole
	tags  map[string]string

	target        string
	tlsServerName string

	ctx     context.Context
	cancel  context.CancelFunc
	conn    *grpc.ClientConn
	connErr error
}

func (p *peer) Name() string                    { return p.name }
func (p *peer) Conn() (*grpc.ClientConn, error) { return p.conn, p.connErr }
func (p *peer) Roles() []ttnpb.ClusterRole      { return p.roles }
func (p *peer) Tags() map[string]string         { return p.tags }

func (p *peer) HasRole(wanted ttnpb.ClusterRole) bool {
	roles := p.Roles()
	for _, role := range roles {
		if role == wanted {
			return true
		}
	}
	return false
}

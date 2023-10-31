package cryptoutil

import (
	"context"
	"net"
	"net/url"
	"strings"

	"go.thethings.network/lorawan-stack/v3/pkg/types"
)

// ComponentPrefixKEKLabeler is a ComponentKEKLabeler that joins the component prefix, separators and host.
type ComponentPrefixKEKLabeler struct {
	// Separator is the string to join parts.
	Separator string
	// ReplaceOldNew is a set of old and new string pairs to replace in parts.
	ReplaceOldNew []string
}

// hostFromAddress returns the cluster host from the given address.
func hostFromAddress(ctx context.Context, addr string) string {
	host := addr
	if u, err := url.Parse(addr); err == nil && u.Host != "" {
		host = u.Host
	}
	if h, _, err := net.SplitHostPort(host); err == nil {
		host = h
	}
	return host
}

func (c ComponentPrefixKEKLabeler) join(parts ...string) string {
	sep := c.Separator
	if sep == "" {
		sep = "/"
	}
	if len(c.ReplaceOldNew) > 0 {
		replacer := strings.NewReplacer(c.ReplaceOldNew...)
		parts = append(parts[:0:0], parts...)
		for i := range parts {
			parts[i] = replacer.Replace(parts[i])
		}
	}
	return strings.Join(parts, sep)
}

// NsKEKLabel returns a KEK label in the form `ns:netID:host` from the given NetID and address,
// where `:` is the default separator. Empty parts are omitted.
func (c ComponentPrefixKEKLabeler) NsKEKLabel(ctx context.Context, netID *types.NetID, addr string) string {
	parts := make([]string, 0, 3)
	parts = append(parts, "ns")
	if netID != nil {
		parts = append(parts, netID.String())
	}
	if addr != "" {
		parts = append(parts, hostFromAddress(ctx, addr))
	}
	return c.join(parts...)
}

// AsKEKLabel returns a KEK label in the form `as:host` from the given address,
// where `:` is the default separator. Empty parts are omitted.
func (c ComponentPrefixKEKLabeler) AsKEKLabel(ctx context.Context, addr string) string {
	parts := make([]string, 0, 2)
	parts = append(parts, "as")
	if addr != "" {
		parts = append(parts, hostFromAddress(ctx, addr))
	}
	return c.join(parts...)
}

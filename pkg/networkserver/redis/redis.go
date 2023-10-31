// Package redis provides Redis implementations of interfaces used by networkserver.
package redis

import (
	"encoding/base64"
	"hash/fnv"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
)

type keyer interface {
	Key(ks ...string) string
}

func UIDKey(r keyer, uid string) string {
	return r.Key("uid", uid)
}

var keyEncoding = base64.RawStdEncoding

func uplinkPayloadHash(b []byte) string {
	h := fnv.New64a()
	_, _ = h.Write(b)
	return keyEncoding.EncodeToString(h.Sum(nil))
}

var errDatabaseCorruption = errors.DefineCorruption("database_corruption", "database is corrupted")

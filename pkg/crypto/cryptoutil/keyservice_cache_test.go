package cryptoutil_test

import (
	"context"
	"crypto/tls"
	"errors"
	"testing"

	"github.com/smarty/assertions"
	. "go.thethings.network/lorawan-stack/v3/pkg/crypto/cryptoutil"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

type mockKeyService struct {
	params struct {
		ctx        context.Context
		ciphertext []byte
		kekLabel   string
	}
	results struct {
		key []byte
		err error
	}
	calls struct {
		count int
	}
}

func (*mockKeyService) Wrap(context.Context, []byte, string) ([]byte, error) {
	return nil, errors.New("not implemented")
}

func (m *mockKeyService) Unwrap(ctx context.Context, ciphertext []byte, label string) ([]byte, error) {
	m.params.ctx, m.params.ciphertext, m.params.kekLabel = ctx, ciphertext, label
	m.calls.count++
	return m.results.key, m.results.err
}

func (*mockKeyService) Encrypt(context.Context, []byte, string) ([]byte, error) {
	return nil, errors.New("not implemented")
}

func (*mockKeyService) Decrypt(context.Context, []byte, string) ([]byte, error) {
	return nil, errors.New("not implemented")
}

func (*mockKeyService) ServerCertificate(context.Context, string) (tls.Certificate, error) {
	return tls.Certificate{}, errors.New("not implemented")
}

func (*mockKeyService) ClientCertificate(context.Context, string) (tls.Certificate, error) {
	return tls.Certificate{}, errors.New("not implemented")
}

func (*mockKeyService) HMACHash(context.Context, []byte, string) ([]byte, error) {
	return nil, errors.New("not implemented")
}

func TestCacheUsed(t *testing.T) {
	t.Parallel()
	a := assertions.New(t)
	m := &mockKeyService{}

	ctx := test.Context()
	ck := NewCacheKeyService(m, test.Delay, 1)

	// Cache is empty, expect a miss
	m.results.key = []byte{0x01, 0x02, 0x03}
	m.results.err = nil

	key, err := ck.Unwrap(ctx, []byte{0x02, 0x03, 0x04}, "foo")
	a.So(key, should.Resemble, m.results.key)
	a.So(err, should.Equal, m.results.err)
	a.So(m.params.ctx, should.HaveParentContextOrEqual, ctx)
	a.So(m.params.ciphertext, should.Resemble, []byte{0x02, 0x03, 0x04})
	a.So(m.params.kekLabel, should.Equal, "foo")
	a.So(m.calls.count, should.Equal, 1)

	// Expect to be served from cache
	key, err = ck.Unwrap(ctx, []byte{0x02, 0x03, 0x04}, "foo")
	a.So(key, should.Resemble, m.results.key)
	a.So(err, should.Equal, m.results.err)
	a.So(m.calls.count, should.Equal, 1)

	// Expect a cache miss with the same KEK label but different ciphertext
	m.results.key = []byte{0x04, 0x05, 0x06}
	m.results.err = nil
	key, err = ck.Unwrap(ctx, []byte{0x05, 0x06, 0x07}, "foo")

	a.So(key, should.Resemble, m.results.key)
	a.So(err, should.Equal, m.results.err)
	a.So(m.calls.count, should.Equal, 2)

	// Expect the old element to be evicted, and a cache miss to occur
	m.results.key = []byte{0x05, 0x06, 0x07}
	m.results.err = nil
	m.calls.count = 0

	key, err = ck.Unwrap(ctx, []byte{0x03, 0x04, 0x05}, "bar")
	a.So(key, should.Resemble, m.results.key)
	a.So(err, should.Equal, m.results.err)
	a.So(m.params.ctx, should.HaveParentContextOrEqual, ctx)
	a.So(m.params.ciphertext, should.Resemble, []byte{0x03, 0x04, 0x05})
	a.So(m.params.kekLabel, should.Equal, "bar")
	a.So(m.calls.count, should.Equal, 1)

	// Expect the cache miss
	m.results.key = []byte{0x01, 0x02, 0x03}
	m.results.err = nil
	m.calls.count = 0

	key, err = ck.Unwrap(ctx, []byte{0x02, 0x03, 0x04}, "foo")
	a.So(key, should.Resemble, m.results.key)
	a.So(err, should.Equal, m.results.err)
	a.So(m.params.ctx, should.HaveParentContextOrEqual, ctx)
	a.So(m.params.ciphertext, should.Resemble, []byte{0x02, 0x03, 0x04})
	a.So(m.params.kekLabel, should.Equal, "foo")
	a.So(m.calls.count, should.Equal, 1)

	// Delay based evictions are left out since they may lead to flakyness.
}

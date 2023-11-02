package cryptoutil_test

import (
	"testing"

	"github.com/smarty/assertions"

	"go.thethings.network/lorawan-stack/v3/pkg/crypto/cryptoutil"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestEmptyKeyVault(t *testing.T) {
	t.Parallel()
	a := assertions.New(t)

	_, err := cryptoutil.EmptyKeyVault.Key(test.Context(), "test")
	a.So(errors.IsNotFound(err), should.BeTrue)

	_, err = cryptoutil.EmptyKeyVault.ServerCertificate(test.Context(), "test")
	a.So(errors.IsNotFound(err), should.BeTrue)

	_, err = cryptoutil.EmptyKeyVault.ClientCertificate(test.Context(), "test")
	a.So(errors.IsNotFound(err), should.BeTrue)
}

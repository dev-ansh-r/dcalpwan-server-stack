package identityserver

import (
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	bunstore "go.thethings.network/lorawan-stack/v3/pkg/identityserver/bunstore"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	storeutil "go.thethings.network/lorawan-stack/v3/pkg/util/store"
)

func (is *IdentityServer) setupStore() error {
	return is.setupBunStore()
}

func (is *IdentityServer) setupBunStore() (err error) {
	is.db, err = storeutil.OpenDB(is.Context(), is.config.DatabaseURI)
	if err != nil {
		return err
	}
	bunDB := bun.NewDB(is.db, pgdialect.New())
	if is.LogDebug() {
		bunDB.AddQueryHook(storeutil.NewLoggerHook(log.FromContext(is.Context()).WithField("namespace", "db")))
	}
	bunStore, err := bunstore.NewStore(is.Context(), bunDB)
	if err != nil {
		return err
	}
	is.store = bunStore

	return nil
}

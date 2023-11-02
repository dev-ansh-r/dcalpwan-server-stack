package store

import (
	"context"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
	ismigrations "go.thethings.network/lorawan-stack/v3/pkg/identityserver/store/migrations"
)

// Migrate migrates the database.
func Migrate(ctx context.Context, db *bun.DB) error {
	migrator := migrate.NewMigrator(db, ismigrations.Migrations)
	err := migrator.Init(ctx)
	if err != nil {
		return err
	}
	_, err = migrator.Migrate(ctx)
	return err
}

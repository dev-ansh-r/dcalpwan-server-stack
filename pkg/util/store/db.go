package store

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/uptrace/bun/driver/pgdriver"
	"go.thethings.network/lorawan-stack/v3/pkg/experimental"
)

var pgdriverFeatureFlag = experimental.DefineFeature("db.pgdriver", false)

// OpenDB opens the database connection.
func OpenDB(ctx context.Context, databaseURI string) (*sql.DB, error) {
	if pgdriverFeatureFlag.GetValue(ctx) {
		return sql.OpenDB(pgdriver.NewConnector(
			pgdriver.WithDSN(databaseURI),
			pgdriver.WithResetSessionFunc(func(ctx context.Context, cn *pgdriver.Conn) error {
				return checkConn(cn.Conn())
			}),
		)), nil
	}
	config, err := pgx.ParseConfig(databaseURI)
	if err != nil {
		return nil, err
	}
	config.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol
	return stdlib.OpenDB(*config), nil
}

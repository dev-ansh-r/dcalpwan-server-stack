// Package migrations contains Identity Server store migrations.
package migrations

import (
	"embed"

	"github.com/uptrace/bun/migrate"
)

// Migrations is the collection of schema migrations.
var Migrations = migrate.NewMigrations()

//go:embed *.sql
var sqlMigrations embed.FS

func init() {
	if err := Migrations.Discover(sqlMigrations); err != nil {
		panic(err)
	}
}

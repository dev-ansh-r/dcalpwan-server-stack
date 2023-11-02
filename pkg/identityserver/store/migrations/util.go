package migrations

import (
	"context"

	"github.com/uptrace/bun"
)

func tableExists(ctx context.Context, db *bun.DB, tableName string) (bool, error) {
	c, err := db.NewSelect().
		TableExpr("INFORMATION_SCHEMA.tables").
		Where("table_name = ?", tableName).
		Where("table_type = 'BASE TABLE'").
		Where("table_schema = CURRENT_SCHEMA()").
		Count(ctx)
	if err != nil {
		return false, err
	}
	return c == 1, nil
}

package store

import (
	"context"

	"github.com/uptrace/bun"
)

func selectWithContext(ctx context.Context) func(*bun.SelectQuery) *bun.SelectQuery {
	return func(q *bun.SelectQuery) *bun.SelectQuery {
		return q
	}
}

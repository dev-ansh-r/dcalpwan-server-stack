package redis

import (
	"context"
)

type paginationOptionsKeyType struct{}

var paginationOptionsKey paginationOptionsKeyType

type paginationOptions struct {
	limit  int64
	offset int64
	total  *int64
}

// NewContextWithPagination instructs the store to paginate the results.
func NewContextWithPagination(ctx context.Context, limit, page int64, total *int64) context.Context {
	if page == 0 {
		page = 1
	}
	return context.WithValue(ctx, paginationOptionsKey, paginationOptions{
		limit:  limit,
		offset: (page - 1) * limit,
		total:  total,
	})
}

// SetPaginationTotal sets the total number of results inside the paginated context, if it was not set already.
func SetPaginationTotal(ctx context.Context, total int64) {
	if opts, ok := ctx.Value(paginationOptionsKey).(paginationOptions); ok && opts.total != nil && *opts.total == 0 {
		*opts.total = total
	}
}

// PaginationLimitAndOffsetFromContext returns the pagination limit and the offset if they are present.
func PaginationLimitAndOffsetFromContext(ctx context.Context) (limit, offset int64) {
	if opts, ok := ctx.Value(paginationOptionsKey).(paginationOptions); ok {
		return opts.limit, opts.offset
	}
	return
}

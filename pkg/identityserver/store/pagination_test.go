package store

import (
	"fmt"
	"testing"

	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestPagination(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		limit          uint32
		page           uint32
		expectedLimit  uint64
		expectedOffset uint64
	}{
		{
			limit:          0,
			page:           0,
			expectedLimit:  0,
			expectedOffset: 0,
		},
		{
			limit:          10,
			page:           0,
			expectedLimit:  10,
			expectedOffset: 0,
		},
		{
			limit:          10,
			page:           1,
			expectedLimit:  10,
			expectedOffset: 0,
		},
		{
			limit:          10,
			page:           2,
			expectedLimit:  10,
			expectedOffset: 10,
		},
	} {
		t.Run(fmt.Sprintf("limitAndOffsetFromContext, limit:%v, offset:%v", tc.limit, tc.page), func(t *testing.T) {
			t.Parallel()

			a, ctx := test.New(t)

			limit, offset := LimitAndOffsetFromContext(WithPagination(ctx, tc.limit, tc.page, nil))

			a.So(limit, should.Equal, tc.expectedLimit)
			a.So(offset, should.Equal, tc.expectedOffset)
		})
	}

	t.Run("SetTotalCount", func(t *testing.T) {
		a, ctx := test.New(t)

		var totalCount uint64
		total := uint64(10)

		SetTotal(ctx, total)
		a.So(totalCount, should.BeZeroValue)

		ctx = WithPagination(ctx, 5, 1, &totalCount)
		a.So(totalCount, should.BeZeroValue)

		SetTotal(ctx, total)
		a.So(totalCount, should.Equal, total)
	})
}

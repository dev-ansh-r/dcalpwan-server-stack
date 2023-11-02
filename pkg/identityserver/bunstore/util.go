package store

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/uptrace/bun"
	"go.thethings.network/lorawan-stack/v3/pkg/identityserver/store"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
)

const timePrecision = time.Microsecond

func cleanTime(t time.Time) time.Time {
	return t.Truncate(timePrecision)
}

func cleanTimePtr(t *time.Time) *time.Time {
	if t == nil {
		return nil
	}
	clean := cleanTime(*t)
	return &clean
}

var nowFunc func() time.Time

func now() time.Time {
	if nowFunc != nil {
		return cleanTime(nowFunc())
	}
	return cleanTime(time.Now())
}

func combineApply[Q bun.Query](f ...func(Q) Q) func(Q) Q {
	return func(q Q) Q {
		for _, f := range f {
			q = f(q)
		}
		return q
	}
}

func mapSlice[X any, Y any](in []X, f func(X) Y) []Y {
	out := make([]Y, len(in))
	for i, x := range in {
		out[i] = f(x)
	}
	return out
}

func equalTime(a, b *time.Time) bool {
	if a == nil {
		return b == nil
	}
	if b == nil {
		return false
	}
	return a.Equal(*b)
}

func idStrings[X ttnpb.IDStringer](in ...X) []string {
	out := make([]string, len(in))
	for i, x := range in {
		out[i] = x.IDString()
	}
	return out
}

func noopSelectModifier(q *bun.SelectQuery) *bun.SelectQuery { return q }

func selectWithLimitAndOffsetFromContext(ctx context.Context) func(*bun.SelectQuery) *bun.SelectQuery {
	limit, offset := store.LimitAndOffsetFromContext(ctx)
	if limit == 0 {
		return noopSelectModifier
	}
	return func(q *bun.SelectQuery) *bun.SelectQuery {
		return q.Limit(int(limit)).Offset(int(offset))
	}
}

func selectWithOrderFromContext(
	ctx context.Context, defaultColumn string, fieldToColumn map[string]string,
) func(*bun.SelectQuery) *bun.SelectQuery {
	order := store.OrderOptionsFromContext(ctx)
	if column, ok := fieldToColumn[order.Field]; ok {
		return func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order(
				fmt.Sprintf("%s %s", column, order.Direction),
				fmt.Sprintf("%s %s", defaultColumn, order.Direction),
			)
		}
	}
	return func(q *bun.SelectQuery) *bun.SelectQuery {
		return q.Order(fmt.Sprintf("%s ASC", defaultColumn))
	}
}

func convertIntSlice[A, B int | ~int32](in []A) []B {
	out := make([]B, len(in))
	for i, el := range in {
		out[i] = B(el)
	}
	return out
}

func secretFromBytes(b []byte) *ttnpb.Secret {
	if len(b) == 0 {
		return nil
	}
	blocks := bytes.SplitN(b, []byte(":"), 2)
	if len(blocks) == 2 {
		return &ttnpb.Secret{
			KeyId: string(blocks[0]),
			Value: blocks[1],
		}
	}
	return nil
}

func secretToBytes(secret *ttnpb.Secret) []byte {
	if secret == nil {
		return nil
	}
	var buf bytes.Buffer
	buf.WriteString(secret.KeyId)
	buf.WriteByte(':')
	buf.Write(secret.Value)
	return buf.Bytes()
}

func eui64FromString(s *string) *types.EUI64 {
	if s == nil {
		return nil
	}
	var eui types.EUI64
	if err := eui.UnmarshalText([]byte(*s)); err != nil {
		return nil
	}
	return &eui
}

func eui64ToString(eui *types.EUI64) *string {
	if eui == nil {
		return nil
	}
	s := eui.String()
	return &s
}

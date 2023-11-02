package store

import (
	"testing"

	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestFieldMask_Contains(t *testing.T) {
	t.Parallel()
	a, _ := test.New(t)

	a.So(FieldMask{}.Contains(""), should.BeFalse)
	a.So(FieldMask{"a", "b", "c"}.Contains("a"), should.BeTrue)
}

func TestFieldMask_TopLevel(t *testing.T) {
	t.Parallel()
	a, _ := test.New(t)

	a.So(FieldMask{}.TopLevel(), should.BeEmpty)
	a.So(FieldMask{"a.foo", "a.bar", "b", "c"}.TopLevel(), should.Resemble, FieldMask{"a", "b", "c"})
}

func TestFieldMask_TrimPrefix(t *testing.T) {
	t.Parallel()
	a, _ := test.New(t)

	a.So(FieldMask{}.TrimPrefix("a"), should.BeEmpty)
	a.So(FieldMask{"a.foo", "a.bar", "b", "c"}.TrimPrefix("a"), should.Resemble, FieldMask{"foo", "bar"})
}

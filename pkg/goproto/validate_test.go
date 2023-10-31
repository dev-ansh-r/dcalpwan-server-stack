package goproto_test

import (
	"math"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/goproto"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
	"google.golang.org/protobuf/types/known/structpb"
)

func TestValidateStruct(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name string

		st       map[string]any
		expected []string
	}{
		{
			name: "empty",
		},
		{
			name: "simple object",

			st: map[string]any{
				"foo": 123,
			},
		},
		{
			name: "top level NaN",

			st: map[string]any{
				"foo": math.NaN(),
			},
			expected: []string{
				".foo: invalid NaN value",
			},
		},
		{
			name: "nested object Infinity",

			st: map[string]any{
				"foo": map[string]any{
					"bar": math.Inf(1),
				},
			},
			expected: []string{
				".foo.bar: invalid Infinity value",
			},
		},
		{
			name: "nested object -Infinity",
			st: map[string]any{
				"foo": []any{
					math.Inf(-1),
				},
			},
			expected: []string{
				".foo[0]: invalid -Infinity value",
			},
		},
		{
			name: "nested object NaN",
			st: map[string]any{
				"foo": []any{
					map[string]any{
						"bar": math.NaN(),
					},
				},
			},
			expected: []string{
				".foo[0].bar: invalid NaN value",
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			a := assertions.New(t)

			st, err := structpb.NewStruct(tc.st)
			if !a.So(err, should.BeNil) {
				t.FailNow()
			}

			warnings := goproto.ValidateStruct(st)
			a.So(warnings, should.Resemble, tc.expected)
		})
	}
}

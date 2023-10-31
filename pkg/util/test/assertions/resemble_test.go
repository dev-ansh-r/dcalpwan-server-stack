package assertions_test

import (
	"fmt"
	"testing"

	"github.com/smarty/assertions"
	. "go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestHaveEmptyDiff(t *testing.T) {
	for _, tc := range []struct {
		A                any
		B                any
		Assertion        func(any, ...any) string
		InverseAssertion func(any, ...any) string
	}{
		{
			A:                "test",
			B:                "test",
			Assertion:        should.BeEmpty,
			InverseAssertion: should.NotBeEmpty,
		},
		{
			A:                "test",
			B:                "test1",
			Assertion:        should.NotBeEmpty,
			InverseAssertion: should.BeEmpty,
		},
		{
			A:                42,
			B:                42,
			Assertion:        should.BeEmpty,
			InverseAssertion: should.NotBeEmpty,
		},
		{
			A:                1,
			B:                2,
			Assertion:        should.NotBeEmpty,
			InverseAssertion: should.BeEmpty,
		},
		{
			A: struct {
				Foo int
				Bar int
			}{42, 43},
			B: struct {
				Foo int
				Bar int
			}{42, 43},
			Assertion:        should.BeEmpty,
			InverseAssertion: should.NotBeEmpty,
		},
		{
			A:                nil,
			B:                0,
			Assertion:        should.NotBeEmpty,
			InverseAssertion: should.BeEmpty,
		},
		{
			A:                nil,
			B:                "test",
			Assertion:        should.NotBeEmpty,
			InverseAssertion: should.BeEmpty,
		},
		{
			A:                []string{},
			B:                []string(nil),
			Assertion:        should.BeEmpty,
			InverseAssertion: should.NotBeEmpty,
		},
		{
			A:                map[int]int{},
			B:                map[int]int(nil),
			Assertion:        should.BeEmpty,
			InverseAssertion: should.NotBeEmpty,
		},
	} {
		t.Run(fmt.Sprintf("%v/%v", tc.A, tc.B), func(t *testing.T) {
			a := assertions.New(t)
			a.So(ShouldHaveEmptyDiff(tc.A, tc.B), tc.Assertion)
			a.So(ShouldNotHaveEmptyDiff(tc.A, tc.B), tc.InverseAssertion)
		})
	}
}

type testSetFielder struct {
	A int
	B string
	C []string
}

func (dst *testSetFielder) SetFields(src *testSetFielder, paths ...string) error {
	for _, p := range paths {
		switch p {
		case "a":
			dst.A = src.A
		case "b":
			dst.B = src.B
		case "c":
			dst.C = src.C
		default:
			return fmt.Errorf("invalid path '%s'", p)
		}
	}
	return nil
}

func TestShouldResembleFields(t *testing.T) {
	for _, tc := range []struct {
		A         any
		B         any
		Paths     []any
		Assertion func(any, ...any) string
	}{
		{
			A:         &testSetFielder{},
			B:         &testSetFielder{},
			Assertion: should.BeEmpty,
		},
		{
			A: &testSetFielder{
				A: 42,
			},
			B:         &testSetFielder{},
			Assertion: should.BeEmpty,
		},
		{
			A: &testSetFielder{
				A: 42,
			},
			B: &testSetFielder{
				A: 42,
			},
			Assertion: should.BeEmpty,
		},
		{
			A: &testSetFielder{
				A: 42,
			},
			B: &testSetFielder{
				A: 42,
			},
			Paths:     []any{"a"},
			Assertion: should.BeEmpty,
		},
		{
			A: &testSetFielder{
				A: 42,
			},
			B: &testSetFielder{
				A: 42,
			},
			Paths:     []any{[]string{"a"}},
			Assertion: should.BeEmpty,
		},
		{
			A: &testSetFielder{
				A: 42,
			},
			B: &testSetFielder{
				A: 42,
			},
			Paths:     []any{[1]string{"a"}},
			Assertion: should.BeEmpty,
		},
		{
			A: &testSetFielder{
				A: 42,
			},
			B: &testSetFielder{
				A: 42,
			},
			Paths:     []any{"b"},
			Assertion: should.BeEmpty,
		},
		{
			A: &testSetFielder{
				A: 42,
			},
			B: &testSetFielder{
				A: 42,
			},
			Paths:     []any{"a", "b"},
			Assertion: should.BeEmpty,
		},
		{
			A: &testSetFielder{
				A: 42,
			},
			B: &testSetFielder{
				A: 43,
			},
			Paths:     []any{[]string{"b"}, "c"},
			Assertion: should.BeEmpty,
		},
		{
			A: &testSetFielder{
				A: 42,
			},
			B: &testSetFielder{
				A: 43,
			},
			Paths:     []any{"a", "b"},
			Assertion: should.NotBeEmpty,
		},
		{
			A: &testSetFielder{
				A: 42,
			},
			B: &testSetFielder{
				A: 43,
			},
			Paths:     []any{[2]string{"a", "a"}, []string{"b"}, "c"},
			Assertion: should.NotBeEmpty,
		},
	} {
		t.Run(fmt.Sprintf("%+v/%+v/%+v", tc.A, tc.B, tc.Paths), func(t *testing.T) {
			a := assertions.New(t)
			a.So(ShouldResembleFields(tc.A, append([]any{tc.B}, tc.Paths...)...), tc.Assertion)
		})
	}
}

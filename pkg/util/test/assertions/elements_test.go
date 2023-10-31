package assertions_test

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	. "go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestShouldHaveSameElements(t *testing.T) {
	for i, tc := range []struct {
		A             any
		B             any
		ShouldFunc    func(actual any, expected ...any) string
		ShouldNotFunc func(actual any, expected ...any) string
	}{
		{
			[][]byte{{42}, {43}},
			[][]byte{{43}, {44}},
			should.NotBeEmpty,
			should.BeEmpty,
		},
		{
			[][]byte{{43}, {43}},
			[][]byte{{43}, {44}},
			should.NotBeEmpty,
			should.BeEmpty,
		},
		{
			[][]byte{{43}, {43}, {43}},
			[][]byte{{43}, {44}},
			should.NotBeEmpty,
			should.BeEmpty,
		},
		{
			[][]byte{{42}, {43}, {43}},
			[][]byte{{43}, {42}, {43}},
			should.BeEmpty,
			should.NotBeEmpty,
		},
		{
			[][]byte{},
			[][]byte{{43}, {42}, {43}},
			should.NotBeEmpty,
			should.BeEmpty,
		},
		{
			[][]byte{{43}, {42}, {43}},
			[][]byte{},
			should.NotBeEmpty,
			should.BeEmpty,
		},
		{
			[]int{42},
			[]int{42},
			should.BeEmpty,
			should.NotBeEmpty,
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			a := assertions.New(t)

			a.So(ShouldHaveSameElementsFunc(tc.A, reflect.DeepEqual, tc.B), tc.ShouldFunc)
			a.So(ShouldNotHaveSameElementsFunc(tc.A, reflect.DeepEqual, tc.B), tc.ShouldNotFunc)
			a.So(ShouldHaveSameElementsFunc(tc.A, test.DiffEqual, tc.B), tc.ShouldFunc)
			a.So(ShouldNotHaveSameElementsFunc(tc.A, test.DiffEqual, tc.B), tc.ShouldNotFunc)
			a.So(ShouldHaveSameElementsDeep(tc.A, tc.B), tc.ShouldFunc)
			a.So(ShouldNotHaveSameElementsDeep(tc.A, tc.B), tc.ShouldNotFunc)
			a.So(ShouldHaveSameElementsDiff(tc.A, tc.B), tc.ShouldFunc)
			a.So(ShouldNotHaveSameElementsDiff(tc.A, tc.B), tc.ShouldNotFunc)
		})
	}
}

package nats

import (
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestCombineSubjects(t *testing.T) {
	a := assertions.New(t)

	for _, tc := range []struct {
		name     string
		subject1 string
		subject2 string
		expected string
	}{
		{
			name:     "EmptySubject1",
			subject1: "",
			subject2: "bar.bar2",
			expected: "bar.bar2",
		},
		{
			name:     "EmptySubject2",
			subject1: "foo.foo2",
			subject2: "",
			expected: "foo.foo2",
		},
		{
			name:     "BothProvided",
			subject1: "foo.foo2",
			subject2: "bar.bar2",
			expected: "foo.foo2.bar.bar2",
		},
		{
			name:     "NoneProvided",
			subject1: "",
			subject2: "",
			expected: "",
		},
		{
			name:     "Trailing1",
			subject1: "foo.",
			subject2: "",
			expected: "foo",
		},
		{
			name:     "Trailing2",
			subject1: "foo.",
			subject2: ".bar",
			expected: "foo.bar",
		},
		{
			name:     "Trailing3",
			subject1: ".foo.test.",
			subject2: ".bar.",
			expected: "foo.test.bar",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			a.So(combineSubjects(tc.subject1, tc.subject2), should.Equal, tc.expected)
		})
	}
}

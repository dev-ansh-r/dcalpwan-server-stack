package events

import (
	"testing"

	"github.com/smarty/assertions"
)

func TestUniqueStrings(t *testing.T) {
	t.Parallel()
	a := assertions.New(t)

	a.So(uniqueStrings([]string{
		"aaa", "bbb", "ccc",
	}), assertions.ShouldResemble, []string{
		"aaa", "bbb", "ccc",
	})

	a.So(uniqueStrings([]string{
		"aaa", "aaa", "bbb", "ccc",
	}), assertions.ShouldResemble, []string{
		"aaa", "bbb", "ccc",
	})

	a.So(uniqueStrings([]string{
		"aaa", "aaa", "bbb", "bbb", "bbb", "ccc",
	}), assertions.ShouldResemble, []string{
		"aaa", "bbb", "ccc",
	})

	a.So(uniqueStrings([]string{
		"aaa", "aaa", "bbb", "bbb", "bbb", "ccc", "ccc", "ccc", "ccc",
	}), assertions.ShouldResemble, []string{
		"aaa", "bbb", "ccc",
	})
}

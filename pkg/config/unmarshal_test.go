package config_test

import (
	"strings"
	"testing"

	"github.com/smarty/assertions"
	. "go.thethings.network/lorawan-stack/v3/pkg/config"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestUnmarshal(t *testing.T) {
	t.Parallel()
	a := assertions.New(t)

	mgr := Initialize("test", "test", defaults)
	a.So(mgr, should.NotBeNil)

	err := mgr.Parse()
	a.So(err, should.BeNil)

	err = mgr.MergeConfig(strings.NewReader(`file-only: 10`))
	a.So(err, should.BeNil)

	var res map[string]any
	err = mgr.Unmarshal(&res)
	a.So(err, should.BeNil)
	a.So(res, should.ContainKey, "file-only")
	a.So(res["file-only"], should.Resemble, 10)
}

func TestUnmarshalKey(t *testing.T) {
	t.Parallel()
	a := assertions.New(t)

	mgr := Initialize("test", "test", defaults)
	a.So(mgr, should.NotBeNil)

	err := mgr.Parse()
	a.So(err, should.BeNil)

	err = mgr.MergeConfig(strings.NewReader(`file-only: 10`))
	a.So(err, should.BeNil)

	var res int
	err = mgr.UnmarshalKey("file-only", &res)
	a.So(err, should.BeNil)
	a.So(res, should.Resemble, 10)
}

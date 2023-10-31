package config_test

import (
	"path"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/smarty/assertions"
	. "go.thethings.network/lorawan-stack/v3/pkg/config"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

type singleFileConfig struct {
	ConfigPath string            `name:"config"`
	Foo        string            `name:"foo"`
	Bar        map[string]string `name:"bar"`
}

func TestReadSingleConfig(t *testing.T) {
	t.Parallel()
	a := assertions.New(t)

	_, filename, _, _ := runtime.Caller(0)

	first := path.Join(filepath.Dir(filename), "first.yml")
	second := path.Join(filepath.Dir(filename), "second.yml")
	_ = second

	defaults := &singleFileConfig{}

	mgr := InitializeWithDefaults("empty", "empty", defaults)
	a.So(mgr, should.NotBeNil)

	err := mgr.Parse("--config", first)
	a.So(err, should.BeNil)

	err = mgr.ReadInConfig()
	a.So(err, should.BeNil)

	res := new(singleFileConfig)
	err = mgr.Unmarshal(res)
	a.So(err, should.BeNil)

	a.So(res.Foo, should.Resemble, "10")
	a.So(res.Bar, should.Resemble, map[string]string{
		"a": "baz",
		"b": "quu",
	})
}

type multiFileConfig struct {
	ConfigPaths []string          `name:"config"`
	Foo         string            `name:"foo"`
	Bar         map[string]string `name:"bar"`
}

func TestReadMultiConfig(t *testing.T) {
	t.Parallel()
	a := assertions.New(t)

	_, filename, _, _ := runtime.Caller(0)

	first := path.Join(filepath.Dir(filename), "first.yml")
	second := path.Join(filepath.Dir(filename), "second.yml")
	_ = second

	defaults := &multiFileConfig{}

	mgr := InitializeWithDefaults("empty", "empty", defaults)
	a.So(mgr, should.NotBeNil)

	err := mgr.Parse("--config", first, "--config", second)
	a.So(err, should.BeNil)

	err = mgr.ReadInConfig()
	a.So(err, should.BeNil)

	res := new(multiFileConfig)
	err = mgr.Unmarshal(res)
	a.So(err, should.BeNil)

	a.So(res.Foo, should.Resemble, "10")
	a.So(res.Bar, should.Resemble, map[string]string{
		"a": "baz",
		"b": "quu",
		"c": "yo!",
	})
}

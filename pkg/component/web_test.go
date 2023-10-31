package component

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/config"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

const (
	httpAddress     = "0.0.0.0:8097"
	metricsPassword = "secret-metrics-test-password"
	pprofPassword   = "secret-pprof-test-password"
	healthPassword  = "secret-health-test-password"
)

func TestPProf(t *testing.T) {
	a := assertions.New(t)

	config := &Config{
		ServiceBase: config.ServiceBase{
			HTTP: config.HTTP{
				Listen: httpAddress,
				Metrics: config.Metrics{
					Enable:   true,
					Password: metricsPassword,
				},
				PProf: config.PProf{
					Enable:   true,
					Password: pprofPassword,
				},
				Health: config.Health{
					Enable:   true,
					Password: healthPassword,
				},
			},
		},
	}
	c, err := New(test.GetLogger(t), config)
	a.So(err, should.BeNil)

	err = c.listenWeb()
	a.So(err, should.BeNil)

	client := &http.Client{}

	for _, tc := range []struct {
		path               string
		username, password string
	}{
		{
			path:     "debug/pprof",
			username: pprofUsername,
			password: pprofPassword,
		},
		{
			path:     "metrics",
			username: metricsUsername,
			password: metricsPassword,
		},
		{
			path:     "healthz",
			username: healthUsername,
			password: healthPassword,
		},
	} {
		t.Run(fmt.Sprintf("%s endpoint", tc.path), func(t *testing.T) {
			a := assertions.New(t)

			url := fmt.Sprintf("http://%s/%s", httpAddress, tc.path)
			res, err := client.Get(url)
			if !a.So(err, should.BeNil) {
				t.FailNow()
			}
			a.So(res.StatusCode, should.BeIn, []int{401, 403})

			req, err := http.NewRequest("GET", url, nil)
			if !a.So(err, should.BeNil) {
				t.FailNow()
			}
			req.SetBasicAuth(tc.username, tc.password)
			res, err = client.Do(req)
			if !a.So(err, should.BeNil) {
				t.FailNow()
			}
			a.So(res.StatusCode, should.BeBetweenOrEqual, 200, 299)
		})
	}
}

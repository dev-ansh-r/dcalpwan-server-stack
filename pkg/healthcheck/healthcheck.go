// Package healthcheck provides configuration of startup probes.
package healthcheck

import (
	"context"
	"net/http"
)

// Check is a function that determines the health of the HealtChecker.
type Check func(ctx context.Context) error

// HealthChecker manages checks for determining the healhiness of a component.
type HealthChecker interface {
	AddHTTPCheck(name string, url string) error
	AddPgCheck(name string, dsn string) error
	AddCheck(name string, check Check) error
	GetHandler() http.Handler
}

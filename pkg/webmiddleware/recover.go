package webmiddleware

import (
	"fmt"
	"net/http"
	"os"
	"runtime/debug"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/webhandlers"
)

// errHTTPRecovered is returned when a panic is caught from an HTTP handler.
var errHTTPRecovered = errors.DefineInternal("http_recovered", "Internal Server Error")

// Recover returns middleware that recovers from panics.
func Recover() MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if p := recover(); p != nil {
					fmt.Fprintln(os.Stderr, p)
					os.Stderr.Write(debug.Stack())
					var err error
					if pErr, ok := p.(error); ok {
						err = errHTTPRecovered.WithCause(pErr)
					} else {
						err = errHTTPRecovered.WithAttributes("panic", p)
					}
					log.FromContext(r.Context()).WithError(err).Error("Handler panicked")
					webhandlers.Error(w, r, err)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

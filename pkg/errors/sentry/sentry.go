// Package sentry converts errors to Sentry events.
package sentry

import (
	"strings"

	"github.com/getsentry/sentry-go"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
)

// NewEvent creates a new Sentry event for the given error.
func NewEvent(err error) *sentry.Event {
	evt := sentry.NewEvent()
	if err == nil {
		return evt
	}

	evt.Level = sentry.LevelError

	errStack := errors.Stack(err)

	messages := make([]string, 0, len(errStack))

	for i, err := range errStack {
		messages = append(messages, err.Error())
		exception := sentry.Exception{Value: err.Error()}
		if ttnErr, ok := errors.From(err); ok && ttnErr != nil {
			exception.Module = ttnErr.Namespace()
			exception.Type = ttnErr.Name()
			if i == 0 { // We set the namespace, name and ID from the first error in the chain.
				evt.Tags["error.namespace"] = ttnErr.Namespace()
				evt.Tags["error.name"] = ttnErr.Name()
				if correlationID := ttnErr.CorrelationID(); correlationID != "" {
					evt.EventID = sentry.EventID(correlationID)
				}
			}
			evt.Contexts[ttnErr.FullName()+" attributes"] = ttnErr.Attributes()
			if stackTrace := sentry.ExtractStacktrace(err); stackTrace != nil {
				exception.Stacktrace = stackTrace
			}
		}
		evt.Exception = append(evt.Exception, exception)
	}

	evt.Message = strings.Join(messages, "\n--- ")

	return evt
}

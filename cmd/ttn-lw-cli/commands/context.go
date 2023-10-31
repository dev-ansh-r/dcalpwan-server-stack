package commands

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

var cancelSignals = []os.Signal{syscall.SIGHUP, os.Interrupt, syscall.SIGTERM}

func newContext(parent context.Context) context.Context {
	ctx, _ := signal.NotifyContext(parent, cancelSignals...)
	return ctx
}

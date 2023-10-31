package io

import (
	"context"
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/task"
)

var (
	dialTaskBackoffIntervals = append(task.DialBackoffIntervals[:],
		time.Minute,
	)
	dialTaskExtendedBackoffIntervals = append(dialTaskBackoffIntervals[:],
		5*time.Minute,
		15*time.Minute,
		30*time.Minute,
	)
	// DialTaskBackoffConfig derives the component.DialTaskBackoffConfig and dynamically determines the backoff duration
	// based on recent error codes.
	DialTaskBackoffConfig = &task.BackoffConfig{
		Jitter: task.DefaultBackoffJitter,
		IntervalFunc: func(ctx context.Context, executionDuration time.Duration, invocation uint, err error) time.Duration {
			intervals := dialTaskBackoffIntervals
			switch {
			case errors.IsFailedPrecondition(err),
				errors.IsUnauthenticated(err),
				errors.IsPermissionDenied(err),
				errors.IsInvalidArgument(err),
				errors.IsAlreadyExists(err),
				errors.IsCanceled(err),
				errors.IsNotFound(err):
				intervals = dialTaskExtendedBackoffIntervals
			}
			switch {
			case executionDuration > task.DefaultBackoffResetDuration:
				return intervals[0]
			case invocation >= uint(len(intervals)):
				return intervals[len(intervals)-1]
			default:
				return intervals[invocation-1]
			}
		},
	}
)

package ratelimit

import (
	"fmt"
	"net/http"
	"time"

	"google.golang.org/grpc/metadata"
)

// Result contains rate limiting metadata.
type Result struct {
	Limit      int
	Remaining  int
	ResetAfter time.Duration
	RetryAfter time.Duration
}

// IsZero checks if the result is empty.
func (r Result) IsZero() bool {
	return r.Limit == 0 && r.Remaining == 0 && r.ResetAfter == 0 && r.RetryAfter == 0
}

// GRPCHeaders returns gRPC headers from a rate limiting result.
func (r Result) GRPCHeaders() metadata.MD {
	if r.IsZero() {
		return metadata.MD{}
	}
	return metadata.Pairs(
		"x-rate-limit-limit", fmt.Sprintf("%d", r.Limit),
		"x-rate-limit-available", fmt.Sprintf("%d", r.Remaining),
		"x-rate-limit-reset", fmt.Sprintf("%d", r.ResetAfter/time.Second),
		"x-rate-limit-retry", fmt.Sprintf("%d", r.RetryAfter/time.Second),
	)
}

// SetHTTPHeaders sets HTTP headers from a rate limiting result.
func (r Result) SetHTTPHeaders(header http.Header) {
	if r.IsZero() {
		return
	}
	header.Add("x-rate-limit-limit", fmt.Sprintf("%d", r.Limit))
	header.Add("x-rate-limit-available", fmt.Sprintf("%d", r.Remaining))
	header.Add("x-rate-limit-reset", fmt.Sprintf("%d", r.ResetAfter/time.Second))
	header.Add("x-rate-limit-retry", fmt.Sprintf("%d", r.RetryAfter/time.Second))
}

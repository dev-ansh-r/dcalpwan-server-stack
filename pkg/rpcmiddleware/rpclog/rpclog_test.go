package rpclog

import (
	"testing"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
)

func TestShouldSuppressLogEvaluatesCorrectly(t *testing.T) {
	ignoredError := errors.Define("ignored", "ignored")
	nonIgnoredError := errors.Define("non_ignored", "non_ignored")
	ignoredErrorKey := ignoredError.FullName()

	tests := []struct {
		inputCfg methodLogConfig
		inputErr error
		expected bool
	}{
		// test behavior when no error occurs
		{
			inputCfg: methodLogConfig{
				IgnoreSuccess: true,
				IgnoredErrors: map[string]struct{}{},
			},
			inputErr: nil,
			expected: true,
		},
		{
			inputCfg: methodLogConfig{
				IgnoreSuccess: true,
				IgnoredErrors: map[string]struct{}{
					ignoredErrorKey: {},
				},
			},
			inputErr: nil,
			expected: true,
		},
		{
			inputCfg: methodLogConfig{
				IgnoreSuccess: false,
				IgnoredErrors: map[string]struct{}{},
			},
			inputErr: nil,
			expected: false,
		},
		{
			inputCfg: methodLogConfig{
				IgnoreSuccess: false,
				IgnoredErrors: map[string]struct{}{
					ignoredErrorKey: {},
				},
			},
			inputErr: nil,
			expected: false,
		},
		// test behavior when an error occurs that is ignored
		{
			inputCfg: methodLogConfig{
				IgnoreSuccess: true,
				IgnoredErrors: map[string]struct{}{},
			},
			inputErr: ignoredError,
			expected: false,
		},
		{
			inputCfg: methodLogConfig{
				IgnoreSuccess: true,
				IgnoredErrors: map[string]struct{}{
					ignoredErrorKey: {},
				},
			},
			inputErr: ignoredError,
			expected: true,
		},
		{
			inputCfg: methodLogConfig{
				IgnoreSuccess: false,
				IgnoredErrors: map[string]struct{}{},
			},
			inputErr: ignoredError,
			expected: false,
		},
		{
			inputCfg: methodLogConfig{
				IgnoreSuccess: false,
				IgnoredErrors: map[string]struct{}{
					ignoredErrorKey: {},
				},
			},
			inputErr: ignoredError,
			expected: true,
		},
		// test behavior when an error occurs that is not ignored
		{
			inputCfg: methodLogConfig{
				IgnoreSuccess: true,
				IgnoredErrors: map[string]struct{}{},
			},
			inputErr: nonIgnoredError,
			expected: false,
		},
		{
			inputCfg: methodLogConfig{
				IgnoreSuccess: true,
				IgnoredErrors: map[string]struct{}{
					ignoredErrorKey: {},
				},
			},
			inputErr: nonIgnoredError,
			expected: false,
		},
		{
			inputCfg: methodLogConfig{
				IgnoreSuccess: false,
				IgnoredErrors: map[string]struct{}{},
			},
			inputErr: nonIgnoredError,
			expected: false,
		},
		{
			inputCfg: methodLogConfig{
				IgnoreSuccess: false,
				IgnoredErrors: map[string]struct{}{
					ignoredErrorKey: {},
				},
			},
			inputErr: nonIgnoredError,
			expected: false,
		},
	}

	for _, test := range tests {
		actual := shouldSuppressLog(test.inputCfg, test.inputErr)
		if actual != test.expected {
			t.Errorf("Expected %t, got %t", test.expected, actual)
		}
	}
}

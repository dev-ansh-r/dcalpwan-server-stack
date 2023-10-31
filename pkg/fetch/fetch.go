// Package fetch offers abstractions to fetch a file with the same method,
// regardless of a location (filesystem, HTTP...).
package fetch

import (
	"fmt"
	"net/url"
	"path"
	"path/filepath"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// realPath replaces path root of p by r if r != "" and returns p otherwise.
// realPath assumes paths separated by forward slashes and assumes that both paths are cleaned.
func realPath(r, p string) (string, error) {
	if r == "" {
		return p, nil
	}
	if !path.IsAbs(p) {
		return path.Join(r, p), nil
	}
	return path.Join(r, p[1:]), nil
}

// realURLPath replaces path root of p by r if r != "" and returns p otherwise.
// realURLPath assumes URL paths and assumes that both paths are cleaned.
func realURLPath(rURL *url.URL, p string) (string, error) {
	pURL, err := url.Parse(p)
	if err != nil {
		return "", err
	}
	if rURL == nil {
		if !pURL.IsAbs() {
			return "", errSchemeNotSpecified.New()
		}
		return p, nil
	}
	if pURL.IsAbs() {
		return "", errSchemeSpecified.New()
	}
	if !rURL.IsAbs() {
		return "", errSchemeNotSpecified.New()
	}
	return fmt.Sprintf("%s/%s", rURL, pURL.EscapedPath()), nil
}

// realOSPath replaces path root of p by r if r != "" and returns p otherwise.
// realOSPath assumes operating system paths and that both paths are cleaned.
func realOSPath(r, p string) (string, error) {
	if r == "" {
		return p, nil
	}
	if !filepath.IsAbs(p) {
		return filepath.Join(r, p), nil
	}
	if filepath.VolumeName(p) != "" {
		return "", errVolumeSpecified.New()
	}
	return filepath.Join(r, p[1:]), nil
}

// Interface is an abstraction for file retrieval.
type Interface interface {
	File(pathElements ...string) ([]byte, error)
}

type baseFetcher struct {
	latency prometheus.Observer
}

func (f baseFetcher) observeLatency(d time.Duration) {
	if f.latency != nil {
		f.latency.Observe(d.Seconds())
	}
}

package fetch

import (
	"path"
	"time"
)

// MemFetcher represents the memory fetcher.
type memFetcher struct {
	baseFetcher
	store map[string][]byte
}

// File gets content from memory.
func (f *memFetcher) File(pathElements ...string) ([]byte, error) {
	if len(pathElements) == 0 {
		return nil, errFilenameNotSpecified.New()
	}

	start := time.Now()

	p := path.Join(pathElements...)
	content, ok := f.store[p]
	if !ok {
		return nil, errFileNotFound.WithAttributes("filename", p)
	}

	f.observeLatency(time.Since(start))
	return content, nil
}

// NewMemFetcher initializes a new memory fetcher.
func NewMemFetcher(store map[string][]byte) Interface {
	return &memFetcher{
		store: store,
	}
}

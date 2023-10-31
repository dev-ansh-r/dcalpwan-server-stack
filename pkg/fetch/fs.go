package fetch

import (
	"os"
	"path/filepath"
	"time"
)

type fsFetcher struct {
	baseFetcher
	root string
}

func (f fsFetcher) File(pathElements ...string) ([]byte, error) {
	if len(pathElements) == 0 {
		return nil, errFilenameNotSpecified.New()
	}

	start := time.Now()

	p := filepath.Join(pathElements...)
	rp, err := realOSPath(f.root, p)
	if err != nil {
		return nil, err
	}
	content, err := os.ReadFile(rp)
	if err == nil {
		f.observeLatency(time.Since(start))
		return content, nil
	}

	if os.IsNotExist(err) {
		return nil, errFileNotFound.WithAttributes("filename", p)
	}
	return nil, errCouldNotReadFile.WithCause(err).WithAttributes("filename", p)
}

// FromFilesystem returns an interface that fetches files from the local filesystem.
func FromFilesystem(rootElements ...string) Interface {
	root := filepath.Join(rootElements...)
	return fsFetcher{
		baseFetcher: baseFetcher{
			latency: fetchLatency.WithLabelValues("fs", root),
		},
		root: root,
	}
}

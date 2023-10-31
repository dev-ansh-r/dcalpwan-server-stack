package fetch

import (
	"path/filepath"
)

type basePathFetcher struct {
	Interface
	basePath []string
}

func (f basePathFetcher) File(pathElements ...string) ([]byte, error) {
	if len(pathElements) == 0 {
		return nil, errFilenameNotSpecified.New()
	}

	// NOTE: filepath.IsAbs returns true for paths starting with '/' on all supported operating systems.
	if filepath.IsAbs(pathElements[0]) {
		return f.Interface.File(pathElements...)
	}
	return f.Interface.File(append(f.basePath, pathElements...)...)
}

// WithBasePath returns an Interface, which preprends basePath to non-absolute requested paths.
func WithBasePath(f Interface, basePath ...string) Interface {
	return basePathFetcher{
		Interface: f,
		basePath:  append(basePath[:0:0], basePath...),
	}
}

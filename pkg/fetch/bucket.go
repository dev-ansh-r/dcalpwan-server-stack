package fetch

import (
	"context"
	"path"
	"time"

	"gocloud.dev/blob"
	"gocloud.dev/gcerrors"
)

type bucketFetcher struct {
	baseFetcher
	context context.Context
	bucket  *blob.Bucket
	root    string
}

func (f *bucketFetcher) File(pathElements ...string) ([]byte, error) {
	if len(pathElements) == 0 {
		return nil, errFilenameNotSpecified.New()
	}

	start := time.Now()

	p := path.Join(pathElements...)
	rp, err := realPath(f.root, p)
	if err != nil {
		return nil, err
	}
	content, err := f.bucket.ReadAll(f.context, rp)
	if err == nil {
		f.observeLatency(time.Since(start))
		return content, nil
	}

	if gcerrors.Code(err) == gcerrors.NotFound {
		return nil, errFileNotFound.WithAttributes("filename", p)
	}
	return nil, errCouldNotReadFile.WithCause(err).WithAttributes("filename", p)
}

// FromBucket returns an interface that fetches files from the given blob bucket.
func FromBucket(ctx context.Context, bucket *blob.Bucket, root string) Interface {
	root = path.Clean(root)
	return &bucketFetcher{
		baseFetcher: baseFetcher{
			latency: fetchLatency.WithLabelValues("bucket", root),
		},
		context: ctx,
		bucket:  bucket,
		root:    root,
	}
}

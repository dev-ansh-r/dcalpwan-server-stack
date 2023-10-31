package fetch_test

import (
	"fmt"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/fetch"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func ExampleNewMemFetcher() {
	fetcher := fetch.NewMemFetcher(map[string][]byte{
		"file.txt":     []byte("content"),
		"dir/file.txt": []byte("content"),
	})
	content, err := fetcher.File("dir/file.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Content of myFile.yml")
	fmt.Println(string(content))
}

func TestMemFetcher(t *testing.T) {
	a := assertions.New(t)
	fetcher := fetch.NewMemFetcher(map[string][]byte{
		"file.txt":     []byte("content1"),
		"dir/file.txt": []byte("content2"),
	})

	// Read a file and test content retrieval.
	{
		content, err := fetcher.File("file.txt")
		a.So(err, should.BeNil)
		a.So(string(content), should.Equal, "content1")
	}

	// Read from a subdirectory and test content retrieval.
	{
		content, err := fetcher.File("dir", "file.txt")
		a.So(err, should.BeNil)
		a.So(string(content), should.Equal, "content2")
	}

	// Read from a non existing path.
	{
		_, err := fetcher.File("notfound.txt")
		a.So(err, should.NotBeNil)
	}
}

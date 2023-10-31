package fetch_test

import (
	"fmt"
	"net/http"

	"go.thethings.network/lorawan-stack/v3/pkg/fetch"
)

func Example() {
	fetcher, err := fetch.FromHTTP(http.DefaultClient, "http://webserver.thethings.network/repository")
	if err != nil {
		panic(err)
	}
	content, err := fetcher.File("README.md")
	if err != nil {
		panic(err)
	}

	fmt.Println("Content of http://webserver.thethings.network/repository/README.md:")
	fmt.Println(string(content))
}

package api

import (
	"encoding/json"
	"io"
	"net/http"
	"unicode/utf8"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/goproto"
)

const maxResponseSize = (1 << 24) // 16 MiB

var errRequest = errors.DefineUnavailable("request", "LoRa Cloud GLS request")

func parse(result any, res *http.Response) error {
	defer res.Body.Close()
	defer io.Copy(io.Discard, res.Body) // nolint:errcheck
	reader := io.LimitReader(res.Body, maxResponseSize)
	if res.StatusCode < 200 || res.StatusCode > 299 {
		body, _ := io.ReadAll(reader)
		m := map[string]any{"status_code": res.StatusCode}
		if utf8.Valid(body) {
			m["body"] = string(body)
		}
		detail, err := goproto.Struct(m)
		if err != nil {
			return err
		}
		return errRequest.WithDetails(detail)
	}
	return json.NewDecoder(reader).Decode(result)
}

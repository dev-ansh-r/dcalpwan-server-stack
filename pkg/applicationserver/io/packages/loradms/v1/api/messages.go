package api

import (
	"encoding/json"
	"io"
	"net/http"
	"unicode/utf8"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/goproto"
	"google.golang.org/protobuf/proto"
)

const (
	sendOperation = "send"

	maxResponseSize = (1 << 24) // 16 MiB
)

type baseResponse struct {
	Result any      `json:"result"`
	Errors []string `json:"errors"`
}

var errRequest = errors.DefineUnavailable("request", "LoRaCloud DMS request")

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
	r := &baseResponse{
		Result: result,
	}
	if err := json.NewDecoder(reader).Decode(r); err != nil {
		return err
	}
	if len(r.Errors) == 0 {
		return nil
	}
	var details []proto.Message
	for _, message := range r.Errors {
		detail, err := goproto.Value(message)
		if err != nil {
			return err
		}
		details = append(details, detail)
	}
	return errRequest.WithDetails(details...)
}

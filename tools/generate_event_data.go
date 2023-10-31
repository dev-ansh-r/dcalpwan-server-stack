//go:build ignore
// +build ignore

package main

import (
	"log"
	"os"

	jsoniter "github.com/json-iterator/go"
	_ "go.thethings.network/lorawan-stack/v3/cmd/ttn-lw-cli/commands"
	_ "go.thethings.network/lorawan-stack/v3/cmd/ttn-lw-stack/commands"
	"go.thethings.network/lorawan-stack/v3/pkg/events"
	"go.thethings.network/lorawan-stack/v3/pkg/jsonpb"
)

var streamPool jsoniter.StreamPool = jsoniter.Config{
	IndentionStep: 2,
}.Froze()

var json = jsonpb.TTN()

func main() {
	messagesFile := "events.json"
	if len(os.Args) == 2 {
		messagesFile = os.Args[1]
	}
	f, err := os.OpenFile(messagesFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	enc := streamPool.BorrowStream(f)
	defer func() {
		enc.Flush()
		streamPool.ReturnStream(enc)
	}()
	enc.WriteObjectStart()
	for i, e := range events.All().Definitions() {
		if i > 0 {
			enc.WriteMore()
		}
		enc.WriteObjectField(e.Name())
		enc.WriteObjectStart()

		enc.WriteObjectField("name")
		enc.WriteString(e.Name())

		enc.WriteMore()

		enc.WriteObjectField("description")
		enc.WriteString(e.Description())

		if e.DataType() != nil {
			enc.WriteMore()

			enc.WriteObjectField("data")

			raw, err := json.Marshal(e.DataType())
			if err != nil {
				log.Fatal(err)
			}
			enc.WriteVal(jsoniter.RawMessage(raw))
		}

		enc.WriteObjectEnd()
	}
	enc.WriteObjectEnd()
}

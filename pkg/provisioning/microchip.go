package provisioning

import (
	"strings"

	"google.golang.org/protobuf/types/known/structpb"
)

// Microchip is the provisioning ID for Microchip devices.
const Microchip = "microchip"

type microchip struct{}

// UniqueID returns the serial number.
func (p *microchip) UniqueID(entry *structpb.Struct) (string, error) {
	if entry == nil {
		return "", errEntry.New()
	}
	sn := entry.Fields["uniqueId"].GetStringValue()
	if sn == "" {
		return "", errEntry.New()
	}
	return strings.ToUpper(sn), nil
}

func init() {
	Register(Microchip, new(microchip))
}

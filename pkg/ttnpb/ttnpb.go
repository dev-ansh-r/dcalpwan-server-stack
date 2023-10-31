package ttnpb

import (
	"math"

	"go.thethings.network/lorawan-stack/v3/pkg/util/byteutil"
)

func marshalBinaryEnum(v int32) []byte {
	switch v := uint32(v); {
	case v <= 255:
		return []byte{byte(v)}
	case v <= math.MaxUint16:
		return byteutil.AppendUint16(make([]byte, 2), uint16(v), 2)
	case v <= byteutil.MaxUint24:
		return byteutil.AppendUint32(make([]byte, 3), v, 3)
	default:
		return byteutil.AppendUint32(make([]byte, 4), v, 4)
	}
}

func unmarshalEnumFromBinary(typName string, names map[int32]string, b []byte) (int32, error) {
	if len(b) > 4 {
		return 0, errCouldNotParse(typName)(string(b))
	}
	i := int32(byteutil.ParseUint32(b))
	if _, ok := names[i]; !ok {
		return 0, errCouldNotParse(typName)(string(b))
	}
	return i, nil
}

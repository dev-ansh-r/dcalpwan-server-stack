package loraclouddevicemanagementv1

import (
	"go.thethings.network/lorawan-stack/v3/pkg/applicationserver/io/packages/loradms/v1/api/objects"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
)

var errTLVRecordTooSmall = errors.DefineInvalidArgument("tlv_record_too_small", "TLV record payload is too small")

func parseTLVPayload(record objects.Hex, f func(uint8, []byte) error) error {
	for len(record) >= 2 {
		tag := record[0]
		length := int(record[1])
		if length+2 > len(record) {
			return errTLVRecordTooSmall.New()
		}

		bytes := []byte(record[2 : 2+length])
		record = record[length+2:]

		if err := f(tag, bytes); err != nil {
			return err
		}
	}
	return nil
}

package ttnpb

import (
	"strings"
	"testing"

	"go.thethings.network/lorawan-stack/v3/pkg/i18n"
)

func TestEnumMessages(t *testing.T) {
	testCases := []struct {
		desc  string
		names map[int32]string
	}{
		{desc: "GrantType", names: GrantType_name},
		{desc: "State", names: State_name},
		{desc: "ContactType", names: ContactType_name},
		{desc: "ContactMethod", names: ContactMethod_name},
		{desc: "MType", names: MType_name},
		{desc: "JoinRequestType", names: JoinRequestType_name},
		{desc: "RejoinRequestType", names: RejoinRequestType_name},
		{desc: "CFListType", names: CFListType_name},
		{desc: "MACCommandIdentifier", names: MACCommandIdentifier_name},
		{desc: "LocationSource", names: LocationSource_name},
		{desc: "PayloadFormatter", names: PayloadFormatter_name},
		{desc: "Right", names: Right_name},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			for _, s := range tc.names {
				if strings.ToLower(s) == s {
					continue // ignore "private" values.
				}
				if i18n.Get("enum:"+s) == nil {
					t.Errorf("message descriptor for %q is nil", s)
				}
			}
		})
	}
}

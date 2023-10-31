package rpclog

import "testing"

func TestParseMethodLogCfgParsesInputsCorrectly(t *testing.T) {
	tests := []struct {
		input              string
		expectedMethodName string
		expectedCfg        methodLogConfig
	}{
		{
			input:              "/ttn.lorawan.v3.GsNs/HandleUplink",
			expectedMethodName: "/ttn.lorawan.v3.GsNs/HandleUplink",
			expectedCfg: methodLogConfig{
				IgnoreSuccess: true,
				IgnoredErrors: map[string]struct{}{},
			},
		},
		{
			input:              "/ttn.lorawan.v3.GsNs/HandleUplink:pkg/networkserver:duplicate_uplink",
			expectedMethodName: "/ttn.lorawan.v3.GsNs/HandleUplink",
			expectedCfg: methodLogConfig{
				IgnoreSuccess: false,
				IgnoredErrors: map[string]struct{}{
					"pkg/networkserver:duplicate_uplink": {},
				},
			},
		},
		{
			input:              "/ttn.lorawan.v3.GsNs/HandleUplink:pkg/networkserver:duplicate_uplink;pkg/networkserver:device_not_found",
			expectedMethodName: "/ttn.lorawan.v3.GsNs/HandleUplink",
			expectedCfg: methodLogConfig{
				IgnoreSuccess: false,
				IgnoredErrors: map[string]struct{}{
					"pkg/networkserver:duplicate_uplink": {},
					"pkg/networkserver:device_not_found": {},
				},
			},
		},
		{
			input:              "/ttn.lorawan.v3.GsNs/HandleUplink:;pkg/networkserver:duplicate_uplink;pkg/networkserver:device_not_found",
			expectedMethodName: "/ttn.lorawan.v3.GsNs/HandleUplink",
			expectedCfg: methodLogConfig{
				IgnoreSuccess: true,
				IgnoredErrors: map[string]struct{}{
					"pkg/networkserver:duplicate_uplink": {},
					"pkg/networkserver:device_not_found": {},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			actualMethodName, actualCfg := parseMethodLogCfg(test.input)
			expectedMethodName := test.expectedMethodName
			expectedCfg := test.expectedCfg
			if actualMethodName != expectedMethodName {
				t.Fatalf("Expected method name to be %s, got %s", expectedMethodName, actualMethodName)
			}
			if actualCfg.IgnoreSuccess != expectedCfg.IgnoreSuccess {
				t.Fatalf("Expected method log config to be %v, got %v", expectedCfg, actualCfg)
			}
			if len(actualCfg.IgnoredErrors) != len(expectedCfg.IgnoredErrors) {
				t.Fatalf("Expected method log config to be %v, got %v", expectedCfg, actualCfg)
			}
			for err := range actualCfg.IgnoredErrors {
				if _, ok := expectedCfg.IgnoredErrors[err]; !ok {
					t.Fatalf("Expected method log config to be %v, got %v", expectedCfg, actualCfg)
				}
			}
		})
	}
}

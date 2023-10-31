// Code generated by protoc-gen-go-json. DO NOT EDIT.
// versions:
// - protoc-gen-go-json v1.5.1
// - protoc             v4.22.2
// source: ttn/lorawan/v3/networkserver.proto

package ttnpb

import (
	jsonplugin "github.com/TheThingsIndustries/protoc-gen-go-json/jsonplugin"
	types "go.thethings.network/lorawan-stack/v3/pkg/types"
)

// MarshalProtoJSON marshals the GenerateDevAddrResponse message to JSON.
func (x *GenerateDevAddrResponse) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.DevAddr) > 0 || s.HasField("dev_addr") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("dev_addr")
		types.MarshalHEXBytes(s.WithField("dev_addr"), x.DevAddr)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the GenerateDevAddrResponse to JSON.
func (x *GenerateDevAddrResponse) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the GenerateDevAddrResponse message from JSON.
func (x *GenerateDevAddrResponse) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "dev_addr", "devAddr":
			s.AddField("dev_addr")
			x.DevAddr = types.Unmarshal4Bytes(s.WithField("dev_addr", false))
		}
	})
}

// UnmarshalJSON unmarshals the GenerateDevAddrResponse from JSON.
func (x *GenerateDevAddrResponse) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the GetDefaultMACSettingsRequest message to JSON.
func (x *GetDefaultMACSettingsRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.FrequencyPlanId != "" || s.HasField("frequency_plan_id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("frequency_plan_id")
		s.WriteString(x.FrequencyPlanId)
	}
	if x.LorawanPhyVersion != 0 || s.HasField("lorawan_phy_version") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("lorawan_phy_version")
		x.LorawanPhyVersion.MarshalProtoJSON(s)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the GetDefaultMACSettingsRequest to JSON.
func (x *GetDefaultMACSettingsRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the GetDefaultMACSettingsRequest message from JSON.
func (x *GetDefaultMACSettingsRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "frequency_plan_id", "frequencyPlanId":
			s.AddField("frequency_plan_id")
			x.FrequencyPlanId = s.ReadString()
		case "lorawan_phy_version", "lorawanPhyVersion":
			s.AddField("lorawan_phy_version")
			x.LorawanPhyVersion.UnmarshalProtoJSON(s)
		}
	})
}

// UnmarshalJSON unmarshals the GetDefaultMACSettingsRequest from JSON.
func (x *GetDefaultMACSettingsRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the GetNetIDResponse message to JSON.
func (x *GetNetIDResponse) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.NetId) > 0 || s.HasField("net_id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("net_id")
		types.MarshalHEXBytes(s.WithField("net_id"), x.NetId)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the GetNetIDResponse to JSON.
func (x *GetNetIDResponse) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the GetNetIDResponse message from JSON.
func (x *GetNetIDResponse) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "net_id", "netId":
			s.AddField("net_id")
			x.NetId = types.Unmarshal3Bytes(s.WithField("net_id", false))
		}
	})
}

// UnmarshalJSON unmarshals the GetNetIDResponse from JSON.
func (x *GetNetIDResponse) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the GetDeviceAdressPrefixesResponse message to JSON.
func (x *GetDeviceAdressPrefixesResponse) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.DevAddrPrefixes) > 0 || s.HasField("dev_addr_prefixes") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("dev_addr_prefixes")
		types.MarshalDevAddrPrefixSlice(s.WithField("dev_addr_prefixes"), x.DevAddrPrefixes)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the GetDeviceAdressPrefixesResponse to JSON.
func (x *GetDeviceAdressPrefixesResponse) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the GetDeviceAdressPrefixesResponse message from JSON.
func (x *GetDeviceAdressPrefixesResponse) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "dev_addr_prefixes", "devAddrPrefixes":
			s.AddField("dev_addr_prefixes")
			if s.ReadNil() {
				x.DevAddrPrefixes = nil
				return
			}
			x.DevAddrPrefixes = types.UnmarshalDevAddrPrefixSlice(s.WithField("dev_addr_prefixes", false))
		}
	})
}

// UnmarshalJSON unmarshals the GetDeviceAdressPrefixesResponse from JSON.
func (x *GetDeviceAdressPrefixesResponse) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

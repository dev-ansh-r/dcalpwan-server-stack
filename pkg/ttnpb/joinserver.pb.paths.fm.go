// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

var SessionKeyRequestFieldPathsNested = []string{
	"dev_eui",
	"join_eui",
	"session_key_id",
}

var SessionKeyRequestFieldPathsTopLevel = []string{
	"dev_eui",
	"join_eui",
	"session_key_id",
}
var NwkSKeysResponseFieldPathsNested = []string{
	"f_nwk_s_int_key",
	"f_nwk_s_int_key.encrypted_key",
	"f_nwk_s_int_key.kek_label",
	"f_nwk_s_int_key.key",
	"nwk_s_enc_key",
	"nwk_s_enc_key.encrypted_key",
	"nwk_s_enc_key.kek_label",
	"nwk_s_enc_key.key",
	"s_nwk_s_int_key",
	"s_nwk_s_int_key.encrypted_key",
	"s_nwk_s_int_key.kek_label",
	"s_nwk_s_int_key.key",
}

var NwkSKeysResponseFieldPathsTopLevel = []string{
	"f_nwk_s_int_key",
	"nwk_s_enc_key",
	"s_nwk_s_int_key",
}
var AppSKeyResponseFieldPathsNested = []string{
	"app_s_key",
	"app_s_key.encrypted_key",
	"app_s_key.kek_label",
	"app_s_key.key",
}

var AppSKeyResponseFieldPathsTopLevel = []string{
	"app_s_key",
}
var CryptoServicePayloadRequestFieldPathsNested = []string{
	"ids",
	"ids.application_ids",
	"ids.application_ids.application_id",
	"ids.dev_addr",
	"ids.dev_eui",
	"ids.device_id",
	"ids.join_eui",
	"lorawan_version",
	"payload",
	"provisioner_id",
	"provisioning_data",
}

var CryptoServicePayloadRequestFieldPathsTopLevel = []string{
	"ids",
	"lorawan_version",
	"payload",
	"provisioner_id",
	"provisioning_data",
}
var CryptoServicePayloadResponseFieldPathsNested = []string{
	"payload",
}

var CryptoServicePayloadResponseFieldPathsTopLevel = []string{
	"payload",
}
var JoinAcceptMICRequestFieldPathsNested = []string{
	"dev_nonce",
	"join_request_type",
	"payload_request",
	"payload_request.ids",
	"payload_request.ids.application_ids",
	"payload_request.ids.application_ids.application_id",
	"payload_request.ids.dev_addr",
	"payload_request.ids.dev_eui",
	"payload_request.ids.device_id",
	"payload_request.ids.join_eui",
	"payload_request.lorawan_version",
	"payload_request.payload",
	"payload_request.provisioner_id",
	"payload_request.provisioning_data",
}

var JoinAcceptMICRequestFieldPathsTopLevel = []string{
	"dev_nonce",
	"join_request_type",
	"payload_request",
}
var DeriveSessionKeysRequestFieldPathsNested = []string{
	"dev_nonce",
	"ids",
	"ids.application_ids",
	"ids.application_ids.application_id",
	"ids.dev_addr",
	"ids.dev_eui",
	"ids.device_id",
	"ids.join_eui",
	"join_nonce",
	"lorawan_version",
	"net_id",
	"provisioner_id",
	"provisioning_data",
}

var DeriveSessionKeysRequestFieldPathsTopLevel = []string{
	"dev_nonce",
	"ids",
	"join_nonce",
	"lorawan_version",
	"net_id",
	"provisioner_id",
	"provisioning_data",
}
var GetRootKeysRequestFieldPathsNested = []string{
	"ids",
	"ids.application_ids",
	"ids.application_ids.application_id",
	"ids.dev_addr",
	"ids.dev_eui",
	"ids.device_id",
	"ids.join_eui",
	"provisioner_id",
	"provisioning_data",
}

var GetRootKeysRequestFieldPathsTopLevel = []string{
	"ids",
	"provisioner_id",
	"provisioning_data",
}
var ProvisionEndDevicesRequestFieldPathsNested = []string{
	"application_ids",
	"application_ids.application_id",
	"end_devices",
	"end_devices.from_data",
	"end_devices.from_data.join_eui",
	"end_devices.list",
	"end_devices.list.end_device_ids",
	"end_devices.list.join_eui",
	"end_devices.range",
	"end_devices.range.join_eui",
	"end_devices.range.start_dev_eui",
	"provisioner_id",
	"provisioning_data",
}

var ProvisionEndDevicesRequestFieldPathsTopLevel = []string{
	"application_ids",
	"end_devices",
	"provisioner_id",
	"provisioning_data",
}
var ApplicationActivationSettingsFieldPathsNested = []string{
	"application_server_id",
	"home_net_id",
	"kek",
	"kek.encrypted_key",
	"kek.kek_label",
	"kek.key",
	"kek_label",
}

var ApplicationActivationSettingsFieldPathsTopLevel = []string{
	"application_server_id",
	"home_net_id",
	"kek",
	"kek_label",
}
var GetApplicationActivationSettingsRequestFieldPathsNested = []string{
	"application_ids",
	"application_ids.application_id",
	"field_mask",
}

var GetApplicationActivationSettingsRequestFieldPathsTopLevel = []string{
	"application_ids",
	"field_mask",
}
var SetApplicationActivationSettingsRequestFieldPathsNested = []string{
	"application_ids",
	"application_ids.application_id",
	"field_mask",
	"settings",
	"settings.application_server_id",
	"settings.home_net_id",
	"settings.kek",
	"settings.kek.encrypted_key",
	"settings.kek.kek_label",
	"settings.kek.key",
	"settings.kek_label",
}

var SetApplicationActivationSettingsRequestFieldPathsTopLevel = []string{
	"application_ids",
	"field_mask",
	"settings",
}
var DeleteApplicationActivationSettingsRequestFieldPathsNested = []string{
	"application_ids",
	"application_ids.application_id",
}

var DeleteApplicationActivationSettingsRequestFieldPathsTopLevel = []string{
	"application_ids",
}
var JoinEUIPrefixFieldPathsNested = []string{
	"join_eui",
	"length",
}

var JoinEUIPrefixFieldPathsTopLevel = []string{
	"join_eui",
	"length",
}
var JoinEUIPrefixesFieldPathsNested = []string{
	"prefixes",
}

var JoinEUIPrefixesFieldPathsTopLevel = []string{
	"prefixes",
}
var GetDefaultJoinEUIResponseFieldPathsNested = []string{
	"join_eui",
}

var GetDefaultJoinEUIResponseFieldPathsTopLevel = []string{
	"join_eui",
}
var ProvisionEndDevicesRequest_IdentifiersListFieldPathsNested = []string{
	"end_device_ids",
	"join_eui",
}

var ProvisionEndDevicesRequest_IdentifiersListFieldPathsTopLevel = []string{
	"end_device_ids",
	"join_eui",
}
var ProvisionEndDevicesRequest_IdentifiersRangeFieldPathsNested = []string{
	"join_eui",
	"start_dev_eui",
}

var ProvisionEndDevicesRequest_IdentifiersRangeFieldPathsTopLevel = []string{
	"join_eui",
	"start_dev_eui",
}
var ProvisionEndDevicesRequest_IdentifiersFromDataFieldPathsNested = []string{
	"join_eui",
}

var ProvisionEndDevicesRequest_IdentifiersFromDataFieldPathsTopLevel = []string{
	"join_eui",
}

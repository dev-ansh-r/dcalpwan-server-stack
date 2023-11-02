
import { defineMessages } from 'react-intl'

export default defineMessages({
  applicationDataAllowDesc: 'Allow downlink messages with FPort between 1 and 255',
  applicationDataDesc: 'Forward uplink messages with FPort 1-255',
  joinAcceptDesc: 'Allow join accept messages',
  joinRequest: 'Join request',
  joinRequestDesc: 'Forward join-request messages',
  localizationInformation: 'Localization  information',
  localizationInformationDesc: 'Forward gateway location, RSSI, SNR and fine timestamp',
  macData: 'MAC data',
  macDataAllowDesc: 'Allow downlink messages with FPort of 0',
  macDataDesc: 'Forward uplink messages with FPort 0',
  signalQualityInformation: 'Signal quality information',
  signalQualityInformationDesc: 'Forward RSSI and SNR',
  forwardsJoinRequest: 'Join request messages are forwarded',
  doesNotForwardJoinRequest: 'Join request messages are not forwarded',
  forwardsMacData: 'MAC data is forwarded',
  doesNotForwardMacData: 'MAC data is not forwarded',
  forwardsApplicationData: 'Application data is forwarded',
  doesNotForwardApplicationData: 'Application data is not forwarded',
  forwardsSignalQuality: 'Signal quality information is forwarded',
  doesNotForwardSignalQuality: 'Signal quality information is not forwarded',
  forwardsLocalization: 'Localization information is forwarded',
  doesNotForwardLocalization: 'Localization information is not forwarded',
  allowsJoinAccept: 'Join accept messages are allowed',
  doesNotAllowJoinAccept: 'Join accept messages are not allowed',
  allowsMacData: 'MAC data is allowed',
  doesNotAllowMacData: 'MAC data is not allowed',
  allowsApplicationData: 'Application data is allowed',
  doesNotAllowApplicationData: 'Application data is not allowed',
  uplinkPolicies: 'This top row shows the uplink forwarding policies of this network',
  downlinkPolicies: 'This bottom row shows the downlink policies of this network',
  gatewayAntennaPlacementLabel: 'Antenna placement',
  gatewayAntennaPlacementDescription: 'Show antenna placement (indoor/outdoor)',
  gatewayAntennaCountLabel: 'Antenna count',
  gatewayFineTimestampsLabel: 'Fine timestamps',
  gatewayFineTimestampsDescription: 'Whether the gateway produces fine timestamps',
  gatewayContactInfoDescription: 'Show means to contact the gateway owner or operator',
  gatewayStatusDescription: 'Show whether the gateway is online or offline',
  gatewayPacketRatesLabel: 'Packet rates',
  gatewayPacketRatesDescription: 'Receive and transmission packet rates',
})

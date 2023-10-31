

import { defineMessages } from 'react-intl'

const messages = defineMessages({
  MACPayload: 'MAC payload',
  devAddr: 'DevAddr',
  fPort: 'FPort',
  fCnt: 'FCnt',
  rawPayload: 'Raw payload',
  txPower: 'Tx Power',
  dataRate: 'Data rate',
  bandwidth: 'Bandwidth',
  metrics: 'Metrics',
  versions: 'Versions',
  snr: 'SNR',
  rssi: 'RSSI',
  sessionKeyId: 'Session key ID',
  selectedMacVersion: 'Selected MAC version',
  rx1Delay: 'Rx1 Delay',
  rx1DataRateIndex: 'Rx1 Data Rate Index',
  rx1Frequency: 'Rx1 Frequency',
  rx2DataRateIndex: 'Rx2 Data Rate Index',
  rx2Frequency: 'Rx2 Frequency',
  class: 'Class',

  // Generic messages
  eventDetails: 'Event details',
  rawEvent: 'Raw event',
  errorOverviewEntry:
    'There was an error and the event cannot be displayed. The raw event can by viewed by clicking this row.',
  dataPreview: 'Data preview',
  dataFormats: 'Data Formats',
  dataFormatsInformation:
    'For more information on event message types, please see our {dataFormatsDocumentationLink} documentation.',
  seeAllActivity: 'See all activity',
  syntheticEvent:
    'Note: This meta event did not originate from the network but was generated automatically by the Console. It is not related to any end device or gateway activity.',
  eventsTruncated:
    'Old events have been truncated to save memory. The current event limit per stream is {limit}.',
  eventUnavailable: 'This event is not available anymore. It was likely truncated to save memory.',
  verboseStream: 'Verbose stream',
  confirmedUplink: 'Confirmed uplink',
})

export default messages

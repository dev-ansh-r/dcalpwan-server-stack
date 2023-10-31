

import React from 'react'

import { getDataRate, getSignalInformation } from '@console/components/events/utils'

import PropTypes from '@ttn-lw/lib/prop-types'
import sharedMessages from '@ttn-lw/lib/shared-messages'

import messages from '../messages'

import DescriptionList from './shared/description-list'
import JSONPayload from './shared/json-payload'

const ApplicationUplinkPreview = React.memo(({ event }) => {
  const { data, identifiers } = event
  const deviceIds = identifiers[0].device_ids
  const { snr, rssi } = getSignalInformation(data)
  const dataRate = getDataRate(data)

  return (
    <DescriptionList>
      <DescriptionList.Byte title={messages.devAddr} data={deviceIds.dev_addr} />
      {'decoded_payload' in data ? (
        <DescriptionList.Item title={sharedMessages.payload}>
          <JSONPayload data={data.decoded_payload} />
          {data.frm_payload && (
            <DescriptionList.Byte key="frm_payload" data={data.frm_payload} convertToHex />
          )}
        </DescriptionList.Item>
      ) : (
        <DescriptionList.Byte title={messages.payload} data={data.frm_payload} convertToHex />
      )}
      <DescriptionList.Item title={messages.fPort} data={data.f_port} />
      <DescriptionList.Item title={messages.dataRate} data={dataRate} />
      <DescriptionList.Item title={messages.snr} data={snr} />
      <DescriptionList.Item title={messages.rssi} data={rssi} />
    </DescriptionList>
  )
})

ApplicationUplinkPreview.propTypes = {
  event: PropTypes.event.isRequired,
}

export default ApplicationUplinkPreview

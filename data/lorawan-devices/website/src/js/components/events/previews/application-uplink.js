
import React from 'react'

import PropTypes from 'prop-types'
import getByPath from '../../../lib/get-by-path'

import DescriptionList from './shared/description-list'
import JSONPayload from './shared/json-payload'

const ApplicationUplinkPreview = React.memo(({ event }) => {
  const { data, identifiers } = event
  const deviceIds = identifiers[0].device_ids
  let snr

  if ('rx_metadata' in data) {
    snr = data.uplink_message.rx_metadata[0].snr
  }

  const bandwidth = getByPath(data.uplink_message, 'settings.data_rate.lora.bandwidth')

  return (
    <DescriptionList>
      <DescriptionList.Byte title={'DevAddr'} data={deviceIds.dev_addr} />
      {'decoded_payload' in data.uplink_message && Boolean(data.uplink_message.decoded_payload) ? (
        <DescriptionList.Item title={'Payload'}>
          <JSONPayload data={data.uplink_message.decoded_payload} />
          {data.uplink_message.frm_payload && (
            <DescriptionList.Byte
              key="frm_payload"
              data={data.uplink_message.frm_payload}
              convertToHex
            />
          )}
        </DescriptionList.Item>
      ) : (
        <DescriptionList.Byte
          title={'MAC payload'}
          data={data.uplink_message.frm_payload}
          convertToHex
        />
      )}
      <DescriptionList.Item title={'FPort'} data={data.uplink_message.f_port} />
      <DescriptionList.Item title={'SNR'} data={snr} />
      <DescriptionList.Item title={'Bandwidth'} data={bandwidth} />
    </DescriptionList>
  )
})

ApplicationUplinkPreview.propTypes = {
  event: PropTypes.shape({}).isRequired,
}

export default ApplicationUplinkPreview

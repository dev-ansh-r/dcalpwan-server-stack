

import React from 'react'

import PropTypes from '@ttn-lw/lib/prop-types'
import sharedMessages from '@ttn-lw/lib/shared-messages'

import messages from '../messages'

import DescriptionList from './shared/description-list'
import JSONPayload from './shared/json-payload'

const ApplicationDownlinkPreview = React.memo(({ event }) => {
  const { data, identifiers } = event
  const deviceIds = identifiers[0].device_ids

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
      <DescriptionList.Item title={messages.fPort}>{data.f_port}</DescriptionList.Item>
    </DescriptionList>
  )
})

ApplicationDownlinkPreview.propTypes = {
  event: PropTypes.event.isRequired,
}

export default ApplicationDownlinkPreview

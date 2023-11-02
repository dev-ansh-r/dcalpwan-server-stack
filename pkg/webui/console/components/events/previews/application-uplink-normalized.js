

import React from 'react'

import PropTypes from '@ttn-lw/lib/prop-types'
import sharedMessages from '@ttn-lw/lib/shared-messages'

import messages from '../messages'

import DescriptionList from './shared/description-list'
import JSONPayload from './shared/json-payload'

const ApplicationUplinkNormalizedPreview = React.memo(({ event }) => {
  const { data, identifiers } = event
  const deviceIds = identifiers[0].device_ids

  return (
    <DescriptionList>
      <DescriptionList.Byte title={messages.devAddr} data={deviceIds.dev_addr} />
      {data.normalized_payload.soil && (
        <DescriptionList.Item title={sharedMessages.normalizedPayloadSoil}>
          <JSONPayload data={data.normalized_payload.soil} />
        </DescriptionList.Item>
      )}
      {data.normalized_payload.air && (
        <DescriptionList.Item title={sharedMessages.normalizedPayloadAir}>
          <JSONPayload data={data.normalized_payload.air} />
        </DescriptionList.Item>
      )}
      {data.normalized_payload.wind && (
        <DescriptionList.Item title={sharedMessages.normalizedPayloadWind}>
          <JSONPayload data={data.normalized_payload.wind} />
        </DescriptionList.Item>
      )}
    </DescriptionList>
  )
})

ApplicationUplinkNormalizedPreview.propTypes = {
  event: PropTypes.event.isRequired,
}

export default ApplicationUplinkNormalizedPreview

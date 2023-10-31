

import React from 'react'

import PropTypes from '@ttn-lw/lib/prop-types'
import getByPath from '@ttn-lw/lib/get-by-path'

import messages from '../messages'

import DescriptionList from './shared/description-list'
import JSONPayload from './shared/json-payload'

const GatewayStatusPreview = React.memo(({ event }) => {
  const metrics = getByPath(event, 'data.metrics')
  const versions = getByPath(event, 'data.versions')

  return (
    <DescriptionList>
      {Boolean(metrics) && (
        <DescriptionList.Item title={messages.metrics}>
          <JSONPayload data={metrics} />
        </DescriptionList.Item>
      )}
      {Boolean(versions) && (
        <DescriptionList.Item title={messages.versions}>
          <JSONPayload data={versions} />
        </DescriptionList.Item>
      )}
    </DescriptionList>
  )
})

GatewayStatusPreview.propTypes = {
  event: PropTypes.event.isRequired,
}

export default GatewayStatusPreview



import React from 'react'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import PropTypes from '@ttn-lw/lib/prop-types'
import getByPath from '@ttn-lw/lib/get-by-path'

import messages from '../messages'

import DescriptionList from './shared/description-list'

const JoinResponsePreview = React.memo(({ event }) => {
  const { identifiers, data } = event
  const ids = identifiers[0].device_ids

  const sessionKeyId = getByPath(data, 'session_keys.session_key_id')

  return (
    <DescriptionList>
      <DescriptionList.Byte title={messages.devAddr} data={ids.dev_addr} />
      <DescriptionList.Byte title={sharedMessages.joinEUI} data={ids.join_eui} />
      <DescriptionList.Byte title={sharedMessages.devEUI} data={ids.dev_eui} />
      <DescriptionList.Byte title={messages.sessionKeyId} data={sessionKeyId} convertToHex />
    </DescriptionList>
  )
})

JoinResponsePreview.propTypes = {
  event: PropTypes.event.isRequired,
}

export default JoinResponsePreview

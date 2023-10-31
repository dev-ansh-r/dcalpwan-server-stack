

import React from 'react'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import PropTypes from '@ttn-lw/lib/prop-types'

import messages from '../messages'

import DescriptionList from './shared/description-list'

const JoinRequestPreview = React.memo(({ event }) => {
  const { identifiers } = event
  const ids = identifiers[0].device_ids

  const selectedMacVersion = event.data.selected_mac_version

  return (
    <DescriptionList>
      <DescriptionList.Byte title={messages.devAddr} data={ids.dev_addr} />
      <DescriptionList.Byte title={sharedMessages.joinEUI} data={ids.join_eui} />
      <DescriptionList.Byte title={sharedMessages.devEUI} data={ids.dev_eui} />
      <DescriptionList.Item title={messages.selectedMacVersion} data={selectedMacVersion} />
    </DescriptionList>
  )
})

JoinRequestPreview.propTypes = {
  event: PropTypes.event.isRequired,
}

export default JoinRequestPreview

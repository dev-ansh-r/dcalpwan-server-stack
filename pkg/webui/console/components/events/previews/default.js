

import React from 'react'

import PropTypes from '@ttn-lw/lib/prop-types'

import messages from '../messages'

import DescriptionList from './shared/description-list'

const DefaultPreview = React.memo(({ event }) => {
  const { identifiers } = event

  if (identifiers && 'device_ids' in identifiers[0]) {
    return (
      <DescriptionList>
        <DescriptionList.Byte title={messages.devAddr} data={identifiers[0].device_ids.dev_addr} />
      </DescriptionList>
    )
  }

  return null
})

DefaultPreview.propTypes = {
  event: PropTypes.event.isRequired,
}

export default DefaultPreview

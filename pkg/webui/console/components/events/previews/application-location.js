

import React from 'react'

import PropTypes from '@ttn-lw/lib/prop-types'
import sharedMessages from '@ttn-lw/lib/shared-messages'

import DescriptionList from './shared/description-list'

const ApplicationLocationPreview = React.memo(({ event }) => {
  const { data } = event
  const { latitude, longitude, altitude, accuracy, source } = data.location

  return (
    <DescriptionList>
      <DescriptionList.Item title={sharedMessages.latitude} data={latitude} />
      <DescriptionList.Item title={sharedMessages.longitude} data={longitude} />
      <DescriptionList.Item title={sharedMessages.altitude} data={altitude} />
      <DescriptionList.Item title={sharedMessages.accuracy} data={accuracy} />
      {source && (
        <DescriptionList.Item
          title={sharedMessages.source}
          data={source.replace(/^(SOURCE_)/, '')}
        />
      )}
    </DescriptionList>
  )
})

ApplicationLocationPreview.propTypes = {
  event: PropTypes.event.isRequired,
}

export default ApplicationLocationPreview

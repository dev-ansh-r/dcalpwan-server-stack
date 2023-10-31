

import React from 'react'

import PropTypes from '@ttn-lw/lib/prop-types'

import { getApplicationUpMessage, getPreviewComponentByApplicationUpMessage } from '../utils'

const ApplicationUpPreview = React.memo(({ event }) => {
  const { data } = event
  const upMessage = getApplicationUpMessage(data)
  const PreviewComponent = getPreviewComponentByApplicationUpMessage(upMessage)

  return <PreviewComponent event={{ ...event, data: event.data[upMessage] }} />
})

ApplicationUpPreview.propTypes = {
  event: PropTypes.event.isRequired,
}

export default ApplicationUpPreview

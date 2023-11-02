
import React, { useEffect } from 'react'

import Notification from '@ttn-lw/components/notification'

import { isBackend, toMessageProps, ingestError } from '@ttn-lw/lib/errors/utils'
import PropTypes from '@ttn-lw/lib/prop-types'

const ErrorNotification = ({ content, title, details, noIngest, ...rest }) => {
  const message = toMessageProps(content)
  let passedDetails = details

  useEffect(() => {
    if (!noIngest) {
      ingestError(details || content, {
        ingestedBy: 'ErrorNotification',
      })
    }
  }, [content, details, noIngest])

  if (isBackend(content) && !details) {
    passedDetails = content
  }
  return (
    <Notification
      error
      content={message.content}
      title={title || message.title}
      messageValues={message.values}
      details={passedDetails}
      data-test-id="error-notification"
      {...rest}
    />
  )
}

ErrorNotification.propTypes = {
  content: PropTypes.oneOfType([PropTypes.message, PropTypes.error, PropTypes.string]).isRequired,
  details: PropTypes.oneOfType([PropTypes.string, PropTypes.shape({})]),
  noIngest: PropTypes.bool,
  title: PropTypes.message,
}

ErrorNotification.defaultProps = {
  details: undefined,
  noIngest: false,
  title: undefined,
}

export default ErrorNotification

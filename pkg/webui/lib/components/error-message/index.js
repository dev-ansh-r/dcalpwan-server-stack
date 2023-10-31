

import React from 'react'

import Message from '@ttn-lw/lib/components/message'

import PropTypes from '@ttn-lw/lib/prop-types'
import { toMessageProps, isBackend } from '@ttn-lw/lib/errors/utils'

const ErrorMessage = ({ content, withRootCause, useTopmost, className, ...rest }) => {
  const baseProps = {
    className,
    firstToUpper: true,
    convertBackticks: Boolean(isBackend(content)),
    ...rest,
  }

  const messageProps = toMessageProps(content, true)

  if (withRootCause && messageProps.length > 1) {
    return (
      <span className={baseProps.className}>
        <Message {...baseProps} {...messageProps[1]} />:{' '}
        <Message {...baseProps} {...messageProps[0]} />
      </span>
    )
  }

  const index = useTopmost ? messageProps.length - 1 : 0

  return <Message {...baseProps} {...messageProps[index]} />
}

ErrorMessage.propTypes = {
  className: PropTypes.string,
  /**
   * Content contains the error data. It will be marshalled into a `react-intl`
   * message in case of backend errors and then output as such. Can also
   * be a usual message type, in case of frontend-defined errors.
   */
  content: PropTypes.error.isRequired,
  useTopmost: PropTypes.bool,
  withRootCause: PropTypes.bool,
}

ErrorMessage.defaultProps = {
  className: undefined,
  useTopmost: false,
  withRootCause: false,
}

export default ErrorMessage

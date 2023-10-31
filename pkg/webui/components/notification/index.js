

import React from 'react'
import classnames from 'classnames'

import Icon from '@ttn-lw/components/icon'

import Message from '@ttn-lw/lib/components/message'

import PropTypes from '@ttn-lw/lib/prop-types'

import Details from './details'

import style from './notification.styl'

const Notification = ({
  content,
  className,
  title,
  error,
  warning,
  details,
  info,
  small,
  success,
  messageValues = {},
  children,
  'data-test-id': dataTestId,
}) => {
  const classname = classnames(style.notification, className, {
    [style.error]: error,
    [style.warning]: warning,
    [style.info]: info,
    [style.small]: small,
    [style.success]: success,
    [style.withDetails]: Boolean(details),
  })

  let icon = 'info'
  if (error) {
    icon = 'error'
  } else if (warning) {
    icon = 'warning'
  }

  return (
    <div className={classname} data-test-id={dataTestId}>
      <div className={style.container}>
        <Icon className={style.icon} icon={icon} large={!small} />
        <div className={style.content}>
          {title && <Message className={style.title} content={title} component="h4" />}
          <div className={style.message}>
            <Message content={content} values={messageValues} firstToUpper />
            {children}
          </div>
        </div>
      </div>
      {Boolean(details) && <Details className={style.details} details={details} isError={error} />}
    </div>
  )
}

Notification.propTypes = {
  children: PropTypes.node,
  className: PropTypes.string,
  content: PropTypes.oneOfType([PropTypes.message, PropTypes.error, PropTypes.string]),
  'data-test-id': PropTypes.string,
  details: PropTypes.error,
  error: PropTypes.bool,
  info: PropTypes.bool,
  messageValues: PropTypes.shape({}),
  small: PropTypes.bool,
  success: PropTypes.bool,
  title: PropTypes.message,
  warning: PropTypes.bool,
}

Notification.defaultProps = {
  children: undefined,
  className: undefined,
  content: undefined,
  'data-test-id': 'notification',
  error: false,
  info: false,
  small: false,
  title: '',
  warning: false,
  success: false,
  messageValues: undefined,
  details: undefined,
}

export default Notification

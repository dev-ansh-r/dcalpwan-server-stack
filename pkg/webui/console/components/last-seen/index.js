

import React from 'react'
import classnames from 'classnames'

import Status from '@ttn-lw/components/status'

import DateTime from '@ttn-lw/lib/components/date-time'
import Message from '@ttn-lw/lib/components/message'

import PropTypes from '@ttn-lw/lib/prop-types'
import sharedMessages from '@ttn-lw/lib/shared-messages'

import style from './last-seen.styl'

const computeDeltaInSeconds = (from, to) => {
  // Avoid situations when server clock is ahead of the browser clock.
  if (from > to) {
    return 0
  }

  return Math.floor((from - to) / 1000)
}

const LastSeen = React.forwardRef((props, ref) => {
  const {
    className,
    lastSeen,
    short,
    updateIntervalInSeconds,
    children,
    flipped,
    message,
    status,
    noTitle,
  } = props

  return (
    <Status status={status} pulseTrigger={lastSeen} flipped={flipped} ref={ref}>
      <div className={classnames(className, style.container)}>
        {!short && <Message className={style.message} content={message} />}
        <DateTime.Relative
          value={lastSeen}
          computeDelta={computeDeltaInSeconds}
          firstToLower={!short}
          updateIntervalInSeconds={updateIntervalInSeconds}
          relativeTimeStyle={short ? 'short' : undefined}
          noTitle={noTitle}
        />
      </div>
      {children}
    </Status>
  )
})

LastSeen.propTypes = {
  children: PropTypes.node,
  className: PropTypes.string,
  flipped: PropTypes.bool,
  lastSeen: PropTypes.oneOfType([
    PropTypes.string,
    PropTypes.number, // Support timestamps.
    PropTypes.instanceOf(Date),
  ]).isRequired,
  message: PropTypes.message,
  noTitle: PropTypes.bool,
  short: PropTypes.bool,
  status: PropTypes.oneOf(['good', 'bad', 'mediocre', 'unknown']),
  updateIntervalInSeconds: PropTypes.number,
}

LastSeen.defaultProps = {
  children: undefined,
  className: undefined,
  flipped: false,
  updateIntervalInSeconds: undefined,
  short: false,
  status: 'good',
  message: sharedMessages.lastSeen,
  noTitle: false,
}

export default LastSeen

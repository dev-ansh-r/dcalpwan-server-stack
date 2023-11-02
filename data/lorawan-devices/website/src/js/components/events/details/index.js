

import React from 'react'
import classnames from 'classnames'
import PropTypes from 'prop-types'

import { getEventId } from '../utils'
import RawEventDetails from './raw'

import style from './details.styl'

const EventDetails = ({ className, children, event }) => {
  const hasChildren = Boolean(children)

  return (
    <div className={classnames(className, style.details)}>
      {!hasChildren ? (
        <RawEventDetails className={style.codeEditor} details={event} id={getEventId(event)} />
      ) : (
        children
      )}
    </div>
  )
}

EventDetails.propTypes = {
  children: PropTypes.oneOfType([PropTypes.arrayOf(PropTypes.node), PropTypes.node]),
  className: PropTypes.string,
  event: PropTypes.shape({}),
}

EventDetails.defaultProps = {
  children: null,
  className: undefined,
  event: undefined,
}

export default EventDetails

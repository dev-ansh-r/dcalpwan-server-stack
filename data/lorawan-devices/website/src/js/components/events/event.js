
import React, { useMemo, useCallback } from 'react'
import classnames from 'classnames'
import PropTypes from 'prop-types'

import { getPreviewComponent, getEventIconByName } from './utils'

import Icon from '../icon'

import style from './events.styl'

const formatTime = (isoString) => {
  const time = new Date(isoString)
  return `${('0' + time.getHours()).slice(-2)}:${('0' + time.getMinutes()).slice(-2)}:${(
    '0' + time.getSeconds()
  ).slice(-2)}`
}

const Event = React.memo(({ event, rowStyle, onRowClick, eventId, active }) => {
  const icon = useMemo(() => getEventIconByName(event.name), [event.name])
  const typeValue = 'Forward uplink data message'
  const PreviewComponent = useMemo(() => {
    return getPreviewComponent(event)
  }, [event])

  const handleRowClick = useCallback(() => {
    onRowClick(eventId)
  }, [eventId, onRowClick])
  const eventClasses = classnames(style.event, {
    [style.active]: active,
    [style.synthetic]: event.isSynthetic,
  })

  return (
    <li className={eventClasses} style={rowStyle} onClick={handleRowClick}>
      <div className={style.cellTime} title={`${event.time}: ${typeValue}`}>
        <div>
          <span>{formatTime(event.time)}</span>
        </div>
      </div>
      <div className={style.cellType} title={typeValue}>
        <span>{typeValue}</span>
      </div>
      <div className={style.cellPreview}>
        <PreviewComponent event={event} />
      </div>
    </li>
  )
})

Event.propTypes = {
  active: PropTypes.bool,
  event: PropTypes.shape({}).isRequired,
  eventId: PropTypes.string.isRequired,
  onRowClick: PropTypes.func,
  rowStyle: PropTypes.shape({}),
}

Event.defaultProps = {
  active: false,
  onRowClick: () => null,
  rowStyle: undefined,
}

export default Event

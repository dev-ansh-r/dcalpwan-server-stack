

import React from 'react'
import { FixedSizeList as List } from 'react-window'
import AutoSizer from 'react-virtualized-auto-sizer'
import PropTypes from 'prop-types'

import Event from './event'

import { getEventId } from './utils'

import style from './events.styl'

const EmptyMessage = () => (
  <div className={style.emptyMessageContainer}>
    <div className={style.emptyMessageContent}>No events</div>
  </div>
)

const EventsList = React.memo(({ events, scoped, onRowClick, activeId }) => {
  if (!events.length) {
    return <EmptyMessage />
  }
  return (
    <AutoSizer>
      {({ height, width }) => (
        <ol>
          <List
            height={height}
            width={width}
            itemCount={events.length}
            itemSize={40}
            overscanCount={25}
          >
            {({ index, style }) => {
              const eventId = getEventId(events[index])
              return (
                <Event
                  event={events[index]}
                  eventId={eventId}
                  key={eventId}
                  scoped={scoped}
                  rowStyle={style}
                  onRowClick={onRowClick}
                  index={index}
                  active={eventId === activeId}
                />
              )
            }}
          </List>
        </ol>
      )}
    </AutoSizer>
  )
})

EventsList.propTypes = {
  activeId: PropTypes.string,
  events: PropTypes.arrayOf(PropTypes.shape({})).isRequired,
  onRowClick: PropTypes.func.isRequired,
  scoped: PropTypes.bool,
}

EventsList.defaultProps = {
  activeId: undefined,
  scoped: false,
}

export { EventsList as default, EmptyMessage }

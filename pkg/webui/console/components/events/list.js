

import React from 'react'
import { FixedSizeList as List } from 'react-window'
import AutoSizer from 'react-virtualized-auto-sizer'

import Message from '@ttn-lw/lib/components/message'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import PropTypes from '@ttn-lw/lib/prop-types'

import Event from './event'
import { getEventId } from './utils'

import style from './events.styl'

const EmptyMessage = ({ entityId }) => (
  <div className={style.emptyMessageContainer}>
    <Message
      className={style.emptyMessageContent}
      content={sharedMessages.noEvents}
      values={{ pre: content => <pre key="entity-id">{content}</pre>, entityId }}
    />
  </div>
)

EmptyMessage.propTypes = {
  entityId: PropTypes.string.isRequired,
}

const EventsList = React.memo(({ events, scoped, entityId, onRowClick, activeId }) => {
  if (!events.length) {
    return <EmptyMessage entityId={entityId} />
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
  entityId: PropTypes.string.isRequired,
  events: PropTypes.events.isRequired,
  onRowClick: PropTypes.func.isRequired,
  scoped: PropTypes.bool,
}

EventsList.defaultProps = {
  activeId: undefined,
  scoped: false,
}

export { EventsList as default, EmptyMessage }

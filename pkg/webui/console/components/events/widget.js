

import React from 'react'
import classnames from 'classnames'

import Link from '@ttn-lw/components/link'
import WidgetContainer from '@ttn-lw/components/widget-container'
import Status from '@ttn-lw/components/status'

import Message from '@ttn-lw/lib/components/message'

import PropTypes from '@ttn-lw/lib/prop-types'
import sharedMessages from '@ttn-lw/lib/shared-messages'

import Event from './event'
import { EmptyMessage } from './list'
import { getEventId } from './utils'
import m from './messages'

import style from './events.styl'

const EventsWidget = ({ toAllUrl, className, events, scoped, entityId }) => {
  const title = (
    <Status flipped status="good" pulseTrigger={events.length > 0 ? events[0].time : ''}>
      <Message content={sharedMessages.liveData} className={style.widgetHeaderTitle} />
    </Status>
  )
  return (
    <WidgetContainer
      title={title}
      toAllUrl={toAllUrl}
      linkMessage={m.seeAllActivity}
      className={className}
    >
      <Link className={style.bodyLink} to={toAllUrl}>
        <div className={classnames(style.body, style.widgetContainer)} data-test-id="events-widget">
          {events.length === 0 ? (
            <EmptyMessage entityId={entityId} />
          ) : (
            <ol>
              {events.slice(0, 6).map(event => {
                const eventId = getEventId(event)
                return (
                  <Event event={event} eventId={eventId} key={eventId} scoped={scoped} widget />
                )
              })}
            </ol>
          )}
        </div>
      </Link>
    </WidgetContainer>
  )
}

EventsWidget.propTypes = {
  className: PropTypes.string,
  entityId: PropTypes.string.isRequired,
  events: PropTypes.events.isRequired,
  scoped: PropTypes.bool,
  toAllUrl: PropTypes.string.isRequired,
}
EventsWidget.defaultProps = { className: undefined, scoped: false }

export default EventsWidget

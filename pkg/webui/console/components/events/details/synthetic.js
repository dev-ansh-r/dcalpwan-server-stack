
import React from 'react'

import Notification from '@ttn-lw/components/notification'

import Message from '@ttn-lw/lib/components/message'

import PropTypes from '@ttn-lw/lib/prop-types'

import { eventMessages } from '@console/lib/events/definitions'

import { getEventId } from '../utils'
import messages from '../messages'

import RawEventDetails from './raw'

import style from './details.styl'

const SyntheticEventDetails = ({ event }) => (
  <>
    <Notification content={messages.syntheticEvent} info small />
    {eventMessages[`${event.name}:details`] && (
      <div>
        <Message
          component="h3"
          content={eventMessages[`${event.name}:type`]}
          className={style.eventType}
        />
        <Message
          component="p"
          content={eventMessages[`${event.name}:details`]}
          className={style.eventDescription}
        />
      </div>
    )}
    <Message component="h4" content={messages.rawEvent} />
    <RawEventDetails className={style.codeEditor} details={event} id={getEventId(event)} />
  </>
)

SyntheticEventDetails.propTypes = {
  event: PropTypes.event.isRequired,
}

export default SyntheticEventDetails

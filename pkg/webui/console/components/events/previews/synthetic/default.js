

import React from 'react'

import Message from '@ttn-lw/lib/components/message'

import PropTypes from '@ttn-lw/lib/prop-types'

import { eventMessages } from '@console/lib/events/definitions'

import style from '../previews.styl'

const SyntheticEventPreview = React.memo(({ event }) => (
  <Message className={style.plainText} content={eventMessages[`${event.name}:preview`]} />
))

SyntheticEventPreview.propTypes = {
  event: PropTypes.event.isRequired,
}

export default SyntheticEventPreview



import React from 'react'

import ErrorMessage from '@ttn-lw/lib/components/error-message'

import PropTypes from '@ttn-lw/lib/prop-types'

import { eventMessages } from '@console/lib/events/definitions'

import style from '../previews.styl'

const SyntheticErrorEventPreview = React.memo(({ event }) => (
  <ErrorMessage
    className={style.plainText}
    content={eventMessages[`${event.name}:preview`]}
    withRootCause
  />
))

SyntheticErrorEventPreview.propTypes = {
  event: PropTypes.event.isRequired,
}

export default SyntheticErrorEventPreview

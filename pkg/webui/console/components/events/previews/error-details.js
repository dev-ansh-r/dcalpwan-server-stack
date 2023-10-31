

import React from 'react'

import ErrorMessage from '@ttn-lw/lib/components/error-message'

import PropTypes from '@ttn-lw/lib/prop-types'

import style from './previews.styl'

const ErrorPreview = React.memo(({ event }) => (
  <ErrorMessage className={style.plainText} content={event.data} withRootCause convertBackticks />
))

ErrorPreview.propTypes = {
  event: PropTypes.event.isRequired,
}

export default ErrorPreview

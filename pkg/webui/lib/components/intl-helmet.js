

import React, { useEffect } from 'react'
import { Helmet } from 'react-helmet'
import { useIntl } from 'react-intl'

import PropTypes from '@ttn-lw/lib/prop-types'
import { warn } from '@ttn-lw/lib/log'

const IntlHelmet = ({ children, values, ...rest }) => {
  const intl = useIntl()

  useEffect(() => {
    if (children) {
      warn(
        `Children of <IntlHelmet /> will not be translated. If you tried to translate head elements with <Message />, use props with message objects instead.`,
      )
    }
  }, [children])

  let translatedRest = {}
  for (const key in rest) {
    let prop = rest[key]
    if (typeof prop === 'object' && prop.id && prop.defaultMessage) {
      const messageValues = values || prop.values || {}
      const translatedMessageValues = {}

      for (const entry in messageValues) {
        const content = messageValues[entry]
        if (typeof content === 'object' && prop.id && prop.defaultMessage) {
          translatedMessageValues[entry] = intl.formatMessage(content)
        } else {
          translatedMessageValues[entry] = messageValues[entry]
        }
      }

      prop = intl.formatMessage(prop, translatedMessageValues)
    }

    translatedRest = {
      ...translatedRest,
      [key]: prop,
    }
  }

  return <Helmet {...translatedRest}>{children}</Helmet>
}

IntlHelmet.propTypes = {
  children: PropTypes.node,
  values: PropTypes.shape({}),
}

IntlHelmet.defaultProps = {
  children: undefined,
  values: undefined,
}

export default IntlHelmet

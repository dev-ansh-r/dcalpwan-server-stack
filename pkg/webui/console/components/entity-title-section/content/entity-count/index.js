

import React from 'react'
import classnames from 'classnames'

import Icon from '@ttn-lw/components/icon'
import Link from '@ttn-lw/components/link'

import Message from '@ttn-lw/lib/components/message'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import PropTypes from '@ttn-lw/lib/prop-types'

import style from './entity-count.styl'

const EntityCount = props => {
  const { icon, keyMessage, value, errored, toAllUrl } = props

  return (
    <Link
      className={classnames(style.container, {
        [style.error]: errored,
      })}
      to={toAllUrl}
      disabled={errored}
    >
      <Icon className={style.icon} icon={icon} textPaddedRight />
      {!errored && <span className={style.value}>{value} </span>}
      <Message
        className={style.message}
        content={errored ? sharedMessages.notAvailable : keyMessage}
        values={{
          count: errored ? undefined : value,
        }}
      />
    </Link>
  )
}

EntityCount.propTypes = {
  errored: PropTypes.bool,
  icon: PropTypes.string.isRequired,
  keyMessage: PropTypes.message.isRequired,
  toAllUrl: PropTypes.string.isRequired,
  value: PropTypes.node,
}

EntityCount.defaultProps = {
  errored: false,
  value: undefined,
}

export default EntityCount

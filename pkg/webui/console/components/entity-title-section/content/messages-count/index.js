

import React from 'react'
import classnames from 'classnames'
import { FormattedNumber } from 'react-intl'

import Icon from '@ttn-lw/components/icon'

import PropTypes from '@ttn-lw/lib/prop-types'

import style from './messages-count.styl'

const MessagesCount = React.forwardRef((props, ref) => {
  const { className, icon, value, iconClassName } = props

  return (
    <div ref={ref} className={classnames(style.container, className)}>
      <Icon className={iconClassName} icon={icon} nudgeUp />
      {typeof value === 'number' ? <FormattedNumber value={value} /> : value}
    </div>
  )
})

MessagesCount.propTypes = {
  className: PropTypes.string,
  icon: PropTypes.string.isRequired,
  iconClassName: PropTypes.string,
  value: PropTypes.node.isRequired,
}

MessagesCount.defaultProps = {
  className: undefined,
  iconClassName: undefined,
}

export default MessagesCount

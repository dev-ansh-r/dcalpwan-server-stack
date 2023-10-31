

import React from 'react'
import classnames from 'classnames'

import Message from '@ttn-lw/lib/components/message'

import PropTypes from '@ttn-lw/lib/prop-types'

import style from './submit-bar.styl'

const SubmitBar = props => {
  const { className, children, align } = props

  const cls = classnames(className, style.bar, style[`bar-${align}`])

  return <div className={cls}>{children}</div>
}

const SubmitBarMessage = ({ className, ...rest }) => (
  <Message {...rest} className={classnames(className, style.barMessage)} />
)

SubmitBar.propTypes = {
  align: PropTypes.oneOf(['start', 'end', 'between', 'around']),
  children: PropTypes.node,
  className: PropTypes.string,
}

SubmitBar.defaultProps = {
  align: 'between',
  className: undefined,
  children: undefined,
}

SubmitBarMessage.propTypes = {
  className: PropTypes.string,
}

SubmitBarMessage.defaultProps = {
  className: undefined,
}

SubmitBar.Message = SubmitBarMessage

export default SubmitBar

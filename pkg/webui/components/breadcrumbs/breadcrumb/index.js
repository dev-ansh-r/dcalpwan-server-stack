

import React from 'react'
import classnames from 'classnames'

import Link from '@ttn-lw/components/link'

import Message from '@ttn-lw/lib/components/message'

import PropTypes from '@ttn-lw/lib/prop-types'

import style from './breadcrumb.styl'

const Breadcrumb = ({ className, path, content, isLast }) => {
  const isRawText = typeof content === 'string' || typeof content === 'number'
  let Component
  let componentProps
  if (!isLast) {
    Component = Link
    componentProps = { className: classnames(className, style.link), to: path, secondary: true }
  } else {
    Component = 'span'
    componentProps = { className: classnames(className, style.last) }
  }

  return (
    <Component {...componentProps}>
      {isRawText ? <span>{content}</span> : <Message content={content} />}
    </Component>
  )
}

Breadcrumb.propTypes = {
  className: PropTypes.string,
  /** The content of the breadcrumb. */
  content: PropTypes.message.isRequired,
  /** The flag for rendering last breadcrumb as plain text. */
  isLast: PropTypes.bool,
  /** The path for a breadcrumb. */
  path: PropTypes.string.isRequired,
}

Breadcrumb.defaultProps = {
  className: undefined,
  isLast: false,
}

export default Breadcrumb

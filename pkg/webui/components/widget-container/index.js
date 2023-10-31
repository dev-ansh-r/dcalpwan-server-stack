

import React from 'react'
import classnames from 'classnames'

import Link from '@ttn-lw/components/link'

import Message from '@ttn-lw/lib/components/message'

import PropTypes from '@ttn-lw/lib/prop-types'

import style from './widget-container.styl'

const WidgetContainer = ({ children, title, toAllUrl, linkMessage, className }) => (
  <aside className={classnames(style.container, className)}>
    <div className={style.header} role="heading">
      {typeof title === 'object' && 'id' in title ? (
        <Message content={title} className={style.headerTitle} />
      ) : (
        title
      )}
      <Link className={style.seeAllLink} secondary to={toAllUrl}>
        <Message content={linkMessage} /> →
      </Link>
    </div>
    <div className={style.body}>{children}</div>
  </aside>
)

WidgetContainer.propTypes = {
  children: PropTypes.node.isRequired,
  className: PropTypes.string,
  linkMessage: PropTypes.message.isRequired,
  title: PropTypes.oneOfType([PropTypes.message, PropTypes.node]).isRequired,
  toAllUrl: PropTypes.string.isRequired,
}

WidgetContainer.defaultProps = {
  className: undefined,
}

export default WidgetContainer

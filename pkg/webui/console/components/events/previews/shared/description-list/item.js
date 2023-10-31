

import React from 'react'
import classnames from 'classnames'

import Message from '@ttn-lw/lib/components/message'

import PropTypes from '@ttn-lw/lib/prop-types'

import style from './description-list.styl'

const DescriptionListItem = props => {
  const { className, children, data, title, highlight } = props
  const content = children || data

  if (!Boolean(content)) {
    return null
  }

  if (!Boolean(title)) {
    return (
      <div className={classnames(className, style.container)}>
        <div className={classnames(style.value, style.onlyValue)}>{content}</div>
      </div>
    )
  }

  return (
    <div className={classnames(className, style.container)}>
      <dt className={style.term}>
        <Message content={title} />
      </dt>
      <dd className={classnames(style.value, { [style.highlight]: highlight })}>{content}</dd>
    </div>
  )
}

DescriptionListItem.propTypes = {
  children: PropTypes.oneOfType([PropTypes.arrayOf(PropTypes.node), PropTypes.node]),
  className: PropTypes.string,
  data: PropTypes.oneOfType([PropTypes.string, PropTypes.number]),
  highlight: PropTypes.bool,
  title: PropTypes.message,
}

DescriptionListItem.defaultProps = {
  children: undefined,
  data: undefined,
  className: undefined,
  highlight: false,
  title: undefined,
}

export default DescriptionListItem

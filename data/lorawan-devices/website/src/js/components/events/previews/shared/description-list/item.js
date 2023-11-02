
import React from 'react'
import classnames from 'classnames'
import PropTypes from 'prop-types'

import style from './description-list.styl'

const DescriptionListItem = (props) => {
  const { className, children, data, title } = props
  const content = children || data

  if (!Boolean(content)) {
    return null
  }

  if (!Boolean(title)) {
    return (
      <div className={classnames(className, style.container)}>
        <div className={style.value}>{content}</div>
      </div>
    )
  }

  return (
    <div className={classnames(className, style.container)}>
      <dt className={style.term}>
        <span>{title}</span>
      </dt>
      <dd className={style.value}>{content}</dd>
    </div>
  )
}

DescriptionListItem.propTypes = {
  children: PropTypes.oneOfType([PropTypes.arrayOf(PropTypes.node), PropTypes.node]),
  className: PropTypes.string,
  data: PropTypes.oneOfType([PropTypes.string, PropTypes.number]),
  title: PropTypes.string,
}

DescriptionListItem.defaultProps = {
  children: undefined,
  data: undefined,
  className: undefined,
  title: undefined,
}

export default DescriptionListItem

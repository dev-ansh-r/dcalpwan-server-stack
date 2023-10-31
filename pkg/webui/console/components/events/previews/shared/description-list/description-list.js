

import React from 'react'
import classnames from 'classnames'

import PropTypes from '@ttn-lw/lib/prop-types'

import style from './description-list.styl'

const DescriptionList = props => {
  const { className, children } = props

  return <dl className={classnames(className, style.descriptionList)}>{children}</dl>
}

DescriptionList.propTypes = {
  children: PropTypes.oneOfType([PropTypes.arrayOf(PropTypes.node), PropTypes.node]).isRequired,
  className: PropTypes.string,
}

DescriptionList.defaultProps = {
  className: undefined,
}

export default DescriptionList

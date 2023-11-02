
import React from 'react'
import classnames from 'classnames'

import PropTypes from '@ttn-lw/lib/prop-types'

import style from './field.styl'

const FieldContainer = ({ horizontal, children, className }) => (
  <div className={classnames(style.fieldContainer, className, { [style.horizontal]: horizontal })}>
    {children}
  </div>
)

FieldContainer.propTypes = {
  children: PropTypes.node.isRequired,
  className: PropTypes.string,
  horizontal: PropTypes.bool,
}

FieldContainer.defaultProps = {
  className: undefined,
  horizontal: false,
}

export default FieldContainer



import classNames from 'classnames'
import React from 'react'

import PropTypes from '@ttn-lw/lib/prop-types'

import style from './button.styl'

const ButtonGroup = ({ className, children, align }) => (
  <div className={classNames(style.buttonGroup, className, [style[`align-${align}`]])}>
    {children}
  </div>
)

ButtonGroup.propTypes = {
  align: PropTypes.oneOf(['start', 'end', 'center']),
  children: PropTypes.node.isRequired,
  className: PropTypes.string,
}

ButtonGroup.defaultProps = {
  align: 'start',
  className: undefined,
}

export default ButtonGroup

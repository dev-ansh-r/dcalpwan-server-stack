
import React from 'react'
import PropTypes from 'prop-types'
import classnames from 'classnames'

import style from './icon.styl'

// A map of hardcoded names to their corresponding icons.
// Keep these sorted alphabetically.
const hardcoded = {
  arrow_drop_down: 'arrow_drop_down',
  cancel: 'cancel',
  check_circle: 'check_circle',
  lock: 'lock',
  help: 'help',
}

function Icon({ icon, className, nudgeUp, nudgeDown, small, large, ...rest }) {
  const cls = classnames(style.icon, className, {
    [style.nudgeUp]: nudgeUp,
    [style.nudgeDown]: nudgeDown,
    [style.large]: large,
    [style.small]: small,
  })

  return (
    <span className={cls} {...rest}>
      {hardcoded[icon] || icon}
    </span>
  )
}

Icon.propTypes = {
  className: PropTypes.string,
  icon: PropTypes.string.isRequired,
  large: PropTypes.bool,
  nudgeDown: PropTypes.bool,
  nudgeUp: PropTypes.bool,
  small: PropTypes.bool,
}

Icon.defaultProps = {
  className: '',
  large: false,
  nudgeDown: false,
  nudgeUp: false,
  small: false,
}

export default Icon

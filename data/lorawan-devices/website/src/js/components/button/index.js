

import React from 'react'
import PropTypes from 'prop-types'
import classnames from 'classnames'

import style from './button.styl'

function Button(props) {
  const {
    text,
    onClick,
    type,
    disabled,
    danger,
    naked,
    secondary,
    inline,
    invert,
    subtle,
    busy,
    bold,
    small,
    className,
  } = props

  const cls = classnames(style.button, className, {
    [style.naked]: Boolean(naked),
    [style.danger]: Boolean(danger),
    [style.secondary]: Boolean(secondary),
    [style.subtle]: Boolean(subtle),
    [style.inline]: Boolean(inline),
    [style.bold]: Boolean(bold),
    [style.small]: Boolean(small),
    [style.invert]: Boolean(invert),
  })

  function handleClick(e) {
    if (disabled) {
      return
    }

    onClick(e)
  }

  return (
    <button className={cls} onClick={handleClick} type={type} disabled={busy || disabled}>
      {text}
    </button>
  )
}

Button.propTypes = {
  bold: PropTypes.bool,
  busy: PropTypes.bool,
  className: PropTypes.string,
  danger: PropTypes.bool,
  disabled: PropTypes.bool,
  inline: PropTypes.bool,
  invert: PropTypes.bool,
  naked: PropTypes.bool,
  onClick: PropTypes.func,
  secondary: PropTypes.bool,
  small: PropTypes.bool,
  subtle: PropTypes.bool,
  text: PropTypes.string,
  type: PropTypes.string,
}

Button.defaultProps = {
  bold: false,
  busy: false,
  className: '',
  danger: false,
  disabled: false,
  inline: false,
  invert: false,
  naked: false,
  onClick: () => null,
  secondary: false,
  small: false,
  subtle: false,
  text: '',
  type: 'button',
}

export default Button

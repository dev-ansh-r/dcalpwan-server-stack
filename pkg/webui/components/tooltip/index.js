
import React, { useCallback } from 'react'
import Tippy from '@tippyjs/react'
import classnames from 'classnames'

import PropTypes from '@ttn-lw/lib/prop-types'

import style from './tooltip.styl'

let currentInstance

const popperModifiers = [
  {
    name: 'offset',
    options: {
      offset: [-12, 8],
    },
  },
  {
    name: 'arrow',
    options: {
      element: '.tippy-arrow',
      padding: 12,
    },
  },
]

const Tooltip = props => {
  const {
    appendTo,
    children,
    className,
    content,
    delay,
    hideOnClick,
    interactive,
    onShow,
    placement,
    small,
    trigger,
  } = props

  const handleShow = useCallback(
    instance => {
      if (currentInstance && currentInstance.state && !currentInstance.state.isDestroyed) {
        currentInstance.hide()
      }
      currentInstance = instance

      onShow(instance)
    },
    [onShow],
  )
  return (
    <Tippy
      className={classnames(className, { [style.small]: small })}
      content={content}
      interactive={interactive}
      placement={placement}
      popperOptions={{ modifiers: popperModifiers }}
      delay={delay}
      onShow={handleShow}
      appendTo={appendTo ? appendTo : interactive ? 'parent' : document.body}
      animation="fade"
      duration={250}
      hideOnClick={hideOnClick}
      trigger={trigger}
    >
      {children}
    </Tippy>
  )
}

Tooltip.propTypes = {
  appendTo: PropTypes.oneOfType([PropTypes.shape({}), PropTypes.string]),
  children: PropTypes.node.isRequired,
  className: PropTypes.string,
  content: PropTypes.node.isRequired,
  delay: PropTypes.oneOfType([PropTypes.number, PropTypes.arrayOf(PropTypes.number)]),
  hideOnClick: PropTypes.bool,
  interactive: PropTypes.bool,
  onShow: PropTypes.func,
  placement: PropTypes.oneOf([
    'top',
    'top-start',
    'top-end',
    'right',
    'right-start',
    'right-end',
    'bottom',
    'bottom-start',
    'bottom-end',
    'left',
    'left-start',
    'left-end',
    'auto',
    'auto-start',
    'auto-end',
  ]),
  small: PropTypes.bool,
  trigger: PropTypes.string,
}

Tooltip.defaultProps = {
  appendTo: undefined,
  className: '',
  hideOnClick: true,
  interactive: false,
  placement: 'bottom',
  small: false,
  delay: 300,
  onShow: () => null,
  trigger: 'mouseenter focus',
}

export default Tooltip

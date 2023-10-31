

import React, { useCallback } from 'react'
import PropTypes from 'prop-types'
import classnames from 'classnames'
import { NavLink } from 'react-router-dom'

import style from './tab.styl'

const Tab = props => {
  const {
    className,
    onClick,
    name,
    active = false,
    disabled = false,
    narrow,
    children,
    link,
    exact = true,
    ...rest
  } = props

  const handleClick = useCallback(() => {
    if (!disabled) {
      onClick(name)
    }
  }, [disabled, name, onClick])

  const tabItemClassNames = classnames(className, style.tabItem, {
    [style.tabItemNarrow]: narrow,
    [style.tabItemActive]: !disabled && active,
    [style.tabItemDefault]: !disabled && !active,
    [style.tabItemDisabled]: disabled,
  })

  // There is no support for disabled on anchors in html and hence in
  // `react-router`. So, do not render the link component if the tab is
  // disabled, but render regular tab item instead.
  const canRenderLink = link && !disabled

  const Component = canRenderLink ? NavLink : 'span'
  const componentProps = {
    role: 'button',
    className: tabItemClassNames,
    children,
  }

  if (canRenderLink) {
    componentProps.end = exact
    componentProps.to = link
    componentProps.className = ({ isActive }) =>
      classnames(tabItemClassNames, style.tabItem, {
        [style.tabItemActive]: !disabled && isActive,
      })
  } else {
    componentProps.onClick = handleClick
  }

  return (
    <li {...rest} className={style.tab}>
      <Component {...componentProps} children={children} />
    </li>
  )
}

Tab.propTypes = {
  /** A flag specifying whether the tab is active. */
  active: PropTypes.bool,
  children: PropTypes.node,
  className: PropTypes.string,
  /** A flag specifying whether the tab is disabled. */
  disabled: PropTypes.bool,
  exact: PropTypes.bool,
  link: PropTypes.string,
  /** The name of the tab. */
  name: PropTypes.string.isRequired,
  narrow: PropTypes.bool,
  /**
   * A click handler to be called when the selected tab changes. Passes the
   * name of the new active tab as an argument.
   */
  onClick: PropTypes.func,
}

Tab.defaultProps = {
  children: undefined,
  className: undefined,
  link: undefined,
  onClick: () => null,
  active: false,
  disabled: false,
  narrow: false,
  exact: true,
}

export default Tab

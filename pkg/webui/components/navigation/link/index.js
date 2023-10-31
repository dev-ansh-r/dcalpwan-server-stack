

import React, { useCallback } from 'react'
import classnames from 'classnames'
import { NavLink } from 'react-router-dom'

import Link from '@ttn-lw/components/link'

import PropTypes from '@ttn-lw/lib/prop-types'

import style from './link.styl'

const NavigationLink = ({
  className,
  children,
  path,
  exact,
  activeClassName,
  onClick,
  ...rest
}) => (
  <NavLink
    to={path}
    end={exact}
    className={useCallback(
      ({ isActive }) => classnames(className, style.link, { [activeClassName]: isActive }),
      [className, activeClassName],
    )}
    onClick={onClick}
    {...rest}
  >
    {children}
  </NavLink>
)

const NavigationAnchorLink = ({ className, children, path }) => (
  <Link.BaseAnchor href={path} className={classnames(className, style.link)} children={children} />
)

NavigationLink.propTypes = {
  activeClassName: PropTypes.string,
  children: PropTypes.node,
  className: PropTypes.string,
  /** Boolean flag identifying whether the path shoul be matched exactly. */
  exact: PropTypes.bool,
  onClick: PropTypes.func,
  /** The path for a link. */
  path: PropTypes.string.isRequired,
  /** The name of a css class to be applied on the active tab. */
}

NavigationLink.defaultProps = {
  activeClassName: undefined,
  children: undefined,
  className: undefined,
  exact: false,
  onClick: () => null,
}

NavigationAnchorLink.propTypes = {
  children: PropTypes.node,
  className: PropTypes.string,
  path: PropTypes.string.isRequired,
}

NavigationAnchorLink.defaultProps = {
  className: undefined,
  children: undefined,
}
export { NavigationLink as default, NavigationAnchorLink }

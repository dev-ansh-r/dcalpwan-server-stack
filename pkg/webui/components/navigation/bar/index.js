

import React from 'react'
import classnames from 'classnames'

import Icon from '@ttn-lw/components/icon'

import Message from '@ttn-lw/lib/components/message'

import PropTypes from '@ttn-lw/lib/prop-types'

import NavigationLink from '../link'

import style from './bar.styl'

const NavigationBar = ({ className, children }) => <nav className={className}>{children}</nav>

NavigationBar.propTypes = {
  children: PropTypes.node.isRequired,
  className: PropTypes.string,
}

NavigationBar.defaultProps = {
  className: undefined,
}

const NavigationBarItem = ({ icon, title, className, ...rest }) => (
  <NavigationLink
    {...rest}
    className={classnames(style.link, className)}
    activeClassName={style.linkActive}
  >
    {icon && <Icon icon={icon} className={style.icon} />}
    <Message content={title} />
  </NavigationLink>
)

NavigationBarItem.propTypes = {
  className: PropTypes.string,
  icon: PropTypes.string.isRequired,
  title: PropTypes.message.isRequired,
}

NavigationBarItem.defaultProps = {
  className: undefined,
}

NavigationBar.Item = NavigationBarItem

export default NavigationBar

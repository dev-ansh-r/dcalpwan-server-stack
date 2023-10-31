

import React from 'react'
import classnames from 'classnames'

import PropTypes from '@ttn-lw/lib/prop-types'

import style from './list.styl'

const SideNavigationList = ({ children, className, depth, isExpanded }) => {
  const isRoot = depth === 0
  const listClassNames = classnames(className, style.list, {
    [style.listNested]: !isRoot,
    [style.listExpanded]: isExpanded,
  })
  return <ul className={listClassNames}>{children}</ul>
}

SideNavigationList.propTypes = {
  children: PropTypes.node.isRequired,
  className: PropTypes.string,
  /** The depth of the current list starting at 0 for the root list. */
  depth: PropTypes.number,
  /**
   * A flag specifying whether the side navigation list is expanded or not.
   * Applicable to nested lists.
   */
  isExpanded: PropTypes.bool,
}

SideNavigationList.defaultProps = {
  className: undefined,
  depth: 0,
  isExpanded: false,
}

export default SideNavigationList

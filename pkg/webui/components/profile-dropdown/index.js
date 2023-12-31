

import React from 'react'
import bind from 'autobind-decorator'
import classnames from 'classnames'

import Icon from '@ttn-lw/components/icon'
import Dropdown from '@ttn-lw/components/dropdown'
import ProfilePicture from '@ttn-lw/components/profile-picture'

import PropTypes from '@ttn-lw/lib/prop-types'

import styles from './profile-dropdown.styl'

export default class ProfileDropdown extends React.PureComponent {
  state = {
    expanded: false,
  }

  @bind
  showDropdown() {
    document.addEventListener('mousedown', this.handleClickOutside)
    this.setState({ expanded: true })
  }

  @bind
  hideDropdown() {
    document.removeEventListener('mousedown', this.handleClickOutside)
    this.setState({ expanded: false })
  }

  @bind
  handleClickOutside(e) {
    if (!this.node.contains(e.target)) {
      this.hideDropdown()
    }
  }

  @bind
  toggleDropdown() {
    let { expanded } = this.state
    expanded = !expanded
    if (expanded) {
      this.showDropdown()
    } else {
      this.hideDropdown()
    }
  }

  @bind
  ref(node) {
    this.node = node
  }

  render() {
    const { userName, className, children, profilePicture, ...rest } = this.props

    return (
      <div
        className={classnames(styles.container, className)}
        onClick={this.toggleDropdown}
        onKeyPress={this.toggleDropdown}
        ref={this.ref}
        tabIndex="0"
        role="button"
        {...rest}
      >
        <ProfilePicture profilePicture={profilePicture} className={styles.profilePicture} />
        <span className={styles.id}>{userName}</span>
        <Icon
          className={styles.dropdownIcon}
          icon={this.state.expanded ? 'arrow_drop_up' : 'arrow_drop_down'}
        />
        {this.state.expanded && <Dropdown className={styles.dropdown}>{children}</Dropdown>}
      </div>
    )
  }
}

ProfileDropdown.propTypes = {
  /**
   * A list of items for the dropdown component. See `<Dropdown />`'s `items`
   * proptypes for details.
   */
  children: PropTypes.node.isRequired,
  className: PropTypes.string,
  /** The profile picture of the current user. */
  profilePicture: PropTypes.profilePicture,
  /** The name/id of the current user. */
  userName: PropTypes.string.isRequired,
}

ProfileDropdown.defaultProps = {
  className: undefined,
  profilePicture: undefined,
}

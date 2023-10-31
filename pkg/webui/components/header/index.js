

import React, { useState, useCallback } from 'react'
import classnames from 'classnames'

import hamburgerMenuNormal from '@assets/misc/hamburger-menu-normal.svg'
import hamburgerMenuClose from '@assets/misc/hamburger-menu-close.svg'

import NavigationBar from '@ttn-lw/components/navigation/bar'
import ProfileDropdown from '@ttn-lw/components/profile-dropdown'
import MobileMenu from '@ttn-lw/components/mobile-menu'
import Input from '@ttn-lw/components/input'

import PropTypes from '@ttn-lw/lib/prop-types'

import style from './header.styl'

const Header = ({
  className,
  dropdownItems,
  navigationEntries,
  user,
  searchable,
  logo,
  mobileDropdownItems,
  onLogout,
  onSearchRequest,
  ...rest
}) => {
  const isGuest = !Boolean(user)

  const [mobileMenuOpen, setMobileMenuOpen] = useState(false)
  const handleMobileMenuClick = useCallback(() => {
    setMobileMenuOpen(!mobileMenuOpen)
  }, [mobileMenuOpen])

  const handleMobileMenuItemsClick = useCallback(() => {
    setMobileMenuOpen(false)
  }, [])

  const classNames = classnames(className, style.container, {
    [style.mobileMenuOpen]: mobileMenuOpen,
  })

  const hamburgerGraphic = mobileMenuOpen ? hamburgerMenuClose : hamburgerMenuNormal

  return (
    <header {...rest} className={classNames}>
      <div className={style.bar}>
        <div className={style.left}>
          {logo}
          {!isGuest && <NavigationBar className={style.navList}>{navigationEntries}</NavigationBar>}
        </div>
        {!isGuest && (
          <div className={style.right}>
            {searchable && <Input icon="search" onEnter={onSearchRequest} />}
            <ProfileDropdown
              className={style.profileDropdown}
              userName={user.name || user.ids.user_id}
              data-test-id="profile-dropdown"
              profilePicture={user.profile_picture}
            >
              {dropdownItems}
            </ProfileDropdown>
            <button onClick={handleMobileMenuClick} className={style.mobileMenuButton}>
              <div className={style.hamburger}>
                <img src={hamburgerGraphic} alt="Open Mobile Menu" />
              </div>
            </button>
          </div>
        )}
      </div>
      {mobileMenuOpen && (
        <MobileMenu
          className={style.mobileMenu}
          onItemsClick={handleMobileMenuItemsClick}
          onLogout={onLogout}
          user={user}
        >
          {mobileDropdownItems}
        </MobileMenu>
      )}
    </header>
  )
}

Header.propTypes = {
  /** The classname applied to the component. */
  className: PropTypes.string,
  /** The child node of the dropdown component. */
  dropdownItems: PropTypes.node,
  /** The logo component. */
  logo: PropTypes.node.isRequired,
  /** The child node of the mobile dropdown. */
  mobileDropdownItems: PropTypes.node,
  /** The Child node of the navigation bar. */
  navigationEntries: PropTypes.node,
  /** A handler for when the user used the search input. */
  onLogout: PropTypes.func,
  /** Handler of the search function. */
  onSearchRequest: PropTypes.func,
  /* A flag indicating whether the header has a search input. */
  searchable: PropTypes.bool,
  /**
   * The User object, retrieved from the API. If it is `undefined`, then the
   * guest header is rendered.
   */
  user: PropTypes.user,
}

Header.defaultProps = {
  className: undefined,
  dropdownItems: undefined,
  navigationEntries: undefined,
  mobileDropdownItems: null,
  onSearchRequest: () => null,
  onLogout: () => null,
  searchable: false,
  user: undefined,
}

export default Header

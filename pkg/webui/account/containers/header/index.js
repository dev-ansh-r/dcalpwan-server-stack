

import React, { useCallback } from 'react'
import { useSelector, useDispatch } from 'react-redux'

import HeaderComponent from '@ttn-lw/components/header'
import NavigationBar from '@ttn-lw/components/navigation/bar'
import Dropdown from '@ttn-lw/components/dropdown'

import Logo from '@account/containers/logo'

import PropTypes from '@ttn-lw/lib/prop-types'
import sharedMessages from '@ttn-lw/lib/shared-messages'

import { logout } from '@account/store/actions/user'

import { selectUser } from '@account/store/selectors/user'

const Header = ({ handleSearchRequest, searchable }) => {
  const user = useSelector(selectUser)
  const dispatch = useDispatch()

  const handleLogout = useCallback(() => {
    dispatch(logout())
  }, [dispatch])

  const navigation = [
    {
      title: sharedMessages.overview,
      icon: 'overview',
      path: '/',
      exact: true,
      hidden: false,
    },
    {
      title: sharedMessages.profileSettings,
      icon: 'user',
      path: '/profile-settings',
    },
    {
      title: sharedMessages.sessions,
      icon: 'vpn_key',
      path: '/session-management',
    },
    {
      title: sharedMessages.oauthClients,
      icon: 'oauth_clients',
      path: '/oauth-clients',
    },
    {
      title: sharedMessages.authorizations,
      icon: 'access',
      path: '/client-authorizations',
    },
  ]

  const navigationEntries = (
    <React.Fragment>
      {navigation.map(
        ({ hidden, ...rest }) => !hidden && <NavigationBar.Item {...rest} key={rest.title.id} />,
      )}
    </React.Fragment>
  )

  const dropdownItems = (
    <React.Fragment>
      <Dropdown.Item title={sharedMessages.logout} icon="logout" action={handleLogout} />
    </React.Fragment>
  )

  const mobileDropdownItems = (
    <React.Fragment>
      {navigation.map(
        ({ hidden, ...rest }) => !hidden && <Dropdown.Item {...rest} key={rest.title.id} />,
      )}
    </React.Fragment>
  )

  return (
    <HeaderComponent
      user={user}
      dropdownItems={dropdownItems}
      mobileDropdownItems={mobileDropdownItems}
      navigationEntries={navigationEntries}
      searchable={searchable}
      onSearchRequest={handleSearchRequest}
      onLogout={handleLogout}
      logo={<Logo />}
    />
  )
}

Header.propTypes = {
  /** A handler for when the user used the search input. */
  handleSearchRequest: PropTypes.func,
  /** A flag identifying whether the header should display the search input. */
  searchable: PropTypes.bool,
}

Header.defaultProps = {
  handleSearchRequest: () => null,
  searchable: false,
}

export default Header

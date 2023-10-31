

import React from 'react'
import { action } from '@storybook/addon-actions'

import TtsLogo from '@assets/static/logo.svg'

import Dropdown from '@ttn-lw/components/dropdown'
import NavigationBar from '@ttn-lw/components/navigation/bar'
import Logo from '@ttn-lw/components/logo'
import ExampleLogo from '@ttn-lw/components/logo/story-logo.svg'

import Header from '.'

const user = {
  name: 'kschiffer',
  ids: {
    user_id: 'ksc300',
  },
}

const singleLogo = <Logo logo={{ src: TtsLogo, alt: 'Logo' }} />
const doubleLogo = (
  <Logo
    logo={{ src: TtsLogo, alt: 'Logo' }}
    brandLogo={{ src: ExampleLogo, alt: 'Secondary Logo' }}
  />
)

const navigationEntries = (
  <React.Fragment>
    <NavigationBar.Item title="Overview" icon="overview" path="/overview" />
    <NavigationBar.Item title="Applications" icon="application" path="/application" />
    <NavigationBar.Item title="Gateways" icon="gateway" path="/gateways" />
    <NavigationBar.Item title="Organizations" icon="organization" path="/organization" />
  </React.Fragment>
)

const items = (
  <React.Fragment>
    <Dropdown.Item title="Profile Settings" icon="settings" path="/profile-settings" />
    <Dropdown.Item title="Logout" icon="power_settings_new" path="/logout" />
  </React.Fragment>
)

export default {
  title: 'Header',
}

export const SingleLogo = () => (
  <Header
    dropdownItems={items}
    handleSearchRequest={action('Search')}
    navigationEntries={navigationEntries}
    style={{ margin: '-1rem' }}
    user={user}
    logo={singleLogo}
  />
)

export const DoubleLogo = () => (
  <Header
    dropdownItems={items}
    handleSearchRequest={action('Search')}
    navigationEntries={navigationEntries}
    style={{ margin: '-1rem' }}
    user={user}
    logo={doubleLogo}
  />
)



import React from 'react'
import { withInfo } from '@storybook/addon-info'

import Dropdown from '@ttn-lw/components/dropdown'

import ProfileDropdown from '.'

const handleLogout = () => {
  // eslint-disable-next-line no-console
  console.log('Click')
}

export default {
  title: 'Profile Dropdown',

  decorators: [
    withInfo({
      inline: true,
      header: false,
      source: false,
      propTables: [ProfileDropdown],
    }),
  ],
}

export const Default = () => (
  <div style={{ height: '6rem' }}>
    <ProfileDropdown style={{ marginLeft: '120px' }} userId="johndoe">
      <Dropdown.Item title="Profile Settings" icon="settings" path="/profile-settings" />
      <Dropdown.Item title="Logout" icon="power_settings_new" action={handleLogout} />
    </ProfileDropdown>
  </div>
)

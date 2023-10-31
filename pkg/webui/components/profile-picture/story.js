

import React from 'react'
import { withInfo } from '@storybook/addon-info'

import ProfilePicture from '.'

const pp = {
  sizes: {
    256: 'https://www.gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50',
  },
}

export default {
  title: 'Profile Picture',

  decorators: [
    withInfo({
      inline: true,
      header: false,
      source: false,
      propTables: [ProfilePicture],
    }),
  ],
}

export const Default = () => (
  <div style={{ height: '8rem' }}>
    <ProfilePicture profilePicture={pp} />
  </div>
)



import React from 'react'
import { withInfo } from '@storybook/addon-info'

import Tab from '.'

export default {
  title: 'Tag',

  decorators: [
    withInfo({
      inline: true,
      header: false,
      propTables: [Tab],
    }),
  ],
}

export const Default = () => <Tab content="Info" />

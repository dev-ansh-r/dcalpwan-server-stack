



import React from 'react'
import { withInfo } from '@storybook/addon-info'

import DataSheet from '.'

const testData = [
  {
    header: 'Hardware',
    items: [
      { key: 'Brand', value: 'SemTech' },
      { key: 'Model', value: 'Some Model' },
      { key: 'Hardware Version', value: 'v1.2.5' },
      { key: 'Firmware Version', value: 'v1.1.1' },
    ],
  },
  {
    header: 'Activation Info',
    items: [
      { key: 'Device EUI', value: '1212121212', type: 'byte', sensitive: false },
      {
        key: 'Device EUI with Uint32_t',
        value: '1212121212',
        type: 'byte',
        sensitive: false,
        enableUint32: true,
      },
      { key: 'Join EUI', value: '1212121212', type: 'byte', sensitive: false },
      {
        key: 'Value with Nesting',
        value: 'ae9d78fe9aed8fe',
        type: 'code',
        sensitive: false,
        subItems: [
          { key: 'Application Key', value: 'ae9d78fe9aed8fe', type: 'code', sensitive: true },
          { key: 'Network Key', value: 'ae9d78fe9aed8fe', type: 'code', sensitive: true },
        ],
      },
    ],
  },
]

const containerStyles = {
  maxWidth: '600px',
}

export default {
  title: 'Data Sheet',

  decorators: [
    withInfo({
      inline: true,
      header: false,
      source: true,
      propTables: [DataSheet],
    }),
  ],
}

export const Default = () => (
  <div style={containerStyles}>
    <DataSheet data={testData} />
  </div>
)

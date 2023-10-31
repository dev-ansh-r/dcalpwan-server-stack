

import React from 'react'
import { withInfo } from '@storybook/addon-info'

import ErrorMessage from '.'

const exampleError = {
  code: 2,
  message:
    'error:pkg/assets:http (HTTP error: `` is not a valid ID. Must be at least 2 and at most 36 characters long and may consist of only letters, numbers and dashes. It may not start or end with a dash)',
  details: [
    {
      '@type': 'type.googleapis.com/ttn.lorawan.v3.ErrorDetails',
      namespace: 'pkg/assets',
      name: 'http',
      message_format: 'HTTP error: {message}',
      attributes: {
        message:
          '`` is not a valid ID. Must be at least 2 and at most 36 characters long and may consist of only letters, numbers and dashes. It may not start or end with a dash',
      },
    },
  ],
}

export default {
  title: 'Utility Components/ErrorMessage',

  decorators: [
    withInfo({
      inline: true,
      header: false,
      propTables: [ErrorMessage],
    }),
  ],
}

export const Default = () => <ErrorMessage content={exampleError} />

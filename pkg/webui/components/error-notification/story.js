

import React from 'react'

import ErrorNotification from '.'

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

const testErrorWithMarkup = {
  code: 2,
  message: 'error:pkg/assets:http (HTTP error: `` is not a valid ID.',
  details: [
    {
      '@type': 'type.googleapis.com/ttn.lorawan.v3.ErrorDetails',
      namespace: 'pkg/assets',
      name: 'http',
      message_format: 'HTTP error: {message}',
      attributes: {
        message: '`12345` is not a valid ID.',
      },
    },
  ],
}

export default {
  title: 'Notification/ErrorNotification',
}

export const Error = () => (
  <div>
    <ErrorNotification title="example message title" content="We've got a problem here!" />
    <ErrorNotification content="We've got a problem here" />
    <ErrorNotification
      title="We got a problem here! And the description is quite lengthy as well, which can sometimes be a problem."
      content={exampleError}
      small
    />
    <ErrorNotification content={exampleError} small />
    <ErrorNotification title="example of error with markup" content={testErrorWithMarkup} />
  </div>
)

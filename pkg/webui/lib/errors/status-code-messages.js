

/* eslint-disable quote-props */

import { defineMessages } from 'react-intl'

import sharedMessages from '../shared-messages'

// eslint-disable-next-line capitalized-comments
// prettier-ignore
export default {
  '520': sharedMessages.unknownError,
  ...defineMessages({
    '4××': 'Client error',
    '400': 'Bad request',
    '401': 'Unauthorized',
    '403': 'Forbidden',
    '404': 'Not found',
    '409': 'Conflict',
    '429': 'Too many requests',
    '499': 'Client closed request',
    '500': 'Internal server error',
    '501': 'Not implemented',
    '503': 'Service unavailable',
    '504': 'Gateway timeout',
  })
}

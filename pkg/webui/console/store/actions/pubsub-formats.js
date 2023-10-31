

import createRequestActions from '@ttn-lw/lib/store/actions/create-request-actions'

export const GET_PUBSUB_FORMATS_BASE = 'GET_PUBSUB_FORMATS'
export const [
  {
    request: GET_PUBSUB_FORMATS,
    success: GET_PUBSUB_FORMATS_SUCCESS,
    failure: GET_PUBSUB_FORMATS_FAILURE,
  },
  { request: getPubsubFormats, success: getPubsubFormatsSuccess, failure: getPubsubFormatsFailure },
] = createRequestActions(GET_PUBSUB_FORMATS_BASE)

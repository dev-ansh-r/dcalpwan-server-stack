

import createRequestActions from '@ttn-lw/lib/store/actions/create-request-actions'

export const GET_WEBHOOK_FORMATS_BASE = 'GET_WEBHOOK_FORMATS'
export const [
  {
    request: GET_WEBHOOK_FORMATS,
    success: GET_WEBHOOK_FORMATS_SUCCESS,
    failure: GET_WEBHOOK_FORMATS_FAILURE,
  },
  {
    request: getWebhookFormats,
    success: getWebhookFormatsSuccess,
    failure: getWebhookFormatsFailure,
  },
] = createRequestActions(GET_WEBHOOK_FORMATS_BASE)

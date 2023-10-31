

import { createFetchingSelector } from '@ttn-lw/lib/store/selectors/fetching'
import { createErrorSelector } from '@ttn-lw/lib/store/selectors/error'

import { GET_WEBHOOK_FORMATS_BASE } from '@console/store/actions/webhook-formats'

const selectWebhookFormatsStore = state => state.webhookFormats

export const selectWebhookFormats = state => {
  const store = selectWebhookFormatsStore(state)

  return store.formats || {}
}

export const selectWebhookFormatsError = createErrorSelector(GET_WEBHOOK_FORMATS_BASE)
export const selectWebhookFormatsFetching = createFetchingSelector(GET_WEBHOOK_FORMATS_BASE)

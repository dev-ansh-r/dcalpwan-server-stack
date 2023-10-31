

import { createFetchingSelector } from '@ttn-lw/lib/store/selectors/fetching'
import { createErrorSelector } from '@ttn-lw/lib/store/selectors/error'

import { GET_PUBSUB_FORMATS_BASE } from '@console/store/actions/pubsub-formats'

const selectPubsubFormatsStore = state => state.pubsubFormats

export const selectPubsubFormats = state => {
  const store = selectPubsubFormatsStore(state)

  return store.formats || {}
}

export const selectPubsubFormatsError = createErrorSelector(GET_PUBSUB_FORMATS_BASE)
export const selectPubsubFormatsFetching = createFetchingSelector(GET_PUBSUB_FORMATS_BASE)

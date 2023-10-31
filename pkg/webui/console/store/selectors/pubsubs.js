

import { createFetchingSelector } from '@ttn-lw/lib/store/selectors/fetching'
import { createErrorSelector } from '@ttn-lw/lib/store/selectors/error'

import { GET_PUBSUB_BASE, GET_PUBSUBS_LIST_BASE } from '@console/store/actions/pubsubs'

const selectPubsubStore = state => state.pubsubs

// Pubsub.
export const selectPubsubEntityStore = state => selectPubsubStore(state).entities
export const selectPubsubById = (state, id) => selectPubsubEntityStore(state)[id]
export const selectSelectedPubsubId = state => selectPubsubStore(state).selectedPubsub
export const selectSelectedPubsub = state =>
  selectPubsubEntityStore(state)[selectSelectedPubsubId(state)]
export const selectPubsubError = createErrorSelector(GET_PUBSUB_BASE)
export const selectPubsubFetching = createFetchingSelector(GET_PUBSUB_BASE)

// Pubsubs.
export const selectPubsubs = state => Object.values(selectPubsubEntityStore(state))
export const selectPubsubsTotalCount = state => selectPubsubStore(state).totalCount
export const selectPubsubsFetching = createFetchingSelector(GET_PUBSUBS_LIST_BASE)
export const selectPubsubsError = createErrorSelector(GET_PUBSUBS_LIST_BASE)

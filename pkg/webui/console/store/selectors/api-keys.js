

import {
  createPaginationIdsSelectorByEntity,
  createPaginationTotalCountSelectorByEntity,
} from '@ttn-lw/lib/store/selectors/pagination'

const ENTITY = 'apiKeys'

// Api key.
export const selectApiKeysStore = state => state.apiKeys || {}
export const selectApiKeysEntitiesStore = state => selectApiKeysStore(state).entities
export const selectApiKeyById = (state, id) => selectApiKeysEntitiesStore(state)[id]
export const selectSelectedApiKeyId = state => selectApiKeysStore(state).selectedApiKey
export const selectSelectedApiKey = state => selectApiKeyById(state, selectSelectedApiKeyId(state))

// Api keys.
const createSelectApiKeysIdsSelector = createPaginationIdsSelectorByEntity(ENTITY)
const createSelectApiKeysTotalCountSelector = createPaginationTotalCountSelectorByEntity(ENTITY)

export const selectApiKeys = state =>
  createSelectApiKeysIdsSelector(state).map(id => selectApiKeyById(state, id))
export const selectApiKeysTotalCount = state => createSelectApiKeysTotalCountSelector(state)

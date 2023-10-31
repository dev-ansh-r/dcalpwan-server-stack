

import createRequestActions from '@ttn-lw/lib/store/actions/create-request-actions'
import { createPaginationByParentRequestActions } from '@ttn-lw/lib/store/actions/pagination'

export const SHARED_NAME = 'API_KEYS'

export const GET_API_KEY_BASE = 'GET_API_KEY'
export const [
  { request: GET_API_KEY, success: GET_API_KEY_SUCCESS, failure: GET_API_KEY_FAILURE },
  { request: getApiKey, success: getApiKeySuccess, failure: getApiKeyFailure },
] = createRequestActions(
  GET_API_KEY_BASE,
  (parentType, parentId, keyId) => ({ parentType, parentId, keyId }),
  (parentType, parentId, keyId, selector) => ({ selector }),
)

export const GET_API_KEYS_LIST_BASE = 'GET_API_KEYS_LIST'
export const [
  {
    request: GET_API_KEYS_LIST,
    success: GET_API_KEYS_LIST_SUCCESS,
    failure: GET_API_KEYS_LIST_FAILURE,
  },
  { request: getApiKeysList, success: getApiKeysListSuccess, failure: getApiKeysListFailure },
] = createPaginationByParentRequestActions(SHARED_NAME)

export const CREATE_APP_API_KEY_BASE = 'CREATE_APP_API_KEY'
export const [
  {
    request: CREATE_APP_API_KEY,
    success: CREATE_APP_API_KEY_SUCCESS,
    failure: CREATE_APP_API_KEY_FAILURE,
  },
  {
    request: createApplicationApiKey,
    success: createApplicationApiKeySuccess,
    failure: createApplicationApiKeyFailure,
  },
] = createRequestActions(CREATE_APP_API_KEY_BASE, (id, key) => ({ id, key }))

export const CREATE_GATEWAY_API_KEY_BASE = 'CREATE_GATEWAY_API_KEY'
export const [
  {
    request: CREATE_GATEWAY_API_KEY,
    success: CREATE_GATEWAY_API_KEY_SUCCESS,
    failure: CREATE_GATEWAY_API_KEY_FAILURE,
  },
  {
    request: createGatewayApiKey,
    success: createGatewayApiKeySuccess,
    failure: createGatewayApiKeyFailure,
  },
] = createRequestActions(CREATE_GATEWAY_API_KEY_BASE, (gatewayId, key) => ({ gatewayId, key }))



import { getApiKeyId } from '@ttn-lw/lib/selectors/id'

import {
  GET_API_KEYS_LIST_SUCCESS,
  GET_API_KEY_SUCCESS,
  GET_API_KEY,
} from '@console/store/actions/api-keys'

const defaultState = {
  entities: {},
  selectedApiKey: null,
}

const apiKey = (state = {}, apiKey) => ({
  ...state,
  ...apiKey,
})

const apiKeys = (state = defaultState, { type, payload }) => {
  switch (type) {
    case GET_API_KEY:
      return {
        ...state,
        selectedApiKey: payload.keyId,
      }
    case GET_API_KEY_SUCCESS:
      const id = getApiKeyId(payload)
      return {
        ...state,
        entities: {
          ...state.entities,
          [id]: apiKey(state.entities[id], payload),
        },
      }
    case GET_API_KEYS_LIST_SUCCESS:
      return {
        ...state,
        entities: {
          ...payload.entities.reduce(
            (acc, ak) => {
              const id = getApiKeyId(ak)
              acc[id] = apiKey(state.entities[id], ak)
              return acc
            },
            { ...state.entities },
          ),
        },
      }
    default:
      return state
  }
}

export default apiKeys

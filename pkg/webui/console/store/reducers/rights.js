

import createRequestActions from '@ttn-lw/lib/store/actions/create-request-actions'

import { createGetRightsListActionType } from '@console/store/actions/rights'

const defaultState = {
  rights: {
    regular: [],
    pseudo: [],
  },
}

const createNamedRightsReducer = (reducerName = '') => {
  const GET_LIST_BASE = createGetRightsListActionType(reducerName)
  const [{ success: GET_LIST_SUCCESS }] = createRequestActions(GET_LIST_BASE)

  return (state = defaultState, { type, payload }) => {
    switch (type) {
      case GET_LIST_SUCCESS:
        return {
          ...state,
          rights: {
            regular: payload.filter(right => !right.endsWith('_ALL')),
            pseudo: payload.filter(right => right.endsWith('_ALL')),
          },
        }
      default:
        return state
    }
  }
}

export default createNamedRightsReducer

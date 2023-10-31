

import { INITIALIZE_SUCCESS } from '@ttn-lw/lib/store/actions/init'

import {
  GET_USER_SUCCESS,
  UPDATE_USER_SUCCESS,
  GET_USER_RIGHTS_SUCCESS,
} from '@account/store/actions/user'

const defaultState = {
  user: undefined,
  sessionId: undefined,
  rights: {
    regular: [],
    pseudo: [],
  },
}

const user = (state = defaultState, { type, payload }) => {
  switch (type) {
    case INITIALIZE_SUCCESS:
      if (typeof payload !== 'string') {
        return state
      }

      return {
        ...state,
        sessionId: payload,
      }
    case GET_USER_SUCCESS:
      return {
        ...state,
        user: payload,
      }
    case UPDATE_USER_SUCCESS:
      return {
        ...state,
        user: {
          ...state.user,
          ...payload,
        },
      }
    case GET_USER_RIGHTS_SUCCESS:
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

export default user

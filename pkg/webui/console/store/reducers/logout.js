

import {
  GET_USER_ME_SUCCESS,
  GET_USER_ME_FAILURE,
  GET_USER_ME,
  LOGOUT_FAILURE,
  GET_USER_RIGHTS_SUCCESS,
} from '@console/store/actions/logout'

const defaultState = {
  user: undefined,
  error: undefined,
  rights: undefined,
}

const user = (state = defaultState, { type, payload }) => {
  switch (type) {
    case GET_USER_ME:
      return {
        ...state,
        user: undefined,
        error: undefined,
      }

    case GET_USER_ME_SUCCESS:
      return {
        ...state,
        user: payload,
        error: undefined,
      }
    case GET_USER_ME_FAILURE:
      return {
        ...state,
        user: undefined,
        error: payload,
      }
    case GET_USER_RIGHTS_SUCCESS:
      return {
        ...state,
        rights: payload,
      }
    case LOGOUT_FAILURE:
      return {
        ...state,
        error: payload,
      }
    default:
      return state
  }
}

export default user

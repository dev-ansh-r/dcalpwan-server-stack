

import { GET_USER_SESSIONS_LIST_SUCCESS } from '@account/store/actions/sessions'

const defaultState = {
  sessions: undefined,
  totalCount: undefined,
}

const session = (state = defaultState, { type, payload }) => {
  switch (type) {
    case GET_USER_SESSIONS_LIST_SUCCESS:
      return {
        ...state,
        sessions: payload.sessions,
        totalCount: payload.sessionsTotalCount,
      }
    default:
      return state
  }
}

export default session

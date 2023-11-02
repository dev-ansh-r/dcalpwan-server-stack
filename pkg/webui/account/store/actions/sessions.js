

import createRequestActions from '@ttn-lw/lib/store/actions/create-request-actions'

export const GET_USER_SESSIONS_LIST_BASE = 'GET_USER_SESSIONS_LIST'
export const [
  {
    request: GET_USER_SESSIONS_LIST,
    success: GET_USER_SESSIONS_LIST_SUCCESS,
    failure: GET_USER_SESSIONS_LIST_FAILURE,
  },
  {
    request: getUserSessionsList,
    success: getUserSessionsListSuccess,
    failure: getUserSessionsListFailure,
  },
] = createRequestActions(
  GET_USER_SESSIONS_LIST_BASE,
  (id, params) => ({ id, params }),
  (id, params, selector) => ({ selector }),
)

export const DELETE_USER_SESSION_BASE = 'DELETE_USER_SESSION'
export const [
  {
    request: DELETE_USER_SESSION,
    success: DELETE_USER_SESSION_SUCCESS,
    failure: DELETE_USER_SESSION_FAILURE,
  },
  {
    request: deleteUserSession,
    success: deleteUserSessionSuccess,
    failure: deleteUserSessionFailure,
  },
] = createRequestActions(DELETE_USER_SESSION_BASE, (user, sessionId) => ({ user, sessionId }))

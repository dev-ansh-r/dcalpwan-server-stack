

import createRequestActions from '@ttn-lw/lib/store/actions/create-request-actions'

export const GET_USER_ME_BASE = 'GET_USER_ME'
export const [
  { request: GET_USER_ME, success: GET_USER_ME_SUCCESS, failure: GET_USER_ME_FAILURE },
  { request: getUserMe, success: getUserMeSuccess, failure: getUserMeFailure },
] = createRequestActions(GET_USER_ME_BASE)

export const GET_USER_RIGHTS_BASE = 'GET_USER_RIGHTS'
export const [
  { request: GET_USER_RIGHTS, success: GET_USER_RIGHTS_SUCCESS, failure: GET_USER_RIGHTS_FAILURE },
  { request: getUserRights, success: getUserRightsSuccess, failure: getUserRightsFailure },
] = createRequestActions(GET_USER_RIGHTS_BASE)

export const LOGOUT_BASE = 'LOGOUT'
export const [
  { request: LOGOUT, success: LOGOUT_SUCCESS, failure: LOGOUT_FAILURE },
  { request: logout, success: logoutSuccess, failure: logoutFailure },
] = createRequestActions(LOGOUT_BASE)

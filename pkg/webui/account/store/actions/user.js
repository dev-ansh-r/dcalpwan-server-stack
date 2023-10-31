

import createRequestActions from '@ttn-lw/lib/store/actions/create-request-actions'

export const GET_USER_BASE = 'GET_USER'
export const [
  { request: GET_USER, success: GET_USER_SUCCESS, failure: GET_USER_FAILURE },
  { request: getUser, success: getUserSuccess, failure: getUserFailure },
] = createRequestActions(
  GET_USER_BASE,
  id => ({ id }),
  (id, selector) => ({ selector }),
)

export const UPDATE_USER_BASE = 'UPDATE_USER'
export const [
  { request: UPDATE_USER, success: UPDATE_USER_SUCCESS, failure: UPDATE_USER_FAILURE },
  { request: updateUser, success: updateUserSuccess, failure: updateUserFailure },
] = createRequestActions(UPDATE_USER_BASE, ({ id, patch }) => ({ id, patch }))

export const DELETE_USER_BASE = 'DELETE_USER'
export const [
  { request: DELETE_USER, success: DELETE_USER_SUCCESS, failure: DELETE_USER_FAILURE },
  { request: deleteUser, success: deleteUserSuccess, failure: deleteUserFailure },
] = createRequestActions(
  DELETE_USER_BASE,
  id => ({ id }),
  (id, options = {}) => ({ options }),
)

export const LOGOUT_BASE = 'LOGOUT'
export const [
  { request: LOGOUT, success: LOGOUT_SUCCESS, failure: LOGOUT_FAILURE },
  { request: logout, success: logoutSuccess, failure: logoutFailure },
] = createRequestActions(LOGOUT_BASE)

export const GET_USER_RIGHTS_BASE = 'GET_USER_RIGHTS'
export const [
  { request: GET_USER_RIGHTS, success: GET_USER_RIGHTS_SUCCESS, failure: GET_USER_RIGHTS_FAILURE },
  { request: getUserRights, success: getUserRightsSuccess, failure: getUserRightsFailure },
] = createRequestActions(GET_USER_RIGHTS_BASE, userId => ({ userId }))



import { createAction } from 'redux-actions'

import ONLINE_STATUS from '@ttn-lw/constants/online-status'

import createRequestActions from '@ttn-lw/lib/store/actions/create-request-actions'

export const SET_CONNECTION_STATUS = 'SET_CONNECTION_STATUS'

export const setStatusOnline = createAction(SET_CONNECTION_STATUS, (isOnline = true) => ({
  onlineStatus: isOnline ? ONLINE_STATUS.ONLINE : ONLINE_STATUS.OFFLINE,
}))
export const setStatusChecking = createAction(SET_CONNECTION_STATUS, () => ({
  onlineStatus: ONLINE_STATUS.CHECKING,
}))

export const SET_LOGIN_STATUS = 'SET_LOGIN_STATUS'

export const setLoginStatus = createAction(SET_LOGIN_STATUS, () => ({
  isLoginRequired: status.isLoginRequired,
}))

export const ATTEMPT_RECONNECT = 'ATTEMPT_RECONNECT'
export const attemptReconnect = createAction(ATTEMPT_RECONNECT)
export const [
  { success: ATTEMPT_RECONNECT_SUCCESS, failure: ATTEMPT_RECONNECT_FAILURE },
  { success: attemptReconnectSuccess, failure: attemptReconnectFailure },
] = createRequestActions(ATTEMPT_RECONNECT)

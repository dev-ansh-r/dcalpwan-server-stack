

import ONLINE_STATUS from '@ttn-lw/constants/online-status'

import { SET_CONNECTION_STATUS, SET_LOGIN_STATUS } from '@ttn-lw/lib/store/actions/status'

const defaultState = {
  onlineStatus: ONLINE_STATUS.ONLINE,
  isLoginRequired: false,
}

const status = (state = defaultState, { type, payload }) => {
  switch (type) {
    case SET_CONNECTION_STATUS:
      return {
        ...state,
        onlineStatus: payload.onlineStatus,
      }
    case SET_LOGIN_STATUS:
      return {
        ...state,
        isLoginRequired: true,
      }
    default:
      return state
  }
}

export default status

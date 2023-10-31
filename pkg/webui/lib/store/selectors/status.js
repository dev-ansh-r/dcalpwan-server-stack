

import ONLINE_STATUS from '@ttn-lw/constants/online-status'

export const selectStatusStore = state => state.status

export const selectOnlineStatus = state => selectStatusStore(state).onlineStatus

export const selectIsOnlineStatus = state =>
  selectStatusStore(state).onlineStatus === ONLINE_STATUS.ONLINE

export const selectIsOfflineStatus = state =>
  selectStatusStore(state).onlineStatus === ONLINE_STATUS.OFFLINE

export const selectIsCheckingStatus = state =>
  selectStatusStore(state).onlineStatus === ONLINE_STATUS.CHECKING

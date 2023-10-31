

import { createFetchingSelector } from '@ttn-lw/lib/store/selectors/fetching'
import { createErrorSelector } from '@ttn-lw/lib/store/selectors/error'

import { GET_USER_ME_BASE, LOGOUT_BASE } from '@console/store/actions/logout'

const selectUserStore = state => state.user

export const selectUser = state => selectUserStore(state).user
export const selectUserError = createErrorSelector([GET_USER_ME_BASE, LOGOUT_BASE])
export const selectUserFetching = createFetchingSelector(GET_USER_ME_BASE)

export const selectUserId = state => {
  const user = selectUser(state)

  if (!Boolean(user)) {
    return undefined
  }

  return user.ids.user_id
}

export const selectUserName = state => {
  const user = selectUser(state)

  return user ? user.name : undefined
}

export const selectUserNameOrId = state => {
  const user = selectUser(state)

  return user ? user.name || user.ids.user_id : undefined
}

export const selectUserIsAdmin = state => {
  const user = selectUser(state)

  return user ? user.isAdmin : false
}

export const selectUserRights = state => selectUserStore(state).rights

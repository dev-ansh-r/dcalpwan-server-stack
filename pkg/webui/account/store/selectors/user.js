

import { createSelector } from 'reselect'

import { createFetchingSelector } from '@ttn-lw/lib/store/selectors/fetching'
import { createErrorSelector } from '@ttn-lw/lib/store/selectors/error'

import { GET_USER_RIGHTS_BASE } from '@account/store/actions/user'

const selectUserStore = state => state.user

export const selectUser = state => selectUserStore(state).user

export const selectSessionId = state => selectUserStore(state).sessionId

export const selectUserId = state => {
  const user = selectUser(state)

  if (!Boolean(user)) {
    return undefined
  }

  return user.ids.user_id
}

export const selectUserIsAdmin = state => {
  const user = selectUser(state)
  return user.admin || false
}

export const selectUserName = state => selectUser(state).name

export const selectUserProfilePicture = state => selectUser(state).profile_picture

// Rights.
export const selectUserRights = createSelector(selectUserStore, ({ rights }) => [
  ...rights.regular,
  ...rights.pseudo,
])
export const selectUserRegularRights = state => selectUserStore(state).rights?.regular
export const selectUserPseudoRights = state => selectUserStore(state).rights?.pseudo
export const selectUserRightsError = createErrorSelector(GET_USER_RIGHTS_BASE)
export const selectUserRightsFetching = createFetchingSelector(GET_USER_RIGHTS_BASE)

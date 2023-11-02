

import { createFetchingSelector } from '@ttn-lw/lib/store/selectors/fetching'
import { createErrorSelector } from '@ttn-lw/lib/store/selectors/error'

import { GET_USER_SESSIONS_LIST_BASE } from '@account/store/actions/sessions'

const selectSessionsStore = state => state.session

export const selectUserSessions = state => selectSessionsStore(state).sessions
export const selectUserSessionsTotalCount = state => selectSessionsStore(state).totalCount
export const selectUserSessionsFetching = createFetchingSelector(GET_USER_SESSIONS_LIST_BASE)
export const selectUserSessionsError = createErrorSelector(GET_USER_SESSIONS_LIST_BASE)



import { createFetchingSelector } from '@ttn-lw/lib/store/selectors/fetching'
import { createErrorSelector } from '@ttn-lw/lib/store/selectors/error'
import {
  createPaginationIdsSelectorByEntity,
  createPaginationTotalCountSelectorByEntity,
} from '@ttn-lw/lib/store/selectors/pagination'

import { GET_USER_BASE, GET_USER_RIGHTS_LIST_BASE } from '@console/store/actions/users'

import { createRightsSelector, createPseudoRightsSelector } from './rights'

const ENTITY = 'users'

const selectUserStore = state => state.users

// User.
export const selectUserEntitiesStore = state => selectUserStore(state)?.entities || {}
export const selectUserById = (state, id) => selectUserEntitiesStore(state)[id]
export const selectSelectedUserId = state => selectUserStore(state).selectedUser
export const selectSelectedUser = state => selectUserById(state, selectSelectedUserId(state))
export const selectUserFetching = createFetchingSelector(GET_USER_BASE)
export const selectUserError = createErrorSelector(GET_USER_BASE)

// Users.
const selectUsrsIds = createPaginationIdsSelectorByEntity(ENTITY)
const selectUsrsTotalCount = createPaginationTotalCountSelectorByEntity(ENTITY)

export const selectUsers = state => selectUsrsIds(state).map(id => selectUserById(state, id))
export const selectUsersTotalCount = state => selectUsrsTotalCount(state)

// Rights.
export const selectUserRights = createRightsSelector(ENTITY)
export const selectUserPseudoRights = createPseudoRightsSelector(ENTITY)
export const selectUserRightsError = createErrorSelector(GET_USER_RIGHTS_LIST_BASE)
export const selectUserRightsFetching = createFetchingSelector(GET_USER_RIGHTS_LIST_BASE)

// Invitations.
export const selectUserInvitations = state => selectUserStore(state).invitations
export const selectUserInvitationsTotalCount = state => selectUserStore(state).invitationsTotalCount

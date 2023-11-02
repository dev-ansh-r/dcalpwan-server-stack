

import { createFetchingSelector } from '@ttn-lw/lib/store/selectors/fetching'
import { createErrorSelector } from '@ttn-lw/lib/store/selectors/error'

import {
  GET_AUTHORIZATIONS_LIST_BASE,
  GET_ACCESS_TOKENS_LIST_BASE,
} from '@account/store/actions/authorizations'

const selectAuthorizationsStore = state => state.authorizations

export const selectAuthorizations = state => selectAuthorizationsStore(state).authorizations
export const selectSelectedAuthorization = (state, clientId) => {
  const authorizations = selectAuthorizationsStore(state).authorizations.reduce((acc, cur) => {
    const id = cur.client_ids.client_id
    acc[id] = cur

    return acc
  }, {})

  return authorizations[clientId]
}
export const selectAuthorizationsTotalCount = state =>
  selectAuthorizationsStore(state).authorizationsTotalCount
export const selectAuthorizationsFetching = createFetchingSelector(GET_AUTHORIZATIONS_LIST_BASE)
export const selectAuthorizationsError = createErrorSelector(GET_AUTHORIZATIONS_LIST_BASE)

export const selectTokens = state => selectAuthorizationsStore(state).tokens
export const selectTokenIds = state => selectTokens(state).map(token => token.id)
export const selectTokensTotalCount = state => selectAuthorizationsStore(state).tokensTotalCount
export const selectTokensFetching = createFetchingSelector(GET_ACCESS_TOKENS_LIST_BASE)
export const selectTokensError = createErrorSelector(GET_ACCESS_TOKENS_LIST_BASE)

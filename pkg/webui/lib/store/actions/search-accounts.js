

import createRequestActions from '@ttn-lw/lib/store/actions/create-request-actions'

export const SEARCH_ACCOUNTS_BASE = 'SEARCH_ACCOUNTS'
export const [
  { request: SEARCH_ACCOUNTS, success: SEARCH_ACCOUNTS_SUCCESS, failure: SEARCH_ACCOUNTS_FAILURE },
  { request: searchAccounts, success: searchAccountsSuccess, failure: searchAccountsFailure },
] = createRequestActions(SEARCH_ACCOUNTS_BASE, (query, onlyUsers, collaboratorOf) => ({
  query,
  onlyUsers,
  collaboratorOf,
}))

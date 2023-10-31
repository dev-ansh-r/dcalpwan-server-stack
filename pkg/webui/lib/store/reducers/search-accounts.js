

import { SEARCH_ACCOUNTS_SUCCESS } from '@ttn-lw/lib/store/actions/search-accounts'

const defaultState = {
  searchResults: {
    account_ids: [],
  },
  totalCount: 0,
}

const searchAccounts = (state = defaultState, { type, payload }) => {
  switch (type) {
    case SEARCH_ACCOUNTS_SUCCESS:
      return {
        ...state,
        searchResults: payload,
      }
    default:
      return state
  }
}

export default searchAccounts

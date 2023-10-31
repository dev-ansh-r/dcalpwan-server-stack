

export const selectSearchResultsStore = state => state.searchAccounts
export const selectSearchResults = state => selectSearchResultsStore(state).searchResults
export const selectSearchResultAccountIds = state => selectSearchResults(state).account_ids



/* eslint-disable import/prefer-default-export */

const selectFetchingStore = state => state.ui.fetching

const selectFetchingEntry = (state, id) => selectFetchingStore(state)[id] || false

/**
 * @example
 * const selectFetching = createFetchingSelector([
 * GET_ENTITY_LIST_BASE,
 * SEARCH_ENTITY_LIST_BASE
 * ])
 * const selectEntityFetching = (state) => selectFetching(state)
 *
 * Creates the fetching selector for a set of base action types.
 * @param {Array} actions - A list of base action types or a single base action type.
 * @returns {boolean} `true` if one of the base action types is in the `fetching` state,
 * `false` otherwise.
 */
export const createFetchingSelector = actions => state => {
  if (!Array.isArray(actions)) {
    return selectFetchingEntry(state, actions)
  }

  return actions.some(action => selectFetchingEntry(state, action))
}

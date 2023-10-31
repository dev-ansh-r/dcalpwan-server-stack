

/* eslint-disable import/prefer-default-export */

const selectErrorStore = state => state.ui.error

const getErrorStoreEntrySelector = (state, baseActionType) =>
  selectErrorStore(state)[baseActionType]

/**
 * @example
 * const selectError = createErrorSelector([
 * 'GET_ENTITY_LIST',
 * 'SEARCH_ENTITY_LIST'
 * ])
 * const selectEntityError = (state) => selectError(state)
 *
 * Creates the error selector for a set of base action types.
 * @param {Array} actions - A list of base action types or a single base action type.
 * @returns {object} The error object matching one of the base action types.
 */
export const createErrorSelector = actions => state => {
  if (!Array.isArray(actions)) {
    return getErrorStoreEntrySelector(state, actions)
  }

  for (const action of actions) {
    const error = getErrorStoreEntrySelector(state, action)

    if (Boolean(error)) {
      return error
    }
  }
}

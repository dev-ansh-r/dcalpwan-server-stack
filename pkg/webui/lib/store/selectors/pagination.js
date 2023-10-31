

const selectStore = state => state.pagination

const createStoreSelectorByEntity = entity => state => selectStore(state)[entity] || {}

export const createPaginationIdsSelectorByEntity = entity => {
  const storeSelector = createStoreSelectorByEntity(entity)

  return state => storeSelector(state).ids || []
}

export const createPaginationIdsSelectorByEntityAndId = entity => {
  const storeSelector = createStoreSelectorByEntity(entity)

  return (state, id) => {
    const store = storeSelector(state)
    const storeById = store[id] || {}

    return storeById.ids || []
  }
}

export const createPaginationTotalCountSelectorByEntity = entity => {
  const storeSelector = createStoreSelectorByEntity(entity)

  return state => storeSelector(state).totalCount
}

export const createPaginationTotalCountSelectorByEntityAndId = entity => {
  const storeSelector = createStoreSelectorByEntity(entity)

  return (state, id) => {
    const store = storeSelector(state)
    const storeById = store[id] || {}

    return storeById.totalCount
  }
}

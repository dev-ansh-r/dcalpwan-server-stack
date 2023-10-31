

const selectRightsStore = (state, entity) => state.rights[entity]

export const createRightsSelector = entity => state => {
  const store = selectRightsStore(state, entity)

  return store.rights.regular
}

export const createPseudoRightsSelector = entity => state => {
  const store = selectRightsStore(state, entity)

  return store.rights.pseudo
}

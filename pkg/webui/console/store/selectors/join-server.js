

const selectJsStore = state => state.js

const selectJoinEUIPrefixes = state => {
  const store = selectJsStore(state)

  return store.prefixes
}

export default selectJoinEUIPrefixes



const fetching = (state = {}, action) => {
  const { type } = action
  const matches = /(.*)_(REQUEST|SUCCESS|FAILURE|ABORT)/.exec(type)

  if (!matches) {
    return state
  }

  const [, key, status] = matches
  return {
    ...state,
    [key]: status === 'REQUEST',
  }
}

export default fetching

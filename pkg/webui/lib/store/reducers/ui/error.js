

const error = (state = {}, action) => {
  const { type, payload: errorValue } = action

  const matches = /(.*)_(REQUEST|SUCCESS|FAILURE)/.exec(type)
  if (!matches) {
    return state
  }

  const [, key, status] = matches
  return {
    ...state,
    [key]: status === 'FAILURE' ? errorValue : undefined,
  }
}

export default error

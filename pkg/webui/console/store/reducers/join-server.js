

import { handleActions } from 'redux-actions'

import { GET_JOIN_EUI_PREFIXES_SUCCESS } from '@console/store/actions/join-server'

const defaultState = {
  prefixes: [],
}

export default handleActions(
  {
    [GET_JOIN_EUI_PREFIXES_SUCCESS]: (state, { payload }) => ({
      ...state,
      prefixes: payload,
    }),
  },
  defaultState,
)



import { handleActions } from 'redux-actions'

import { GET_IS_CONFIGURATION_SUCCESS } from '@console/store/actions/identity-server'

const defaultState = {
  configuration: {},
}

export default handleActions(
  {
    [GET_IS_CONFIGURATION_SUCCESS]: (state, { payload }) => ({
      ...state,
      configuration: payload.configuration,
    }),
  },
  defaultState,
)

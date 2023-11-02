

import { handleActions } from 'redux-actions'

import { GET_AS_CONFIGURATION_SUCCESS } from '@console/store/actions/application-server'

const defaultState = {
  configuration: {},
}

export default handleActions(
  {
    [GET_AS_CONFIGURATION_SUCCESS]: (state, { payload }) => ({
      ...state,
      configuration: payload.configuration,
    }),
  },
  defaultState,
)

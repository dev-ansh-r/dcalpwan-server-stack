

import {
  GET_NS_FREQUENCY_PLANS_SUCCESS,
  GET_GS_FREQUENCY_PLANS_SUCCESS,
} from '@console/store/actions/configuration'

const defaultState = {
  nsFrequencyPlans: undefined,
  gsFrequencyPlans: undefined,
}

const configuration = (state = defaultState, { type, payload }) => {
  switch (type) {
    case GET_NS_FREQUENCY_PLANS_SUCCESS:
      return {
        ...state,
        nsFrequencyPlans: payload,
      }
    case GET_GS_FREQUENCY_PLANS_SUCCESS:
      return {
        ...state,
        gsFrequencyPlans: payload,
      }
    default:
      return state
  }
}

export default configuration

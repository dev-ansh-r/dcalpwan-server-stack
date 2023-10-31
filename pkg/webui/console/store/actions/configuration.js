

import createRequestActions from '@ttn-lw/lib/store/actions/create-request-actions'

export const GET_NS_FREQUENCY_PLANS_BASE = 'GET_NS_FREQUENCY_PLANS'
export const [
  {
    request: GET_NS_FREQUENCY_PLANS,
    success: GET_NS_FREQUENCY_PLANS_SUCCESS,
    failure: GET_NS_FREQUENCY_PLANS_FAILURE,
  },
  {
    request: getNsFrequencyPlans,
    success: getNsFrequencyPlansSuccess,
    failure: getNsFrequencyPlansFailure,
  },
] = createRequestActions(GET_NS_FREQUENCY_PLANS_BASE)

export const GET_GS_FREQUENCY_PLANS_BASE = 'GET_GS_FREQUENCY_PLANS'
export const [
  {
    request: GET_GS_FREQUENCY_PLANS,
    success: GET_GS_FREQUENCY_PLANS_SUCCESS,
    failure: GET_GS_FREQUENCY_PLANS_FAILURE,
  },
  {
    request: getGsFrequencyPlans,
    success: getGsFrequencyPlansSuccess,
    failure: getGsFrequencyPlansFailure,
  },
] = createRequestActions(GET_GS_FREQUENCY_PLANS_BASE)



import createRequestActions from '@ttn-lw/lib/store/actions/create-request-actions'

export const GET_IS_CONFIGURATION_BASE = 'GET_IS_CONFIGURATION'
export const [
  {
    request: GET_IS_CONFIGURATION,
    success: GET_IS_CONFIGURATION_SUCCESS,
    failure: GET_IS_CONFIGURATION_FAILURE,
  },
  {
    request: getIsConfiguration,
    success: getIsConfigurationSuccess,
    failure: getIsConfigurationFailure,
  },
] = createRequestActions(GET_IS_CONFIGURATION_BASE)

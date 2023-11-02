

import createRequestActions from '@ttn-lw/lib/store/actions/create-request-actions'

export const GET_AS_CONFIGURATION_BASE = 'GET_AS_CONFIGURATION'
export const [
  {
    request: GET_AS_CONFIGURATION,
    success: GET_AS_CONFIGURATION_SUCCESS,
    failure: GET_AS_CONFIGURATION_FAILURE,
  },
  {
    request: getAsConfiguration,
    success: getAsConfigurationSuccess,
    failure: getAsConfigurationFailure,
  },
] = createRequestActions(GET_AS_CONFIGURATION_BASE)

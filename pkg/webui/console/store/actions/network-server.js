

import createRequestActions from '@ttn-lw/lib/store/actions/create-request-actions'

export const GET_DEFAULT_MAC_SETTINGS_BASE = 'GET_DEFAULT_MAC_SETTINGS'
export const [
  {
    request: GET_DEFAULT_MAC_SETTINGS,
    success: GET_DEFAULT_MAC_SETTINGS_SUCCESS,
    failure: GET_DEFAULT_MAC_SETTINGS_FAILURE,
  },
  {
    request: getDefaultMacSettings,
    success: getDefaultMacSettingsSuccess,
    failure: getDefaultMacSettingsFailure,
  },
] = createRequestActions(GET_DEFAULT_MAC_SETTINGS_BASE, (frequencyPlanId, lorawanPhyVersion) => ({
  frequencyPlanId,
  lorawanPhyVersion,
}))

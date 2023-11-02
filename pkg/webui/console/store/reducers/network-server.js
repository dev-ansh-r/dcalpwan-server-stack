

import { handleActions } from 'redux-actions'

import { GET_DEFAULT_MAC_SETTINGS_SUCCESS } from '@console/store/actions/network-server'

import { generateDefaultMacSettingsKey } from '@console/store/selectors/network-server'

const defaultState = {
  defaultMacSettings: {},
}

export default handleActions(
  {
    [GET_DEFAULT_MAC_SETTINGS_SUCCESS]: (
      state,
      { payload: { frequencyPlanId, lorawanPhyVersion, defaultMacSettings } },
    ) => ({
      ...state,
      defaultMacSettings: {
        ...state.defaultMacSettings,
        [generateDefaultMacSettingsKey(frequencyPlanId, lorawanPhyVersion)]: defaultMacSettings,
      },
    }),
  },
  defaultState,
)

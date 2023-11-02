

import tts from '@console/api/tts'

import createRequestLogic from '@ttn-lw/lib/store/logics/create-request-logic'

import * as ns from '@console/store/actions/network-server'

import { selectDefaultMacSettings } from '@console/store/selectors/network-server'

const getDefaultMacSettings = createRequestLogic({
  type: ns.GET_DEFAULT_MAC_SETTINGS,
  process: async ({ getState, action }) => {
    const { frequencyPlanId, lorawanPhyVersion } = action.payload
    const state = getState()
    const cachedResult = selectDefaultMacSettings(state, frequencyPlanId, lorawanPhyVersion)

    // Default MAC settings typically don't change so we can
    // cache the result until the next refresh.
    const defaultMacSettings =
      cachedResult || (await tts.Ns.getDefaultMacSettings(frequencyPlanId, lorawanPhyVersion))

    return {
      defaultMacSettings,
      frequencyPlanId,
      lorawanPhyVersion,
    }
  },
})

export default [getDefaultMacSettings]

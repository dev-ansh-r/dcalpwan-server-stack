

import tts from '@console/api/tts'

import createRequestLogic from '@ttn-lw/lib/store/logics/create-request-logic'

import * as claim from '@console/store/actions/claim'

const claimDeviceLogic = createRequestLogic({
  type: claim.CLAIM_DEVICE,
  process: async ({ action }) => {
    const { appId, qr_code, authenticatedIdentifiers } = action.payload

    return await tts.DeviceClaim.claim(appId, qr_code, authenticatedIdentifiers)
  },
})

const unclaimDeviceLogic = createRequestLogic({
  type: claim.UNCLAIM_DEVICE,
  process: async ({ action }) => {
    const { applicationId, deviceId, devEui, joinEui } = action.payload

    return await tts.DeviceClaim.unclaim(applicationId, deviceId, devEui, joinEui)
  },
})

const getInfoByJoinEUILogic = createRequestLogic({
  type: claim.GET_INFO_BY_JOIN_EUI,
  process: async ({ action }) => {
    const { joinEUI } = action.payload

    return await tts.DeviceClaim.GetInfoByJoinEUI(joinEUI)
  },
})

export default [claimDeviceLogic, unclaimDeviceLogic, getInfoByJoinEUILogic]

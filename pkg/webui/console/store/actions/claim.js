

import createRequestActions from '@ttn-lw/lib/store/actions/create-request-actions'

const CLAIM_DEVICE_BASE = 'CLAIM_DEVICE'
export const [
  { request: CLAIM_DEVICE, success: CLAIM_DEVICE_SUCCESS, failure: CLAIM_DEVICE_FAILURE },
  { request: claimDevice, success: claimDeviceSuccess, failure: claimDeviceFailure },
] = createRequestActions(CLAIM_DEVICE_BASE, (appId, qr_code, authenticatedIdentifiers) => ({
  appId,
  qr_code,
  authenticatedIdentifiers,
}))

const UNCLAIM_DEVICE_BASE = 'UNCLAIM_DEVICE'
export const [
  { request: UNCLAIM_DEVICE, success: UNCLAIM_DEVICE_SUCCESS, failure: UNCLAIM_DEVICE_FAILURE },
  { request: unclaimDevice, success: unclaimDeviceSuccess, failure: unclaimDeviceFailure },
] = createRequestActions(UNCLAIM_DEVICE_BASE, (applicationId, deviceId, devEui, joinEui) => ({
  applicationId,
  deviceId,
  devEui,
  joinEui,
}))

const GET_INFO_BY_JOIN_EUI_BASE = 'GET_INFO_BY_JOIN_EUI'
export const [
  {
    request: GET_INFO_BY_JOIN_EUI,
    success: GET_INFO_BY_JOIN_EUI_SUCCESS,
    failure: GET_INFO_BY_JOIN_EUI_FAILURE,
  },
  { request: getInfoByJoinEUI, success: getInfoByJoinEUISuccess, failure: getInfoByJoinEUIFailure },
] = createRequestActions(GET_INFO_BY_JOIN_EUI_BASE, joinEUI => ({ joinEUI }))



import autoBind from 'auto-bind'

import Bytes from '../util/bytes'
import Marshaler from '../util/marshaler'

class DeviceClaim {
  constructor(registry, { stackConfig }) {
    this._api = registry
    this._stackConfig = stackConfig
    autoBind(this)
  }

  // Claim
  async claim(applicationId, qrCode, values) {
    const deviceToClaim = qrCode ? { qr_code: qrCode } : values
    const payload = {
      ...deviceToClaim,
      target_application_ids: {
        application_id: applicationId,
      },
    }

    const response = await this._api.EndDeviceClaimingServer.Claim(undefined, payload)

    return Marshaler.payloadSingleResponse(response)
  }

  async GetInfoByJoinEUI(join_eui) {
    const response = await this._api.EndDeviceClaimingServer.GetInfoByJoinEUI(undefined, join_eui)

    return Marshaler.payloadSingleResponse(response)
  }

  /**
   * Unclaim an end device.
   *
   * @param {string} applicationId - The Application ID.
   * @param {string} deviceId - The Device ID.
   * @param {Array} devEui - The Device dev_eui.
   * @param {object} joinEui - The Device join_eui.
   * @returns {object} - An empty object on successful requests, an error otherwise.
   */
  async unclaim(applicationId, deviceId, devEui, joinEui) {
    const params = {
      routeParams: {
        'application_ids.application_id': applicationId,
        device_id: deviceId,
      },
    }

    const response = await this._api.EndDeviceClaimingServer.Unclaim(params, {
      dev_eui: Bytes.hexToBase64(devEui),
      join_eui: Bytes.hexToBase64(joinEui),
    })

    return Marshaler.payloadSingleResponse(response)
  }
}

export default DeviceClaim

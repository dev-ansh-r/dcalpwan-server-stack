

import autoBind from 'auto-bind'

import Marshaler from '../util/marshaler'

class As {
  constructor(service) {
    this._api = service
    autoBind(this)
  }

  async getConfiguration() {
    const result = await this._api.As.GetConfiguration()

    return Marshaler.payloadSingleResponse(result)
  }

  async encodeDownlink(appId, deviceId, data) {
    const result = await this._api.AppAs.EncodeDownlink(
      {
        routeParams: {
          'end_device_ids.application_ids.application_id': appId,
          'end_device_ids.device_id': deviceId,
        },
      },
      data,
    )

    return Marshaler.payloadSingleResponse(result)
  }

  async decodeDownlink(appId, deviceId, data) {
    const result = await this._api.AppAs.DecodeDownlink(
      {
        routeParams: {
          'end_device_ids.application_ids.application_id': appId,
          'end_device_ids.device_id': deviceId,
        },
      },
      data,
    )

    return Marshaler.payloadSingleResponse(result)
  }

  async decodeUplink(appId, deviceId, data) {
    const result = await this._api.AppAs.DecodeUplink(
      {
        routeParams: {
          'end_device_ids.application_ids.application_id': appId,
          'end_device_ids.device_id': deviceId,
        },
      },
      data,
    )

    return Marshaler.payloadSingleResponse(result)
  }
}

export default As

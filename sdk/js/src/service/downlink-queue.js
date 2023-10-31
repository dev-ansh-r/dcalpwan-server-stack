

import autoBind from 'auto-bind'

import Marshaler from '../util/marshaler'

class DownlinkQueue {
  constructor(api, { stackConfig }) {
    this._api = api
    this._stackConfig = stackConfig
    autoBind(this)
  }

  async list(applicationId, deviceId) {
    const result = await this._api.DownlinkQueueList({
      routeParams: {
        'application_ids.application_id': applicationId,
        device_id: deviceId,
      },
    })

    return Marshaler.payloadListResponse('downlinks', result)
  }

  async push(applicationId, deviceId, downlinks) {
    const result = await this._api.DownlinkQueuePush(
      {
        routeParams: {
          'end_device_ids.application_ids.application_id': applicationId,
          'end_device_ids.device_id': deviceId,
        },
      },
      {
        downlinks,
      },
    )

    return Marshaler.payloadSingleResponse(result)
  }

  async replace(applicationId, deviceId, downlinks) {
    const result = await this._api.DownlinkQueueReplace(
      {
        routeParams: {
          'end_device_ids.application_ids.application_id': applicationId,
          'end_device_ids.device_id': deviceId,
        },
      },
      {
        downlinks,
      },
    )

    return Marshaler.payloadSingleResponse(result)
  }
}

export default DownlinkQueue

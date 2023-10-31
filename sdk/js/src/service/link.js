

import autoBind from 'auto-bind'

import Marshaler from '../util/marshaler'

class Link {
  constructor(registry) {
    this._api = registry
    autoBind(this)
  }

  async get(appId, fieldMask) {
    const result = await this._api.GetLink(
      {
        routeParams: { 'application_ids.application_id': appId },
      },
      Marshaler.selectorToFieldMask(fieldMask),
    )

    return Marshaler.payloadSingleResponse(result)
  }

  async set(appId, data, mask = Marshaler.fieldMaskFromPatch(data)) {
    const result = await this._api.SetLink(
      {
        routeParams: { 'application_ids.application_id': appId },
      },
      {
        link: data,
        field_mask: Marshaler.fieldMask(mask),
      },
    )

    return Marshaler.payloadSingleResponse(result)
  }

  async delete(appId) {
    const result = await this._api.DeleteLink({
      routeParams: { application_id: appId },
    })

    return Marshaler.payloadSingleResponse(result)
  }

  async getStats(appId) {
    const result = await this._api.GetLinkStats({
      routeParams: { application_id: appId },
    })

    return Marshaler.payloadSingleResponse(result)
  }
}

export default Link

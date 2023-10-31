

import autoBind from 'auto-bind'

import Marshaler from '../util/marshaler'

class ApplicationPackages {
  constructor(registry) {
    this._api = registry
    autoBind(this)
  }

  async getDefaultAssociation(appId, fPort, selector) {
    const result = await this._api.GetDefaultAssociation(
      {
        routeParams: {
          'ids.application_ids.application_id': appId,
          'ids.f_port': fPort,
        },
      },
      Marshaler.selectorToFieldMask(selector),
    )

    return Marshaler.payloadSingleResponse(result)
  }

  async setDefaultAssociation(
    appId,
    fPort,
    patch,
    mask = Marshaler.fieldMaskFromPatch(
      patch,
      this._api.SetDefaultAssociationAllowedFieldMaskPaths,
    ),
  ) {
    const result = await this._api.SetDefaultAssociation(
      {
        routeParams: {
          'default.ids.application_ids.application_id': appId,
          'default.ids.f_port': fPort,
        },
      },
      {
        default: patch,
        field_mask: Marshaler.fieldMask(mask),
      },
    )

    return Marshaler.payloadSingleResponse(result)
  }

  async deleteDefaultAssociation(appId, fPort) {
    const result = await this._api.DeleteDefaultAssociation({
      routeParams: {
        'application_ids.application_id': appId,
        f_port: fPort,
      },
    })

    return Marshaler.payloadSingleResponse(result)
  }
}

export default ApplicationPackages



import autoBind from 'auto-bind'

import Marshaler from '../util/marshaler'

class ApiKeys {
  constructor(registry, { parentRoutes }) {
    this._api = registry
    this._parentRoutes = parentRoutes
    autoBind(this)
  }

  async getById(entityId, id) {
    const entityIdRoute = this._parentRoutes.get
    const result = await this._api.GetAPIKey({
      routeParams: { [entityIdRoute]: entityId, key_id: id },
    })

    return Marshaler.payloadSingleResponse(result)
  }

  async getAll(entityId, params) {
    const entityIdRoute = this._parentRoutes.list
    const result = await this._api.ListAPIKeys(
      {
        routeParams: { [entityIdRoute]: entityId },
      },
      params,
    )

    return Marshaler.payloadListResponse('api_keys', result)
  }

  async create(entityId, key) {
    const entityIdRoute = this._parentRoutes.create
    const result = await this._api.CreateAPIKey(
      {
        routeParams: { [entityIdRoute]: entityId },
      },
      {
        ...key,
      },
    )

    return Marshaler.payloadSingleResponse(result)
  }

  async deleteById(entityId, id) {
    const entityIdRoute = this._parentRoutes.delete
    const result = await this._api.DeleteAPIKey({
      routeParams: { [entityIdRoute]: entityId, key_id: id },
    })

    return Marshaler.payloadSingleResponse(result)
  }

  async updateById(
    entityId,
    id,
    patch,
    mask = Marshaler.fieldMaskFromPatch(patch, this._api.UpdateAllowedFieldMaskPaths),
  ) {
    const entityIdRoute = this._parentRoutes.update
    const result = await this._api.UpdateAPIKey(
      {
        routeParams: {
          [entityIdRoute]: entityId,
          'api_key.id': id,
        },
      },
      {
        api_key: { ...patch },
        field_mask: Marshaler.fieldMask(mask),
      },
    )

    return Marshaler.payloadSingleResponse(result)
  }
}

export default ApiKeys

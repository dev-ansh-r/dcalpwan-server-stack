

import autoBind from 'auto-bind'

import Marshaler from '../util/marshaler'

class Collaborators {
  constructor(registry, { parentRoutes }) {
    this._api = registry
    this._parentRoutes = parentRoutes
    autoBind(this)
  }

  async _getById(entityId, collaboratorId, isUser) {
    const entityIdRoute = this._parentRoutes.get
    const collaboratorIdRoute = isUser
      ? 'collaborator.user_ids.user_id'
      : 'collaborator.organization_ids.organization_id'

    const result = await this._api.GetCollaborator({
      routeParams: {
        [entityIdRoute]: entityId,
        [collaboratorIdRoute]: collaboratorId,
      },
    })

    return Marshaler.payloadSingleResponse(result)
  }

  async getByUserId(entityId, userId) {
    return this._getById(entityId, userId, true)
  }

  async getByOrganizationId(entityId, organizationId) {
    return this._getById(entityId, organizationId, false)
  }

  async getAll(entityId, params) {
    const entityIdRoute = this._parentRoutes.list
    const result = await this._api.ListCollaborators(
      {
        routeParams: { [entityIdRoute]: entityId },
      },
      params,
    )

    return Marshaler.payloadListResponse('collaborators', result)
  }

  async add(entityId, data) {
    const entityIdRoute = this._parentRoutes.set
    const result = await this._api.SetCollaborator(
      {
        routeParams: { [entityIdRoute]: entityId },
      },
      {
        collaborator: data,
      },
    )

    return Marshaler.payloadSingleResponse(result)
  }

  async update(entityId, data) {
    return await this.add(entityId, data)
  }

  async remove(isUser, entityId, collaboratorId) {
    const entityIdRoute = this._parentRoutes.delete
    const collaboratorIdRoute = isUser
      ? 'collaborator_ids.user_ids.user_id'
      : 'collaborator_ids.organization_ids.organization_id'

    const result = await this._api.DeleteCollaborator({
      routeParams: {
        [entityIdRoute]: entityId,
        [collaboratorIdRoute]: collaboratorId,
      },
    })

    return Marshaler.payloadSingleResponse(result)
  }
}

export default Collaborators

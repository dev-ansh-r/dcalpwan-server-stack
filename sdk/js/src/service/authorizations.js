

import autoBind from 'auto-bind'

import Marshaler from '../util/marshaler'

class Authorizations {
  constructor(registry) {
    this._api = registry

    autoBind(this)
  }

  async getAllAuthorizations(userId, params) {
    const result = await this._api.OAuthAuthorizationRegistry.List(
      {
        routeParams: { 'user_ids.user_id': userId },
      },
      params,
    )

    return Marshaler.payloadListResponse('authorizations', result)
  }

  async deleteAuthorization(userId, client_id) {
    const result = await this._api.OAuthAuthorizationRegistry.Delete({
      routeParams: { 'user_ids.user_id': userId, 'client_ids.client_id': client_id },
    })

    return Marshaler.payloadSingleResponse(result)
  }

  async getAllTokens(userId, client_id, params) {
    const result = await this._api.OAuthAuthorizationRegistry.ListTokens(
      {
        routeParams: { 'user_ids.user_id': userId, 'client_ids.client_id': client_id },
      },
      params,
    )

    return Marshaler.payloadListResponse('tokens', result)
  }

  async deleteToken(userId, client_id, id) {
    const result = await this._api.OAuthAuthorizationRegistry.DeleteToken({
      routeParams: {
        'user_ids.user_id': userId,
        'client_ids.client_id': client_id,
        id,
      },
    })

    return Marshaler.payloadSingleResponse(result)
  }
}

export default Authorizations

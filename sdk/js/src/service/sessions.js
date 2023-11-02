

import autoBind from 'auto-bind'

import Marshaler from '../util/marshaler'

class Sessions {
  constructor(registry) {
    this._api = registry

    autoBind(this)
  }

  async getAllSessions(userId, params) {
    const result = await this._api.UserSessionRegistry.List(
      {
        routeParams: { 'user_ids.user_id': userId },
      },
      params,
    )

    return Marshaler.payloadListResponse('sessions', result)
  }

  async deleteSession(userId, session_id) {
    const result = await this._api.UserSessionRegistry.Delete({
      routeParams: { 'user_ids.user_id': userId, session_id },
    })

    return Marshaler.payloadListResponse('sessions', result)
  }
}

export default Sessions

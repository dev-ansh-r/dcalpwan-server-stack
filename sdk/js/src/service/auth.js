

import autoBind from 'auto-bind'

import Marshaler from '../util/marshaler'

class Auth {
  constructor(service) {
    this._api = service
    autoBind(this)
  }

  async getAuthInfo() {
    const result = await this._api.AuthInfo()

    return Marshaler.payloadSingleResponse(result)
  }
}

export default Auth

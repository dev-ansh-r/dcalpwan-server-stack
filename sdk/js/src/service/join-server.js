

import autoBind from 'auto-bind'

import Marshaler from '../util/marshaler'

class Js {
  constructor(service) {
    this._api = service
    autoBind(this)
  }

  async listJoinEUIPrefixes() {
    const result = await this._api.GetJoinEUIPrefixes()

    return Marshaler.payloadListResponse('prefixes', result)
  }
}

export default Js

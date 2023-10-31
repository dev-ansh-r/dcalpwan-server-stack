

import autoBind from 'auto-bind'

import Marshaler from '../util/marshaler'

class Is {
  constructor(service) {
    this._api = service
    autoBind(this)
  }

  async getConfiguration() {
    const result = await this._api.GetConfiguration()

    return Marshaler.payloadSingleResponse(result)
  }
}

export default Is

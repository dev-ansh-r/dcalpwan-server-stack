

import autoBind from 'auto-bind'

import Marshaler from '../util/marshaler'

class ContactInfo {
  constructor(service) {
    this._api = service
    autoBind(this)
  }

  async validate(token, id) {
    const result = await this._api.Validate(undefined, token, id)

    return Marshaler.payloadSingleResponse(result)
  }

  async requestValidation(ids) {
    const result = await this._api.RequestValidation(undefined, ids)

    return Marshaler.payloadSingleResponse(result)
  }
}

export default ContactInfo

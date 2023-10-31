

import autoBind from 'auto-bind'

import Marshaler from '../util/marshaler'

class QRCodeGenerator {
  constructor(registry, { stackConfig }) {
    this._api = registry
    this._stackConfig = stackConfig
    autoBind(this)
  }

  async parse(qrCode) {
    const response = await this._api.EndDeviceQRCodeGenerator.Parse(undefined, {
      qr_code: btoa(qrCode),
    })

    return Marshaler.payloadSingleResponse(response)
  }
}

export default QRCodeGenerator

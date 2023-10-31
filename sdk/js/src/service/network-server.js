

import autoBind from 'auto-bind'

import Marshaler from '../util/marshaler'

class Ns {
  constructor(service) {
    this._api = service
    autoBind(this)
  }

  async generateDevAddress() {
    const result = await this._api.GenerateDevAddr()

    return Marshaler.payloadSingleResponse(result)
  }

  async getDefaultMacSettings(freqPlan, phyVersion) {
    const result = await this._api.GetDefaultMACSettings({
      routeParams: {
        frequency_plan_id: freqPlan,
        lorawan_phy_version: phyVersion,
      },
    })

    return Marshaler.payloadSingleResponse(result)
  }
}

export default Ns

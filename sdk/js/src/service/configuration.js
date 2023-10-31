

import autoBind from 'auto-bind'

import Marshaler from '../util/marshaler'

class Configuration {
  constructor(service) {
    this._api = service
    autoBind(this)
  }

  async listNsFrequencyPlans() {
    const result = await this._api.ListFrequencyPlans({ component: 'ns' })

    return Marshaler.payloadListResponse('frequency_plans', result)
  }

  async listGsFrequencyPlans() {
    const result = await this._api.ListFrequencyPlans({ component: 'gs' })

    return Marshaler.payloadListResponse('frequency_plans', result)
  }

  async getPhyVersions() {
    const result = await this._api.GetPhyVersions()
    return Marshaler.payloadSingleResponse(result)
  }
}

export default Configuration

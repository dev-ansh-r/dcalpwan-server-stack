

import autoBind from 'auto-bind'

import Marshaler from '../util/marshaler'

class SearchAccounts {
  constructor(registry) {
    this._api = registry

    autoBind(this)
  }

  async searchAccounts(params) {
    const result = await this._api.EntityRegistrySearch.SearchAccounts(undefined, {
      ...params,
    })

    return Marshaler.payloadListResponse('account_ids', result)
  }
}

export default SearchAccounts

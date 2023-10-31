

import createRequestLogic from '@ttn-lw/lib/store/logics/create-request-logic'
import * as search from '@ttn-lw/lib/store/actions/search-accounts'

export default tts => {
  const getUserLogic = createRequestLogic({
    type: search.SEARCH_ACCOUNTS,
    process: async ({ action }) => {
      const { query, onlyUsers, collaboratorOf } = action.payload

      const request = {
        query,
        only_users: onlyUsers,
      }
      if (Boolean(collaboratorOf)) {
        request[collaboratorOf.path] = collaboratorOf.id
      }

      return await tts.SearchAccounts.searchAccounts(request)
    },
  })

  return getUserLogic
}

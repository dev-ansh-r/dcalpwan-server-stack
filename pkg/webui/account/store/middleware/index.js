

import status from '@ttn-lw/lib/store/logics/status'

import init from './init'
import user from './user'
import identityServer from './identity-server'
import sessions from './sessions'
import clients from './clients'
import collaborators from './collaborators'
import authorizations from './authorizations'
import searchAccounts from './search-accounts'

export default [
  ...status,
  ...init,
  ...user,
  ...identityServer,
  ...sessions,
  ...clients,
  ...collaborators,
  ...authorizations,
  ...searchAccounts,
]

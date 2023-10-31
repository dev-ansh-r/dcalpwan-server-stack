

import { combineReducers } from 'redux'

import { getClientId, getCollaboratorId } from '@ttn-lw/lib/selectors/id'
import init from '@ttn-lw/lib/store/reducers/init'
import status from '@ttn-lw/lib/store/reducers/status'
import { createNamedPaginationReducer } from '@ttn-lw/lib/store/reducers/pagination'
import fetching from '@ttn-lw/lib/store/reducers/ui/fetching'
import error from '@ttn-lw/lib/store/reducers/ui/error'
import collaborators from '@ttn-lw/lib/store/reducers/collaborators'
import searchAccounts from '@ttn-lw/lib/store/reducers/search-accounts'

import user from './user'
import is from './identity-server'
import session from './sessions'
import clients from './clients'
import authorizations from './authorizations'

export default combineReducers({
  init,
  clients,
  authorizations,
  user,
  session,
  is,
  ui: combineReducers({
    fetching,
    error,
  }),
  pagination: combineReducers({
    clients: createNamedPaginationReducer('CLIENTS', getClientId),
    collaborators: createNamedPaginationReducer('COLLABORATORS', getCollaboratorId),
  }),
  status,
  collaborators,
  searchAccounts,
})

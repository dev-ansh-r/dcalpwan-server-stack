

import { createSelector } from 'reselect'

import { createPaginationIdsSelectorByEntity } from '@ttn-lw/lib/store/selectors/pagination'
import { createFetchingSelector } from '@ttn-lw/lib/store/selectors/fetching'
import { createErrorSelector } from '@ttn-lw/lib/store/selectors/error'

import {
  GET_CLIENT_BASE,
  GET_CLIENTS_LIST_BASE,
  GET_CLIENT_RIGHTS_BASE,
} from '@account/store/actions/clients'

const selectClientsStore = state => state.clients

export const selectClientEntitiesStore = state => selectClientsStore(state).entities
export const selectClientById = (state, id) => selectClientEntitiesStore(state)[id]
export const selectSelectedClientId = state => selectClientsStore(state).selectedClient
export const selectSelectedClient = state => selectClientById(state, selectSelectedClientId(state))

const selectClientsIds = createPaginationIdsSelectorByEntity('clients')
const selectClientsFetching = createFetchingSelector(GET_CLIENTS_LIST_BASE)
const selectClientsError = createErrorSelector(GET_CLIENTS_LIST_BASE)
export const selectClientFetching = createFetchingSelector(GET_CLIENT_BASE)
export const selectClientError = createErrorSelector(GET_CLIENT_BASE)

export const selectOAuthClients = state =>
  selectClientsIds(state).map(id => selectClientById(state, id))
export const selectOAuthClientsTotalCount = state => selectClientsStore(state).totalCount
export const selectOAuthClientsFetching = state => selectClientsFetching(state)
export const selectOAuthClientsError = state => selectClientsError(state)

// Rights.
export const selectClientRights = createSelector(selectClientsStore, ({ rights }) => [
  ...rights.regular,
  ...rights.pseudo,
])
export const selectClientRegularRights = state => selectClientsStore(state).rights?.regular
export const selectClientPseudoRights = state => selectClientsStore(state).rights?.pseudo
export const selectClientRightsError = createErrorSelector(GET_CLIENT_RIGHTS_BASE)
export const selectClientRightsFetching = createFetchingSelector(GET_CLIENT_RIGHTS_BASE)



import {
  createPaginationIdsSelectorByEntity,
  createPaginationTotalCountSelectorByEntity,
} from '@ttn-lw/lib/store/selectors/pagination'
import { createFetchingSelector } from '@ttn-lw/lib/store/selectors/fetching'
import { createErrorSelector } from '@ttn-lw/lib/store/selectors/error'
import { getOrganizationId } from '@ttn-lw/lib/selectors/id'

import { GET_ORGS_LIST_BASE, GET_ORGS_RIGHTS_LIST_BASE } from '@console/store/actions/organizations'

import {
  createEventsSelector,
  createEventsErrorSelector,
  createEventsStatusSelector,
  createEventsInterruptedSelector,
  createEventsPausedSelector,
  createEventsTruncatedSelector,
  createEventsFilterSelector,
} from './events'
import { createRightsSelector, createPseudoRightsSelector } from './rights'

const ENTITY = 'organizations'

// Organization.
export const selectOrganizationStore = state => state.organizations
export const selectOrganizationEntitiesStore = state => selectOrganizationStore(state).entities
export const selectOrganizationById = (state, id) => selectOrganizationEntitiesStore(state)[id]
export const selectSelectedOrganizationId = state =>
  selectOrganizationStore(state).selectedOrganization
export const selectSelectedOrganization = state =>
  selectOrganizationById(state, selectSelectedOrganizationId(state))
export const selectOrganizationCollaboratorCounts = state =>
  selectOrganizationStore(state).collaboratorCounts
export const selectOrganizationCollaboratorCount = (state, id) =>
  selectOrganizationCollaboratorCounts(state)?.[id] || 0

// Organizations.
const selectOrgsIds = createPaginationIdsSelectorByEntity(ENTITY)
const selectOrgsTotalCount = createPaginationTotalCountSelectorByEntity(ENTITY)
const selectOrgsFetching = createFetchingSelector(GET_ORGS_LIST_BASE)
const selectOrgsError = createErrorSelector(GET_ORGS_LIST_BASE)

export const selectOrganizations = state =>
  selectOrgsIds(state).map(id => selectOrganizationById(state, id))
export const selectOrganizationsTotalCount = state => selectOrgsTotalCount(state)
export const selectOrganizationsFetching = state => selectOrgsFetching(state)
export const selectOrganizationsError = state => selectOrgsError(state)
export const selectOrganizationsWithCollaboratorCount = state =>
  selectOrganizations(state).map(org => ({
    ...org,
    _collaboratorCount: selectOrganizationCollaboratorCount(state, getOrganizationId(org)),
  }))

// Rights.
export const selectOrganizationRights = createRightsSelector(ENTITY)
export const selectOrganizationPseudoRights = createPseudoRightsSelector(ENTITY)
export const selectOrganizationRightsError = createErrorSelector(GET_ORGS_RIGHTS_LIST_BASE)
export const selectOrganizationRightsFetching = createFetchingSelector(GET_ORGS_RIGHTS_LIST_BASE)

// Events.
export const selectOrganizationEvents = createEventsSelector(ENTITY)
export const selectOrganizationEventsError = createEventsErrorSelector(ENTITY)
export const selectOrganizationEventsStatus = createEventsStatusSelector(ENTITY)
export const selectOrganizationEventsInterrupted = createEventsInterruptedSelector(ENTITY)
export const selectOrganizationEventsPaused = createEventsPausedSelector(ENTITY)
export const selectOrganizationEventsTruncated = createEventsTruncatedSelector(ENTITY)
export const selectOrganizationEventsFilter = createEventsFilterSelector(ENTITY)

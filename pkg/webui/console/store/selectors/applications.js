

import {
  createPaginationIdsSelectorByEntity,
  createPaginationTotalCountSelectorByEntity,
} from '@ttn-lw/lib/store/selectors/pagination'
import { createFetchingSelector } from '@ttn-lw/lib/store/selectors/fetching'
import { createErrorSelector } from '@ttn-lw/lib/store/selectors/error'

import { GET_APP_LINK_BASE } from '@console/store/actions/link'
import {
  GET_APPS_RIGHTS_LIST_BASE,
  GET_APP_BASE,
  GET_APP_DEV_COUNT_BASE,
} from '@console/store/actions/applications'

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

const ENTITY = 'applications'

// Application.
export const selectApplicationStore = state => state.applications
export const selectApplicationEntitiesStore = state => selectApplicationStore(state).entities
export const selectApplicationDerivedStore = state => selectApplicationStore(state).derived
export const selectApplicationDerivedById = (state, id) => selectApplicationDerivedStore(state)[id]
export const selectApplicationById = (state, id) => selectApplicationEntitiesStore(state)[id]
export const selectSelectedApplicationId = state =>
  selectApplicationStore(state).selectedApplication
export const selectApplicationNameById = (state, id) =>
  (selectApplicationById(state, id) || {}).name
export const selectSelectedApplication = state =>
  selectApplicationById(state, selectSelectedApplicationId(state))
export const selectApplicationFetching = createFetchingSelector(GET_APP_BASE)
export const selectApplicationError = createErrorSelector(GET_APP_BASE)
export const selectApplicationDeviceCount = (state, id) =>
  selectApplicationStore(state).applicationDeviceCounts[id]
export const selectApplicationDeviceCountFetching = createFetchingSelector(GET_APP_DEV_COUNT_BASE)
export const selectApplicationDeviceCountError = createErrorSelector(GET_APP_DEV_COUNT_BASE)
export const selectApplicationDerivedLastSeen = (state, id) =>
  (selectApplicationDerivedById(state, id) || {}).lastSeen
export const selectApplicationDevEUICount = state =>
  selectApplicationById(state, selectSelectedApplicationId(state)).dev_eui_counter || 0

// Applications.
const selectAppsIds = createPaginationIdsSelectorByEntity(ENTITY)
const selectAppsTotalCount = createPaginationTotalCountSelectorByEntity(ENTITY)

export const selectApplications = state =>
  selectAppsIds(state).map(id => selectApplicationById(state, id))
export const selectApplicationsTotalCount = state => selectAppsTotalCount(state)
export const selectApplicationsWithDeviceCounts = state =>
  selectApplications(state).map(app => ({
    ...app,
    _devices: selectApplicationDeviceCount(state, app.ids.application_id),
  }))

// Events.
export const selectApplicationEvents = createEventsSelector(ENTITY)
export const selectApplicationEventsError = createEventsErrorSelector(ENTITY)
export const selectApplicationEventsStatus = createEventsStatusSelector(ENTITY)
export const selectApplicationEventsInterrupted = createEventsInterruptedSelector(ENTITY)
export const selectApplicationEventsPaused = createEventsPausedSelector(ENTITY)
export const selectApplicationEventsTruncated = createEventsTruncatedSelector(ENTITY)
export const selectApplicationEventsFilter = createEventsFilterSelector(ENTITY)

// Rights.
export const selectApplicationRights = createRightsSelector(ENTITY)
export const selectApplicationPseudoRights = createPseudoRightsSelector(ENTITY)
export const selectApplicationRightsError = createErrorSelector(GET_APPS_RIGHTS_LIST_BASE)

// Link.
const selectLinkStore = state => state.link
export const selectApplicationLink = state => selectLinkStore(state).link
export const selectApplicationLinkError = createErrorSelector(GET_APP_LINK_BASE)
export const selectApplicationLinkFormatters = state => {
  const link = selectApplicationLink(state) || {}

  return link.default_formatters
}

export const selectApplicationLinkSkipPayloadCrypto = state => {
  const link = selectApplicationLink(state) || {}

  return link.skip_payload_crypto || false
}

export const selectMqttConnectionInfo = state => selectApplicationStore(state).mqtt

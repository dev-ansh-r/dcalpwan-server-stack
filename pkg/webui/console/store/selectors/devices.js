

import getHostnameFromUrl from '@ttn-lw/lib/host-from-url'
import { combineDeviceIds, extractDeviceIdFromCombinedId } from '@ttn-lw/lib/selectors/id'
import {
  createPaginationIdsSelectorByEntity,
  createPaginationTotalCountSelectorByEntity,
} from '@ttn-lw/lib/store/selectors/pagination'
import { selectAsConfig, selectJsConfig, selectNsConfig } from '@ttn-lw/lib/selectors/env'

import {
  createEventsSelector,
  createEventsErrorSelector,
  createEventsStatusSelector,
  createEventsInterruptedSelector,
  createEventsPausedSelector,
  createEventsTruncatedSelector,
  createEventsFilterSelector,
} from './events'

const ENTITY = 'devices'

// Device.
export const selectDeviceStore = state => state.devices
export const selectDeviceEntitiesStore = state => selectDeviceStore(state).entities
export const selectDeviceDerivedStore = state => selectDeviceStore(state).derived
export const selectDeviceByIds = (state, appId, devId) =>
  selectDeviceById(state, combineDeviceIds(appId, devId))
export const selectDeviceById = (state, id) => selectDeviceEntitiesStore(state)[id]
export const selectDeviceDerivedById = (state, id) => selectDeviceDerivedStore(state)[id]
export const selectSelectedDeviceId = state =>
  extractDeviceIdFromCombinedId(selectDeviceStore(state).selectedDevice)
export const selectSelectedCombinedDeviceId = state => selectDeviceStore(state).selectedDevice
export const selectSelectedDevice = state =>
  selectDeviceById(state, selectSelectedCombinedDeviceId(state))
export const selectSelectedDeviceFormatters = state => selectSelectedDevice(state).formatters
export const isOtherClusterDevice = device => {
  const isOtherCluster =
    getHostnameFromUrl(selectAsConfig().base_url) !== device.application_server_address &&
    getHostnameFromUrl(selectNsConfig().base_url) !== device.network_server_address &&
    getHostnameFromUrl(selectJsConfig().base_url) !== device.join_server_address

  return isOtherCluster
}
export const selectDeviceLastSeen = (state, appId, devId) => {
  const device = selectDeviceById(state, combineDeviceIds(appId, devId))
  if (!Boolean(device)) return undefined

  return device.last_seen_at
}

// Derived.
export const selectDeviceDerivedUplinkFrameCount = (state, appId, devId) => {
  const derived = selectDeviceDerivedById(state, combineDeviceIds(appId, devId))
  if (!Boolean(derived)) return undefined

  return derived.uplinkFrameCount
}
export const selectDeviceDerivedDownlinkFrameCount = (state, appId, devId) => {
  const derived = selectDeviceDerivedById(state, combineDeviceIds(appId, devId))
  if (!Boolean(derived)) return undefined

  return derived.downlinkFrameCount
}

export const selectSelectedDeviceClaimable = state =>
  selectDeviceStore(state).selectedDeviceClaimable

export const selectDevicesWithLastSeen = state => {
  const devices = selectDevices(state)
  const devicesWithLastSeen = devices.map(device => ({
    ...device,
    _lastSeen: selectDeviceLastSeen(state, device.application_ids, device.ids.device_id),
  }))
  return devicesWithLastSeen
}

// Devices.
const selectDevsIds = createPaginationIdsSelectorByEntity(ENTITY)
const selectDevsTotalCount = createPaginationTotalCountSelectorByEntity(ENTITY)

export const selectDevices = state => selectDevsIds(state).map(id => selectDeviceById(state, id))
export const selectDevicesTotalCount = state => selectDevsTotalCount(state)

// Events.
export const selectDeviceEvents = createEventsSelector(ENTITY)
export const selectDeviceEventsError = createEventsErrorSelector(ENTITY)
export const selectDeviceEventsStatus = createEventsStatusSelector(ENTITY)
export const selectDeviceEventsInterruptted = createEventsInterruptedSelector(ENTITY)
export const selectDeviceEventsPaused = createEventsPausedSelector(ENTITY)
export const selectDeviceEventsTruncated = createEventsTruncatedSelector(ENTITY)
export const selectDeviceEventsFilter = createEventsFilterSelector(ENTITY)

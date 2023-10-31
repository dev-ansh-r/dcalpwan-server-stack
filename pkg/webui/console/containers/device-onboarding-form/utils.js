

import { ACTIVATION_MODES, hasSelectedDeviceRepositoryOther } from '@console/lib/device-utils'

// End device selectors.

export const getActivationMode = device =>
  device.supports_join === true
    ? ACTIVATION_MODES.OTAA
    : device.multicast === true
    ? ACTIVATION_MODES.MULTICAST
    : device.supports_join === false && device.multicast === false
    ? ACTIVATION_MODES.ABP
    : ACTIVATION_MODES.NONE

export const getLorawanVersion = device => device.lorawan_version || '1.1.0'

export const getApplicationServerAddress = device => device.application_server_address
export const getNetworkServerAddress = device => device.network_server_address
export const getJoinServerAddress = device => device.join_server_address

// End device repository utils.

export const hasCompletedDeviceRepositorySelection = version =>
  version && Object.values(version).every(value => value)
export const hasValidDeviceRepositoryType = (version, template) =>
  !hasSelectedDeviceRepositoryOther(version) &&
  hasCompletedDeviceRepositorySelection(version) &&
  Boolean(template)
export const mayProvisionDevice = (values, template) => {
  const { frequency_plan_id, lorawan_version, lorawan_phy_version, _inputMethod, version_ids } =
    values

  return (
    (_inputMethod === 'manual' &&
      Boolean(frequency_plan_id) &&
      Boolean(lorawan_version) &&
      Boolean(lorawan_phy_version)) ||
    (_inputMethod === 'device-repository' &&
      hasValidDeviceRepositoryType(version_ids, template) &&
      Boolean(frequency_plan_id))
  )
}

export const REGISTRATION_TYPES = Object.freeze({
  SINGLE: 'single',
  MULTIPLE: 'multiple',
})

export const DEVICE_CLASS_MAP = {
  CLASS_A: 'class-a',
  CLASS_B: 'class-b',
  CLASS_C: 'class-c',
  CLASS_B_C: 'class-b-c',
}

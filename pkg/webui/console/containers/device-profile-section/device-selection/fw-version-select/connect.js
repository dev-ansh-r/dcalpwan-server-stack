

import { connect } from 'react-redux'

import { isUnknownHwVersion } from '@console/lib/device-utils'

import { selectDeviceModelFirmwareVersions } from '@console/store/selectors/device-repository'

const mapStateToProps = (state, props) => {
  const { brandId, modelId, hwVersion } = props

  const fwVersions = selectDeviceModelFirmwareVersions(state, brandId, modelId).filter(
    ({ supported_hardware_versions = [] }) =>
      (Boolean(hwVersion) && supported_hardware_versions.includes(hwVersion)) ||
      // Include firmware versions when there are no hardware versions configured in device repository
      // for selected end device model.
      isUnknownHwVersion(hwVersion),
  )

  return {
    versions: fwVersions,
  }
}

export default FirmwareVersionSelect => connect(mapStateToProps)(FirmwareVersionSelect)

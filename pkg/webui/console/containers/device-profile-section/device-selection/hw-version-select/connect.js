

import { connect } from 'react-redux'

import { selectDeviceModelHardwareVersions } from '@console/store/selectors/device-repository'

const mapStateToProps = (state, props) => {
  const { brandId, modelId } = props

  return {
    versions: selectDeviceModelHardwareVersions(state, brandId, modelId),
  }
}

export default HardwareVersionSelect => connect(mapStateToProps)(HardwareVersionSelect)

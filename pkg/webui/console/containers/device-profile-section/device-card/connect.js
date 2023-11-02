

import { connect } from 'react-redux'

import { selectDeviceModelById } from '@console/store/selectors/device-repository'

const mapStateToProps = (state, props) => {
  const { brandId, modelId } = props

  return {
    model: selectDeviceModelById(state, brandId, modelId),
  }
}

export default DeviceCard => connect(mapStateToProps)(DeviceCard)

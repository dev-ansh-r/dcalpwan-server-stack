

import { connect } from 'react-redux'

import {
  selectDeviceByIds,
  selectDeviceDerivedUplinkFrameCount,
  selectDeviceDerivedDownlinkFrameCount,
  selectDeviceLastSeen,
} from '@console/store/selectors/devices'

const mapStateToProps = (state, props) => {
  const { devId, appId } = props

  return {
    devId,
    appId,
    device: selectDeviceByIds(state, appId, devId),
    uplinkFrameCount: selectDeviceDerivedUplinkFrameCount(state, appId, devId),
    downlinkFrameCount: selectDeviceDerivedDownlinkFrameCount(state, appId, devId),
    lastSeen: selectDeviceLastSeen(state, appId, devId),
  }
}

export default TitleSection => connect(mapStateToProps)(TitleSection)

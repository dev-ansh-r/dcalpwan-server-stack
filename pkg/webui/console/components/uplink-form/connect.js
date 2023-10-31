

import { connect } from 'react-redux'

import tts from '@console/api/tts'

import {
  selectSelectedApplicationId,
  selectApplicationLinkSkipPayloadCrypto,
} from '@console/store/selectors/applications'
import { selectSelectedDeviceId, selectSelectedDevice } from '@console/store/selectors/devices'

const mapStateToProps = state => {
  const appId = selectSelectedApplicationId(state)
  const devId = selectSelectedDeviceId(state)
  const device = selectSelectedDevice(state)
  const skipPayloadCrypto = selectApplicationLinkSkipPayloadCrypto(state)

  return {
    appId,
    devId,
    device,
    simulateUplink: uplink => tts.Applications.Devices.simulateUplink(appId, devId, uplink),
    skipPayloadCrypto,
  }
}

export default UplinkForm => connect(mapStateToProps)(UplinkForm)

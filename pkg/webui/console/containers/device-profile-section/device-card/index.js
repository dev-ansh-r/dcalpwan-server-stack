

import connect from './connect'
import DeviceCard from './device-card'

const ConnectedDeviceCard = connect(DeviceCard)

export { ConnectedDeviceCard as default, DeviceCard }

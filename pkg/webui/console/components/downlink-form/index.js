

import DownlinkForm from './downlink-form'
import connect from './connect'

const ConnectedDownlinkForm = connect(DownlinkForm)

export { ConnectedDownlinkForm as default, DownlinkForm }

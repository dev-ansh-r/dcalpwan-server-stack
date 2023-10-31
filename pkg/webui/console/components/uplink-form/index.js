

import UplinkForm from './uplink-form'
import connect from './connect'

const ConnectedUplinkForm = connect(UplinkForm)

export { ConnectedUplinkForm as default, UplinkForm }

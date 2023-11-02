

import connect from './connect'
import HardwareVersionSelect from './hw-version-select'

const ConnectedHardwareVersionSelect = connect(HardwareVersionSelect)

export { ConnectedHardwareVersionSelect as default, HardwareVersionSelect }

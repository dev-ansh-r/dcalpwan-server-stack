

import connect from './connect'
import FirmwareVersionSelect from './fw-version-select'

const ConnectedFirmwareVersionSelect = connect(FirmwareVersionSelect)

export { ConnectedFirmwareVersionSelect as default, FirmwareVersionSelect }

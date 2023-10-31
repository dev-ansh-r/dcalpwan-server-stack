

import connect from './connect'
import OwnersSelect from './owners-select'

const ConnectedOwnersSelect = connect(OwnersSelect)

export { ConnectedOwnersSelect as default, OwnersSelect }

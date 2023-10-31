

import { FullViewError, FullViewErrorInner } from './error'
import connect from './connect'

const ConnectedFullErrorView = connect(FullViewError)

export { ConnectedFullErrorView as default, FullViewError, FullViewErrorInner }

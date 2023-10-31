

import Events from './organization-events'
import connect from './connect'

const ConnectedOrganizationEvents = connect(Events)

export { ConnectedOrganizationEvents as default, Events }

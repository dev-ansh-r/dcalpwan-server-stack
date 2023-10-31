

import Connection from './gateway-connection'
import connect from './connect'

const ConnectedGatewayConnection = connect(Connection)

export { ConnectedGatewayConnection as default, Connection }

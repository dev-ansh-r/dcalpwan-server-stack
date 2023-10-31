

const isGsEvent = eventName => Boolean(eventName) && eventName.startsWith('gs.')
const isGsUplinkEvent = eventName => isGsEvent(eventName) && eventName.includes('.up.')
const isGsDownlinkEvent = eventName => isGsEvent(eventName) && eventName.includes('.down.')
const isGsGatewayEvent = eventName => isGsEvent(eventName) && eventName.includes('.gateway.')
export const isGsStatusReceiveEvent = eventName =>
  isGsEvent(eventName) && eventName.includes('.status.receive')
export const isGsUplinkReceiveEvent = eventName =>
  isGsUplinkEvent(eventName) && eventName.endsWith('.receive')
export const isGsDownlinkSendEvent = eventName =>
  isGsDownlinkEvent(eventName) && eventName.includes('.send')
export const isGsGatewayConnectedEvent = eventName =>
  isGsGatewayEvent(eventName) && eventName.includes('.connect')
export const isGsGatewayDisconnectedEvent = eventName =>
  isGsGatewayEvent(eventName) && eventName.includes('.disconnect')

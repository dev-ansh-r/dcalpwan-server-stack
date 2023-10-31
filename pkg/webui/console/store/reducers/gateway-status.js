

import { handleActions } from 'redux-actions'

import {
  isGsStatusReceiveEvent,
  isGsUplinkReceiveEvent,
  isGsGatewayConnectedEvent,
  isGsGatewayDisconnectedEvent,
  isGsDownlinkSendEvent,
} from '@ttn-lw/lib/selectors/event'

import {
  GET_GTW,
  UPDATE_GTW_STATS_SUCCESS,
  GET_GTW_EVENT_MESSAGE_SUCCESS,
} from '@console/store/actions/gateways'

const handleStatsUpdate = (state, { stats = {} }) => {
  const lastSeen = [
    stats.last_status_received_at,
    stats.last_uplink_received_at,
    stats.last_downlink_received_at,
    stats.connected_at,
    stats.disconnected_at,
  ]
    .filter(a => Boolean(a))
    .reduce((max, current) => (max > current ? max : current), state.lastSeen)

  if (lastSeen === state.lastSeen) {
    return state
  }

  return { ...state, lastSeen }
}

const handleEventUpdate = (state, event) => {
  if (
    isGsStatusReceiveEvent(event.name) ||
    isGsUplinkReceiveEvent(event.name) ||
    isGsDownlinkSendEvent(event.name) ||
    isGsGatewayConnectedEvent(event.name) ||
    isGsGatewayDisconnectedEvent(event.name)
  ) {
    const lastSeen = state.lastSeen > event.time ? state.lastSeen : event.time
    return { ...state, lastSeen }
  }

  return state
}

const defaultState = { lastSeen: undefined }

/**
 * The `gatewayStatus` reducer contains connection status information of the
 * current gateway. The connection status is deducted from gateway stats and
 * gateway status events.
 */
const gatewayStatus = handleActions(
  {
    [GET_GTW]: () => defaultState,
    [UPDATE_GTW_STATS_SUCCESS]: (state, { payload }) => handleStatsUpdate(state, payload),
    [GET_GTW_EVENT_MESSAGE_SUCCESS]: (state, { event }) => handleEventUpdate(state, event),
  },
  defaultState,
)

export { gatewayStatus as default, defaultState }



const selectGatewayStatusStore = state => state.gatewayStatus

/* eslint-disable-next-line import/prefer-default-export */
export const selectGatewayLastSeen = state => selectGatewayStatusStore(state).lastSeen

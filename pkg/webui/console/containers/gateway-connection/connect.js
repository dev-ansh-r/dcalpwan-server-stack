

import { connect } from 'react-redux'

import { selectGsConfig } from '@ttn-lw/lib/selectors/env'
import getHostFromUrl from '@ttn-lw/lib/host-from-url'

import {
  startGatewayStatistics,
  stopGatewayStatistics,
  updateGatewayStatistics,
} from '@console/store/actions/gateways'

import {
  selectGatewayStatistics,
  selectGatewayStatisticsError,
  selectGatewayStatisticsIsFetching,
  selectLatestGatewayEvent,
  selectGatewayById,
} from '@console/store/selectors/gateways'
import { selectGatewayLastSeen } from '@console/store/selectors/gateway-status'

import withConnectionReactor from './gateway-connection-reactor'

export default GatewayConnection =>
  connect(
    (state, ownProps) => {
      const gateway = selectGatewayById(state, ownProps.gtwId)
      const gsConfig = selectGsConfig()
      const consoleGsAddress = getHostFromUrl(gsConfig.base_url)
      const gatewayServerAddress = getHostFromUrl(gateway.gateway_server_address)

      return {
        statistics: selectGatewayStatistics(state, ownProps),
        error: selectGatewayStatisticsError(state, ownProps),
        fetching: selectGatewayStatisticsIsFetching(state, ownProps),
        latestEvent: selectLatestGatewayEvent(state, ownProps.gtwId),
        lastSeen: selectGatewayLastSeen(state),
        isOtherCluster: consoleGsAddress !== gatewayServerAddress,
      }
    },
    (dispatch, ownProps) => ({
      startStatistics: () => dispatch(startGatewayStatistics(ownProps.gtwId)),
      stopStatistics: () => dispatch(stopGatewayStatistics()),
      updateGatewayStatistics: () => dispatch(updateGatewayStatistics(ownProps.gtwId)),
    }),
  )(withConnectionReactor(GatewayConnection))

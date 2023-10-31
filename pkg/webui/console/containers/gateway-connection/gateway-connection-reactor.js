

import React from 'react'

import {
  isGsStatusReceiveEvent,
  isGsDownlinkSendEvent,
  isGsUplinkReceiveEvent,
} from '@ttn-lw/lib/selectors/event'
import PropTypes from '@ttn-lw/lib/prop-types'

/**
 * `withConnectionReactor` is a HOC that handles gateway connection statistics
 * updates based on gateway uplink, downlink and connection events.
 *
 * @param {object} Component - React component to be wrapped by the reactor.
 * @returns {object} - A wrapped react component.
 */
const withConnectionReactor = Component => {
  class ConnectionReactor extends React.PureComponent {
    componentDidUpdate(prevProps) {
      const { latestEvent, updateGatewayStatistics } = this.props

      if (Boolean(latestEvent) && latestEvent !== prevProps.latestEvent) {
        const { name } = latestEvent
        const isHeartBeatEvent =
          isGsDownlinkSendEvent(name) ||
          isGsUplinkReceiveEvent(name) ||
          isGsStatusReceiveEvent(name)

        if (isHeartBeatEvent) {
          updateGatewayStatistics()
        }
      }
    }

    render() {
      const { latestEvent, updateGatewayStatistics, ...rest } = this.props
      return <Component {...rest} />
    }
  }

  ConnectionReactor.propTypes = {
    latestEvent: PropTypes.event,
    updateGatewayStatistics: PropTypes.func.isRequired,
  }

  ConnectionReactor.defaultProps = {
    latestEvent: undefined,
  }

  return ConnectionReactor
}

export default withConnectionReactor

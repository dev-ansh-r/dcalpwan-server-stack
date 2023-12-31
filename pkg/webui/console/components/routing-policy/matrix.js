
import React from 'react'
import classnames from 'classnames'

import Icon from '@ttn-lw/components/icon'
import Tooltip from '@ttn-lw/components/tooltip'

import Message from '@ttn-lw/lib/components/message'

import PropTypes from '@ttn-lw/lib/prop-types'

import m from '@console/lib/packet-broker/messages'

import style from './routing-policy-matrix.styl'

const PolicyPoint = ({ symbol, enabled, positiveMessage, negativeMessage, tooltipPlacement }) => (
  <Tooltip
    content={<Message content={enabled ? positiveMessage : negativeMessage} />}
    placement={tooltipPlacement}
  >
    <div className={classnames({ [style.active]: enabled })}>{symbol}</div>
  </Tooltip>
)

PolicyPoint.propTypes = {
  enabled: PropTypes.bool,
  negativeMessage: PropTypes.message.isRequired,
  positiveMessage: PropTypes.message.isRequired,
  symbol: PropTypes.string.isRequired,
  tooltipPlacement: PropTypes.string,
}

PolicyPoint.defaultProps = {
  enabled: false,
  tooltipPlacement: 'top',
}

const RoutingPolicyMatrix = ({ policy }) => {
  const { uplink = {}, downlink = {} } = policy

  return (
    <div className={style.container}>
      <div className={style.uplink}>
        <Tooltip content={<Message content={m.uplinkPolicies} />} placement="top">
          <Icon small icon="uplink" className={style.icon} />
        </Tooltip>
        <PolicyPoint
          symbol="J"
          enabled={uplink.join_request}
          positiveMessage={m.forwardsJoinRequest}
          negativeMessage={m.doesNotForwardJoinRequest}
        />
        <PolicyPoint
          symbol="M"
          enabled={uplink.mac_data}
          positiveMessage={m.forwardsMacData}
          negativeMessage={m.doesNotForwardMacData}
        />
        <PolicyPoint
          symbol="A"
          enabled={uplink.application_data}
          positiveMessage={m.forwardsApplicationData}
          negativeMessage={m.doesNotForwardApplicationData}
        />
        <PolicyPoint
          symbol="S"
          enabled={uplink.signal_quality}
          positiveMessage={m.forwardsSignalQuality}
          negativeMessage={m.doesNotForwardSignalQuality}
        />
        <PolicyPoint
          symbol="L"
          enabled={uplink.localization}
          positiveMessage={m.forwardsLocalization}
          negativeMessage={m.doesNotForwardLocalization}
        />
      </div>
      <div className={style.downlink}>
        <Tooltip content={<Message content={m.downlinkPolicies} />} placement="bottom">
          <Icon small icon="downlink" className={style.icon} />
        </Tooltip>
        <PolicyPoint
          symbol="J"
          enabled={downlink.join_accept}
          positiveMessage={m.allowsJoinAccept}
          negativeMessage={m.doesNotAllowJoinAccept}
          tooltipPlacement="bottom"
        />
        <PolicyPoint
          symbol="M"
          enabled={downlink.mac_data}
          positiveMessage={m.allowsMacData}
          negativeMessage={m.doesNotAllowMacData}
          tooltipPlacement="bottom"
        />
        <PolicyPoint
          symbol="A"
          enabled={downlink.application_data}
          positiveMessage={m.allowsApplicationData}
          negativeMessage={m.doesNotAllowApplicationData}
          tooltipPlacement="bottom"
        />
      </div>
    </div>
  )
}

RoutingPolicyMatrix.propTypes = {
  policy: PropTypes.routingPolicy.isRequired,
}

export default RoutingPolicyMatrix

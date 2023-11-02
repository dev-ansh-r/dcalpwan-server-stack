
import React from 'react'

import RoutingPolicy from '..'

export default {
  title: 'Routing Policy/Matrix',
}

export const _Empty = () => (
  <div style={{ display: 'flex', height: '100vh', width: '100%' }}>
    <RoutingPolicy.Matrix
      policy={{
        uplink: {},
        downlink: {},
      }}
    />
  </div>
)

export const _Full = () => (
  <div style={{ display: 'flex', height: '100vh', width: '100%' }}>
    <RoutingPolicy.Matrix
      policy={{
        uplink: {
          join_request: true,
          mac_data: true,
          application_data: true,
          signal_quality: true,
          localization: true,
        },
        downlink: {
          join_accept: true,
          mac_data: true,
          application_data: true,
        },
      }}
    />
  </div>
)

export const _Random = () => (
  <div style={{ display: 'flex', height: '100vh', width: '100%' }}>
    <RoutingPolicy.Matrix
      policy={{
        uplink: {
          join_request: Math.random() >= 0.5,
          mac_data: Math.random() >= 0.5,
          application_data: Math.random() >= 0.5,
          signal_quality: Math.random() >= 0.5,
          localization: Math.random() >= 0.5,
        },
        downlink: {
          join_accept: Math.random() >= 0.5,
          mac_data: Math.random() >= 0.5,
          application_data: Math.random() >= 0.5,
        },
      }}
    />
  </div>
)

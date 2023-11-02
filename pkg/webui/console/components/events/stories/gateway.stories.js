
import React from 'react'
import { action } from '@storybook/addon-actions'

import Events from '..'

import { gatewayEvents } from './event-data'

export default {
  title: 'Events/Gateway',
}

export const Default = () => (
  <div style={{ display: 'flex', height: '100vh', width: '100%' }}>
    <Events
      entityId="test-app2"
      events={gatewayEvents}
      onPause={action('Gateway onPause')}
      onClear={action('Gateway onClear')}
      paused={false}
      truncated={false}
      scoped
    />
  </div>
)

export const Widget = () => (
  <div style={{ width: '540px' }}>
    <Events.Widget events={gatewayEvents} entityId="test-gtw-id" toAllUrl="/" scoped />
  </div>
)

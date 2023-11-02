
import React from 'react'
import { action } from '@storybook/addon-actions'

import Events from '..'

import { deviceEvents } from './event-data'

export default {
  title: 'Events/Device',
}

export const Default = () => (
  <div style={{ display: 'flex', height: '100vh', width: '100%' }}>
    <Events
      entityId="test-app2"
      events={deviceEvents}
      onClear={action('Device onClear')}
      onPause={action('Device onPause')}
      paused={false}
      truncated={false}
      scoped
    />
  </div>
)

export const Widget = () => (
  <div style={{ width: '540px' }}>
    <Events.Widget events={deviceEvents} entityId="test-dev-c" toAllUrl="/" scoped />
  </div>
)

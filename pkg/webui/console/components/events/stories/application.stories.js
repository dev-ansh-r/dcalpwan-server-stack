

import React from 'react'
import { action } from '@storybook/addon-actions'

import Events from '..'

import { events } from './event-data'

export default {
  title: 'Events/Application',
}

export const Default = () => (
  <div style={{ display: 'flex', height: '100vh', width: '100%' }}>
    <Events
      entityId="test-app2"
      events={events}
      onClear={action('Application onClear')}
      onPause={action('Application onPause')}
      paused={false}
      truncated={false}
    />
  </div>
)

export const Widget = () => (
  <div style={{ width: '540px' }}>
    <Events.Widget events={events} entityId="test-app2" toAllUrl="/" />
  </div>
)

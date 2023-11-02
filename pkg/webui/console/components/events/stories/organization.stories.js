
import React from 'react'
import { action } from '@storybook/addon-actions'

import Events from '..'

import { organizationEvents } from './event-data'

export default {
  title: 'Events/Organization',
}

export const Default = () => (
  <div style={{ display: 'flex', height: '100vh', width: '100%' }}>
    <Events
      entityId="test-app2"
      events={organizationEvents}
      onPause={action('Organization onPause')}
      onClear={action('Organization onClear')}
      paused={false}
      truncated={false}
      scoped
    />
  </div>
)

export const Widget = () => (
  <div style={{ width: '540px' }}>
    <Events.Widget events={organizationEvents} entityId="test-org-id" toAllUrl="/" scoped />
  </div>
)

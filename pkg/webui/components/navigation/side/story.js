

import React from 'react'
import { withInfo } from '@storybook/addon-info'

import SideNavigationItem from './item'

import { SideNavigation } from '.'

export default {
  title: 'Navigation',

  decorators: [
    withInfo({
      inline: true,
      header: false,
      source: false,
      propTables: [SideNavigation],
    }),
  ],
}

export const _SideNavigation = () => {
  const header = {
    title: 'test-application',
    icon: 'application',
    to: '/',
  }

  return (
    <div style={{ width: '300px', height: '700px' }}>
      <SideNavigation modifyAppContainerClasses={false} header={header}>
        <SideNavigationItem title="Overview" path="/" icon="overview" exact />
        <SideNavigationItem title="Devices" path="/devices" icon="devices" />
        <SideNavigationItem title="Data" path="/data" icon="data" />
        <SideNavigationItem title="Linking" path="/link" icon="link" />
        <SideNavigationItem title="Payload Formatters" icon="code">
          <SideNavigationItem title="Uplink" path="/payload-formatters/uplink" icon="uplink" />
          <SideNavigationItem
            title="Downlink"
            path="/payload-formatters/downlink"
            icon="downlink"
          />
        </SideNavigationItem>
        <SideNavigationItem title="Integrations" icon="integration">
          <SideNavigationItem title="MQTT" path="/integrations/mqtt" icon="extension" />
          <SideNavigationItem title="Webhooks" path="/integrations/webhooks" icon="extension" />
          <SideNavigationItem title="Pub/Subs" path="/integrations/pubsubs" icon="extension" />
        </SideNavigationItem>
        <SideNavigationItem title="Collaborators" path="/collaborators" icon="organization" />
        <SideNavigationItem title="API keys" path="/api-keys" icon="api_keys" />
        <SideNavigationItem
          title="General Settings"
          path="/general-settings"
          icon="general_settings"
        />
      </SideNavigation>
    </div>
  )
}

_SideNavigation.story = {
  name: 'SideNavigation',
}

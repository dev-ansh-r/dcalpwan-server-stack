

import React from 'react'
import 'focus-visible/dist/focus-visible'
import { withInfo } from '@storybook/addon-info'

import NavigationBar from '.'

export default {
  title: 'Navigation',

  decorators: [
    withInfo({
      inline: true,
      header: false,
      source: false,
      propTables: [NavigationBar],
    }),
  ],
}

export const _NavigationBar = () => (
  <NavigationBar>
    <NavigationBar.Item title="Overview" icon="overview" path="/overview" />
    <NavigationBar.Item title="Applications" icon="application" path="/application" />
    <NavigationBar.Item title="Gateways" icon="gateway" path="/gateways" />
    <NavigationBar.Item title="Organizations" icon="organization" path="/organization" />
  </NavigationBar>
)

_NavigationBar.story = {
  name: 'NavigationBar',
}

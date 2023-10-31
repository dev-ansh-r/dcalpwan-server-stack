

import React from 'react'
import { withInfo } from '@storybook/addon-info'

import Breadcrumb from './breadcrumb'
import { Breadcrumbs } from './breadcrumbs'

export default {
  title: 'Breadcrumbs',

  decorators: [
    withInfo({
      inline: true,
      header: false,
      source: false,
      propTables: [Breadcrumbs, Breadcrumb],
    }),
  ],
}

export const Default = () => {
  const breadcrumbs = [
    <Breadcrumb key="1" path="/applications" content="Applications" />,
    <Breadcrumb key="2" path="/applications/test-app" content="test-app" />,
    <Breadcrumb key="3" path="/applications/test-app/devices" content="Devices" />,
  ]

  return <Breadcrumbs breadcrumbs={breadcrumbs} />
}

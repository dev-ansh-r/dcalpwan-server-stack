
import React from 'react'

import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'
import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'

import PacketBrokerNetworksTable from '@console/containers/packet-broker-networks-table'

import m from './messages'

const NetworkRoutingPoliciesView = () => {
  useBreadcrumbs(
    'admin-panel.packet-broker.networks',
    <Breadcrumb path={'/admin-panel/packet-broker/networks'} content={m.networks} />,
  )

  return <PacketBrokerNetworksTable />
}

export default NetworkRoutingPoliciesView

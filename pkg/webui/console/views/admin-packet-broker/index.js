
import React from 'react'
import { Routes, Route } from 'react-router-dom'

import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'
import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'

import RequireRequest from '@ttn-lw/lib/components/require-request'
import ValidateRouteParam from '@ttn-lw/lib/components/validate-route-param'

import Require from '@console/lib/components/require'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import { pathId as pathIdRegexp } from '@ttn-lw/lib/regexp'

import { mayConfigurePacketBroker } from '@console/lib/feature-checks'

import { getPacketBrokerInfo } from '@console/store/actions/packet-broker'

import PacketBroker from './admin-packet-broker'
import NetworkRoutingPolicy from './network-routing-policy'

const PacketBrokerRouter = () => {
  useBreadcrumbs(
    'admin-panel.packet-broker',
    <Breadcrumb path={'/admin-panel/packet-broker'} content={sharedMessages.packetBroker} />,
  )

  return (
    <Require featureCheck={mayConfigurePacketBroker} otherwise={{ redirect: '/' }}>
      <RequireRequest requestAction={getPacketBrokerInfo()}>
        <Routes>
          <Route
            path="networks/:netId/:tenantId?"
            element={
              <ValidateRouteParam
                check={{ tenantId: pathIdRegexp, netId: /^\d+$/ }}
                Component={NetworkRoutingPolicy}
              />
            }
          />
          <Route path="*" Component={PacketBroker} />
        </Routes>
      </RequireRequest>
    </Require>
  )
}

export default PacketBrokerRouter

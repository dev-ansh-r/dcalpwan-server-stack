

import React from 'react'
import { Routes, Route } from 'react-router-dom'

import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'

import ValidateRouteParam from '@ttn-lw/lib/components/validate-route-param'

import Require from '@console/lib/components/require'

import Gateway from '@console/views/gateway'
import GatewayAdd from '@console/views/gateway-add'
import GatewaysList from '@console/views/gateways-list'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import { pathId as pathIdRegexp } from '@ttn-lw/lib/regexp'

import { mayViewGateways } from '@console/lib/feature-checks'

const Gateways = () => {
  useBreadcrumbs('gtws', <Breadcrumb path="/gateways" content={sharedMessages.gateways} />)
  return (
    <Require featureCheck={mayViewGateways} otherwise={{ redirect: '/' }}>
      <Routes>
        <Route index Component={GatewaysList} />
        <Route path="add" Component={GatewayAdd} />
        <Route
          path={`:gtwId/*`}
          element={<ValidateRouteParam check={{ gtwId: pathIdRegexp }} Component={Gateway} />}
        />
      </Routes>
    </Require>
  )
}

export default Gateways

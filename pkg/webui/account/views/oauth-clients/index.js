

import React from 'react'
import { Routes, Route } from 'react-router-dom'

import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'

import ValidateRouteParam from '@ttn-lw/lib/components/validate-route-param'
import GenericNotFound from '@ttn-lw/lib/components/full-view-error/not-found'

import OAuthClient from '@account/views/oauth-client'
import ClientsList from '@account/views/oauth-clients-list'
import OAuthClientAdd from '@account/views/oauth-client-add'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import { pathId as pathIdRegexp } from '@ttn-lw/lib/regexp'

const OAuthClients = () => {
  useBreadcrumbs(
    'clients',
    <Breadcrumb path="/oauth-clients" content={sharedMessages.oauthClients} />,
  )

  return (
    <Routes>
      <Route index Component={ClientsList} />
      <Route path="add" Component={OAuthClientAdd} />
      <Route
        path=":clientId/*"
        element={<ValidateRouteParam check={{ clientId: pathIdRegexp }} Component={OAuthClient} />}
      />
      <Route path="*" element={<GenericNotFound />} />
    </Routes>
  )
}
export default OAuthClients



import React from 'react'
import { Routes, Route } from 'react-router-dom'

import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'

import ValidateRouteParam from '@ttn-lw/lib/components/validate-route-param'
import GenericNotFound from '@ttn-lw/lib/components/full-view-error/not-found'

import AuthorizationsList from '@account/views/oauth-client-authorizations-list'
import AuthorizationOverview from '@account/views/oauth-client-authorization-overview'

import { pathId as pathIdRegexp } from '@ttn-lw/lib/regexp'
import sharedMessages from '@ttn-lw/lib/shared-messages'

const OAuthClientAuthorizations = () => {
  useBreadcrumbs(
    'client-authorizations',
    <Breadcrumb path="/client-authorizations" content={sharedMessages.oauthClientAuthorizations} />,
  )

  return (
    <Routes>
      <Route index Component={AuthorizationsList} />
      <Route
        path=":clientId/*"
        element={
          <ValidateRouteParam
            check={{ clientId: pathIdRegexp }}
            Component={AuthorizationOverview}
          />
        }
      />
      <Route path="*" element={<GenericNotFound />} />
    </Routes>
  )
}

export default OAuthClientAuthorizations

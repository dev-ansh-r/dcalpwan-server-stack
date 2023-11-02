
import React from 'react'
import { Routes, Route, useParams } from 'react-router-dom'

import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'

import ErrorView from '@ttn-lw/lib/components/error-view'
import ValidateRouteParam from '@ttn-lw/lib/components/validate-route-param'
import GenericNotFound from '@ttn-lw/lib/components/full-view-error/not-found'

import Require from '@console/lib/components/require'

import GatewayApiKeyEdit from '@console/views/gateway-api-key-edit'
import GatewayApiKeyAdd from '@console/views/gateway-api-key-add'
import GatewayApiKeysList from '@console/views/gateway-api-keys-list'
import SubViewError from '@console/views/sub-view-error'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import { apiKeyPath as apiKeyPathRegexp } from '@console/lib/regexp'
import { mayViewOrEditGatewayApiKeys } from '@console/lib/feature-checks'

const GatewayApiKeys = () => {
  const { gtwId } = useParams()

  useBreadcrumbs(
    'gtws.single.api-keys',
    <Breadcrumb path={`/gateways/${gtwId}/api-keys`} content={sharedMessages.apiKeys} />,
  )

  return (
    <Require
      featureCheck={mayViewOrEditGatewayApiKeys}
      otherwise={{ redirect: `/gateways/${gtwId}` }}
    >
      <ErrorView errorRender={SubViewError}>
        <Routes>
          <Route index Component={GatewayApiKeysList} />
          <Route path="add" Component={GatewayApiKeyAdd} />
          <Route
            path=":apiKeyId/*"
            element={
              <ValidateRouteParam
                check={{ apiKeyId: apiKeyPathRegexp }}
                Component={GatewayApiKeyEdit}
              />
            }
          />
          <Route path="*" Component={GenericNotFound} />
        </Routes>
      </ErrorView>
    </Require>
  )
}

export default GatewayApiKeys

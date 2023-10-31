

import React from 'react'
import { Routes, Route, useParams } from 'react-router-dom'

import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'

import ErrorView from '@ttn-lw/lib/components/error-view'
import ValidateRouteParam from '@ttn-lw/lib/components/validate-route-param'
import GenericNotFound from '@ttn-lw/lib/components/full-view-error/not-found'

import Require from '@console/lib/components/require'

import GatewayCollaboratorEdit from '@console/views/gateway-collaborator-edit'
import GatewayCollaboratorAdd from '@console/views/gateway-collaborator-add'
import GatewayCollaboratorsList from '@console/views/gateway-collaborators-list'
import SubViewError from '@console/views/sub-view-error'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import { userPathId as userPathIdRegexp } from '@ttn-lw/lib/regexp'

import { mayViewOrEditGatewayCollaborators } from '@console/lib/feature-checks'

const GatewayCollaborators = () => {
  const { gtwId } = useParams()

  useBreadcrumbs(
    'gtws.single.collaborators',
    <Breadcrumb path={`/gateways/${gtwId}/collaborators`} content={sharedMessages.collaborators} />,
  )

  return (
    <Require
      featureCheck={mayViewOrEditGatewayCollaborators}
      otherwise={{ redirect: `/gateways/${gtwId}` }}
    >
      <ErrorView errorRender={SubViewError}>
        <Routes>
          <Route index Component={GatewayCollaboratorsList} />
          <Route path="add" Component={GatewayCollaboratorAdd} />
          <Route
            path=":collaboratorType/:collaboratorId"
            element={
              <ValidateRouteParam
                check={{
                  collaboratorType: /^user$|^organization$/,
                  collaboratorId: userPathIdRegexp,
                }}
                Component={GatewayCollaboratorEdit}
              />
            }
          />
          <Route path="*" element={<GenericNotFound />} />
        </Routes>
      </ErrorView>
    </Require>
  )
}

export default GatewayCollaborators

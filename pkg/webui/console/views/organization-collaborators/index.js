

import React from 'react'
import { Routes, Route, useParams } from 'react-router-dom'

import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'

import ErrorView from '@ttn-lw/lib/components/error-view'
import ValidateRouteParam from '@ttn-lw/lib/components/validate-route-param'
import GenericNotFound from '@ttn-lw/lib/components/full-view-error/not-found'

import Require from '@console/lib/components/require'

import SubViewError from '@console/views/sub-view-error'
import OrganizationCollaboratorsList from '@console/views/organization-collaborators-list'
import OrganizationCollaboratorAdd from '@console/views/organization-collaborator-add'
import OrganizationCollaboratorEdit from '@console/views/organization-collaborator-edit'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import { userPathId as userPathIdRegexp } from '@ttn-lw/lib/regexp'

import { mayViewOrEditOrganizationCollaborators } from '@console/lib/feature-checks'

const OrganizationCollaborators = () => {
  const { orgId } = useParams()

  useBreadcrumbs(
    'orgs.single.collaborators',
    <Breadcrumb
      path={`/organizations/${orgId}/collaborators`}
      content={sharedMessages.collaborators}
    />,
  )

  return (
    <Require
      featureCheck={mayViewOrEditOrganizationCollaborators}
      otherwise={{ redirect: `/organizations/${orgId}` }}
    >
      <ErrorView errorRender={SubViewError}>
        <Routes>
          <Route index Component={OrganizationCollaboratorsList} />
          <Route path="add" Component={OrganizationCollaboratorAdd} />
          <Route
            path="user/:collaboratorId"
            element={
              <ValidateRouteParam
                check={{ collaboratorId: userPathIdRegexp }}
                Component={OrganizationCollaboratorEdit}
              />
            }
          />
          <Route path="*" element={<GenericNotFound />} />
        </Routes>
      </ErrorView>
    </Require>
  )
}

export default OrganizationCollaborators



import React from 'react'
import { Routes, Route, useParams } from 'react-router-dom'

import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'

import ErrorView from '@ttn-lw/lib/components/error-view'
import ValidateRouteParam from '@ttn-lw/lib/components/validate-route-param'
import GenericNotFound from '@ttn-lw/lib/components/full-view-error/not-found'

import ApplicationCollaboratorsList from '@console/views/application-collaborators-list'
import ApplicationCollaboratorEdit from '@console/views/application-collaborator-edit'
import SubViewError from '@console/views/sub-view-error'
import ApplicationCollaboratorAdd from '@console/views/application-collaborator-add'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import { userPathId as userPathIdRegexp } from '@ttn-lw/lib/regexp'

const ApplicationCollaborators = () => {
  const { appId } = useParams()

  useBreadcrumbs(
    'apps.single.collaborators',
    <Breadcrumb
      path={`/applications/${appId}/collaborators`}
      content={sharedMessages.collaborators}
    />,
  )

  return (
    <ErrorView errorRender={SubViewError}>
      <Routes>
        <Route index Component={ApplicationCollaboratorsList} />
        <Route path="add" Component={ApplicationCollaboratorAdd} />
        <Route
          path=":collaboratorType/:collaboratorId"
          element={
            <ValidateRouteParam
              check={{
                collaboratorType: /^user$|^organization$/,
                collaboratorId: userPathIdRegexp,
              }}
              Component={ApplicationCollaboratorEdit}
            />
          }
        />
        <Route path="*" element={<GenericNotFound />} />
      </Routes>
    </ErrorView>
  )
}

export default ApplicationCollaborators

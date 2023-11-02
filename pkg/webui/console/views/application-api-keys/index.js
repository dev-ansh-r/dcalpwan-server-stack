
import React from 'react'
import { Routes, Route, useParams } from 'react-router-dom'

import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'

import ErrorView from '@ttn-lw/lib/components/error-view'
import ValidateRouteParam from '@ttn-lw/lib/components/validate-route-param'
import GenericNotFound from '@ttn-lw/lib/components/full-view-error/not-found'

import Require from '@console/lib/components/require'

import ApplicationApiKeysList from '@console/views/application-api-keys-list'
import ApplicationApiKeyAdd from '@console/views/application-api-key-add'
import SubViewError from '@console/views/sub-view-error'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import { apiKeyPath as apiKeyPathRegexp } from '@console/lib/regexp'
import { mayViewOrEditApplicationApiKeys } from '@console/lib/feature-checks'

import ApplicationApiKeyEdit from '../application-api-key-edit'

const ApplicationApiKeys = () => {
  const { appId } = useParams()

  useBreadcrumbs(
    'apps.single.api-keys',
    <Breadcrumb path={`/applications/${appId}/api-keys`} content={sharedMessages.apiKeys} />,
  )

  return (
    <Require
      featureCheck={mayViewOrEditApplicationApiKeys}
      otherwise={{ redirect: `/applications/${appId}` }}
    >
      <ErrorView errorRender={SubViewError}>
        <Routes>
          <Route index Component={ApplicationApiKeysList} />
          <Route path="/add" Component={ApplicationApiKeyAdd} />
          <Route
            path="/:apiKeyId"
            element={
              <ValidateRouteParam
                check={{ apiKeyId: apiKeyPathRegexp }}
                Component={ApplicationApiKeyEdit}
              />
            }
          />
          <Route path="*" element={<GenericNotFound />} />
        </Routes>
      </ErrorView>
    </Require>
  )
}

export default ApplicationApiKeys

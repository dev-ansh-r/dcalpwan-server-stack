

import React from 'react'
import { Routes, Route, useParams } from 'react-router-dom'

import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'

import ErrorView from '@ttn-lw/lib/components/error-view'
import ValidateRouteParam from '@ttn-lw/lib/components/validate-route-param'
import GenericNotFound from '@ttn-lw/lib/components/full-view-error/not-found'

import Require from '@console/lib/components/require'

import SubViewError from '@console/views/sub-view-error'
import OrganizationApiKeysList from '@console/views/organization-api-keys-list'
import OrganizationApiKeyAdd from '@console/views/organization-api-key-add'
import OrganizationApiKeyEdit from '@console/views/organization-api-key-edit'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import { apiKeyPath as apiKeyPathRegexp } from '@console/lib/regexp'
import { mayViewOrEditOrganizationApiKeys } from '@console/lib/feature-checks'

const OrganizationApiKeys = () => {
  const { orgId } = useParams()

  useBreadcrumbs(
    'org.single.api-keys',
    <Breadcrumb path={`/organizations/${orgId}/api-keys`} content={sharedMessages.apiKeys} />,
  )

  return (
    <Require
      featureCheck={mayViewOrEditOrganizationApiKeys}
      otherwise={{ redirect: `/organizations/${orgId}` }}
    >
      <ErrorView errorRender={SubViewError}>
        <Routes>
          <Route index Component={OrganizationApiKeysList} />
          <Route path="add" Component={OrganizationApiKeyAdd} />
          <Route
            path=":apiKeyId"
            element={
              <ValidateRouteParam
                check={{ apiKeyId: apiKeyPathRegexp }}
                Component={OrganizationApiKeyEdit}
              />
            }
          />
          <Route path="*" element={<GenericNotFound />} />
        </Routes>
      </ErrorView>
    </Require>
  )
}

export default OrganizationApiKeys

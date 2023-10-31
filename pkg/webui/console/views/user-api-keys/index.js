

import React from 'react'
import { Routes, Route } from 'react-router-dom'

import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'

import ErrorView from '@ttn-lw/lib/components/error-view'
import GenericNotFound from '@ttn-lw/lib/components/full-view-error/not-found'
import ValidateRouteParam from '@ttn-lw/lib/components/validate-route-param'

import UserApiKeyEdit from '@console/views/user-api-key-edit'
import UserApiKeyAdd from '@console/views/user-api-key-add'
import SubViewError from '@console/views/sub-view-error'
import UserApiKeysList from '@console/views/user-api-keys-list'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import { apiKeyPath as apiKeyPathRegexp } from '@console/lib/regexp'

const UserApiKeys = () => {
  useBreadcrumbs(
    'usr.single.api-keys',
    <Breadcrumb path={`/user/api-keys`} content={sharedMessages.personalApiKeys} />,
  )

  return (
    <ErrorView errorRender={SubViewError}>
      <Routes>
        <Route index Component={UserApiKeysList} />
        <Route path="add" Component={UserApiKeyAdd} />
        <Route
          path=":apiKeyId/*"
          element={
            <ValidateRouteParam check={{ apiKeyId: apiKeyPathRegexp }} Component={UserApiKeyEdit} />
          }
        />
        <Route path="*" Component={GenericNotFound} />
      </Routes>
    </ErrorView>
  )
}

export default UserApiKeys

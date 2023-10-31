

import React from 'react'
import { Routes, Route } from 'react-router-dom'

import Breadcrumbs from '@ttn-lw/components/breadcrumbs'

import IntlHelmet from '@ttn-lw/lib/components/intl-helmet'
import GenericNotFound from '@ttn-lw/lib/components/full-view-error/not-found'

import UserApiKeys from '@console/views/user-api-keys'

import { selectApplicationSiteName } from '@ttn-lw/lib/selectors/env'

const UserView = () => (
  <>
    <Breadcrumbs />
    <IntlHelmet titleTemplate={`%s - User - ${selectApplicationSiteName()}`} />
    <Routes>
      <Route path="api-keys/*" Component={UserApiKeys} />
      <Route path="*" Component={GenericNotFound} />
    </Routes>
  </>
)

export default UserView

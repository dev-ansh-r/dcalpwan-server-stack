

import React from 'react'
import { Routes, Route } from 'react-router-dom'

import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'

import GenericNotFound from '@ttn-lw/lib/components/full-view-error/not-found'
import ValidateRouteParam from '@ttn-lw/lib/components/validate-route-param'

import withFeatureRequirement from '@console/lib/components/with-feature-requirement'

import Application from '@console/views/application'
import ApplicationsList from '@console/views/applications-list'
import ApplicationAdd from '@console/views/application-add'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import { pathId as pathIdRegexp } from '@ttn-lw/lib/regexp'

import { mayViewApplications } from '@console/lib/feature-checks'

const Applications = () => {
  useBreadcrumbs('apps', <Breadcrumb path="/applications" content={sharedMessages.applications} />)

  return (
    <Routes>
      <Route index Component={ApplicationsList} />
      <Route path="add" Component={ApplicationAdd} />
      <Route
        path=":appId/*"
        element={<ValidateRouteParam check={{ appId: pathIdRegexp }} Component={Application} />}
      />
      <Route path="*" Component={GenericNotFound} />
    </Routes>
  )
}
export default withFeatureRequirement(mayViewApplications, { redirect: '/' })(Applications)

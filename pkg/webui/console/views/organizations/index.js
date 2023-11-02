
import React from 'react'
import { Routes, Route } from 'react-router-dom'

import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'

import GenericNotFound from '@ttn-lw/lib/components/full-view-error/not-found'
import ValidateRouteParam from '@ttn-lw/lib/components/validate-route-param'

import Require from '@console/lib/components/require'

import Organization from '@console/views/organization'
import OrganizationAdd from '@console/views/organization-add'
import OrganizationsList from '@console/views/organizations-list'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import { pathId as pathIdRegexp } from '@ttn-lw/lib/regexp'

import { mayViewOrganizationsOfUser } from '@console/lib/feature-checks'

const Organizations = () => {
  useBreadcrumbs(
    'orgs',
    <Breadcrumb path="/organizations" content={sharedMessages.organizations} />,
  )

  return (
    <Require featureCheck={mayViewOrganizationsOfUser} otherwise={{ redirect: '/' }}>
      <Routes>
        <Route index Component={OrganizationsList} />
        <Route path="add" Component={OrganizationAdd} />
        <Route
          path=":orgId/*"
          element={<ValidateRouteParam check={{ orgId: pathIdRegexp }} Component={Organization} />}
        />
        <Route path="*" element={<GenericNotFound />} />
      </Routes>
    </Require>
  )
}

export default Organizations

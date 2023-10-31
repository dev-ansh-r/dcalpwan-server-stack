

import React from 'react'

import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'
import PageTitle from '@ttn-lw/components/page-title'

import RequireRequest from '@ttn-lw/lib/components/require-request'

import NetworkInformationContainer from '@console/containers/network-information-container'
import DeploymentComponentStatus from '@console/containers/deployment-component-status'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import { getApplicationsList } from '@console/store/actions/applications'
import { getGatewaysList } from '@console/store/actions/gateways'
import { getUsersList } from '@console/store/actions/users'
import { getOrganizationsList } from '@console/store/actions/organizations'

const NetworkInformation = () => {
  useBreadcrumbs(
    'admin-panel.network-information',
    <Breadcrumb
      path={`/admin-panel/network-information`}
      content={sharedMessages.networkInformation}
    />,
  )

  const requestActions = [
    getApplicationsList(),
    getGatewaysList(),
    getUsersList(),
    getOrganizationsList(),
  ]

  return (
    <>
      <RequireRequest requestAction={requestActions}>
        <PageTitle title={sharedMessages.networkInformation} className="panel-title mb-0" />
        <NetworkInformationContainer />
        <DeploymentComponentStatus />
      </RequireRequest>
    </>
  )
}

export default NetworkInformation

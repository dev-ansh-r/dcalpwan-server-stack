

import React from 'react'

import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'
import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'

import RequireRequest from '@ttn-lw/lib/components/require-request'

import UserDataFormAdd from '@console/containers/user-data-form/add'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import { getIsConfiguration } from '@console/store/actions/identity-server'

const UserManagementAdd = () => {
  useBreadcrumbs(
    'admin-panel.user-management.add',
    <Breadcrumb path={`/admin-panel/user-management/add`} content={sharedMessages.add} />,
  )

  return (
    <RequireRequest requestAction={getIsConfiguration()}>
      <UserDataFormAdd />
    </RequireRequest>
  )
}

export default UserManagementAdd
